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

	"gorm.io/gorm"

	BookService "github.com/dreezy305/library-core-service/internal/books/service"
	OrderService "github.com/dreezy305/library-core-service/internal/orders/service"
	PaymentService "github.com/dreezy305/library-core-service/internal/payments/service"
	"github.com/dreezy305/library-core-service/internal/types"
	UserRepository "github.com/dreezy305/library-core-service/internal/users/repository"
)

type PaystackWebhookService struct {
	DB             *gorm.DB
	bookService    *BookService.BookService
	userRepo       *UserRepository.UserRepository
	orderService   *OrderService.OrderService
	paymentService *PaymentService.PaymentService

	paystackSecret string
}

func NewPaystackWebhookService(orderService *OrderService.OrderService, bookService *BookService.BookService, userRepo *UserRepository.UserRepository, paymentService *PaymentService.PaymentService, db *gorm.DB, paystackSecret string) *PaystackWebhookService {
	return &PaystackWebhookService{orderService: orderService, bookService: bookService, userRepo: userRepo, paymentService: paymentService, DB: db, paystackSecret: paystackSecret}
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

	// update payment info
	err = s.paymentService.UpdatePaymentInfo(reference, &types.UpdatePaymentPayload{
		Status:         "success",
		Reference:      reference,
		PaymentGateway: "paystack",
		Currency:       verified.Data.Currency,
		Metadata:       verified.Data.Metadata,
		PaymentMethod:  verified.Data.Channel,
	})
	if err != nil {
		return err
	}

	// mark order as paid
	err = s.orderService.MarkOrderAsPaid(event.Data.Metadata.OrderID)
	if err != nil {
		return err
	}

	// DECREMENT BOOK STOCK
	err = s.bookService.DecrementAvailable(event.Data.Metadata.BookID)
	if err != nil {
		return err
	}

	return nil
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
	client := &http.Client{}
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
