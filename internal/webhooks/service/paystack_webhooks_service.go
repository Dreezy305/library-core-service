package service

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"gorm.io/gorm"

	"github.com/dreezy305/library-core-service/internal/constants"
	OrderService "github.com/dreezy305/library-core-service/internal/orders/service"
	PaymentService "github.com/dreezy305/library-core-service/internal/payments/service"
	"github.com/dreezy305/library-core-service/internal/types"
)

type PaystackWebhookService struct {
	DB             *gorm.DB
	orderService   *OrderService.OrderService
	paymentService *PaymentService.PaymentService
	paystackSecret string
}

func NewPaystackWebhookService(orderService *OrderService.OrderService, paymentService *PaymentService.PaymentService, db *gorm.DB, paystackSecret string) *PaystackWebhookService {
	return &PaystackWebhookService{orderService: orderService, paymentService: paymentService, DB: db, paystackSecret: paystackSecret}
}

func (s *PaystackWebhookService) HandleWebhook(body []byte, signature string) error {
	mac := hmac.New(sha512.New, []byte(s.paystackSecret))
	mac.Write(body)
	expected := hex.EncodeToString(mac.Sum(nil))
	if !hmac.Equal([]byte(expected), []byte(signature)) {
		return errors.New("unauthorized")
	}
	var event types.PaystackWebhookEvent
	if err := json.Unmarshal(body, &event); err != nil {
		return err
	}
	if event.Event != "charge.success" {
		return nil
	}
	return s.ProcessWebhook(event)
}

func (s *PaystackWebhookService) ProcessWebhook(event types.PaystackWebhookEvent) error {
	// extract key data from event
	reference := event.Data.Reference
	amount := event.Data.Amount
	currency := event.Data.Currency
	var metaData types.PaymentMetadata
	if err := json.Unmarshal(event.Data.Metadata, &metaData); err != nil {
		return err
	}

	// call paystack verify API
	verified, err := s.VerifyTransaction(reference)
	if err != nil {
		return err
	}
	// process verified transaction
	if verified.Data.Status != "success" {
		return fmt.Errorf("transaction not successful")
	}
	if verified.Data.Amount != amount {
		return fmt.Errorf("amount mismatch")
	}
	if verified.Data.Currency != currency {
		return fmt.Errorf("currency mismatch")
	}
	if reference != verified.Data.Reference {
		return fmt.Errorf("reference mismatch")
	}

	// persist all DB changes in one atomic transaction
	return s.DB.Transaction(func(tx *gorm.DB) error {
		order, err := s.orderService.GetOrderByIDTx(tx, metaData.OrderID)
		if err != nil {
			return err
		}
		if order == nil {
			return errors.New("order not found")
		}
		if order.Status == string(constants.OrderPaid) {
			return errors.New("order already processed")
		}

		if err := s.paymentService.UpdatePaymentInfoTx(tx, reference, &types.UpdatePaymentPayload{
			Status:         "success",
			Reference:      reference,
			PaymentGateway: "paystack",
			Currency:       verified.Data.Currency,
			Metadata:       verified.Data.Metadata,
			PaymentMethod:  verified.Data.Channel,
			PaidAt:         func() *time.Time { t := time.Now(); return &t }(),
		}); err != nil {
			return err
		}

		return s.orderService.MarkOrderAsPaidTx(tx, metaData.OrderID)
	})
}

func (s *PaystackWebhookService) VerifyTransaction(reference string) (*types.PaystackVerifyResponse, error) {
	url := "https://api.paystack.co/transaction/verify/" + reference

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Add headers
	req.Header.Set("Authorization", "Bearer "+s.paystackSecret)
	req.Header.Set("Content-Type", "application/json")

	// Make request
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed with status: %d", resp.StatusCode)
	}

	var result types.PaystackVerifyResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
