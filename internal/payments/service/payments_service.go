package service

import (
	"encoding/json"

	"github.com/dreezy305/library-core-service/internal/constants"
	"github.com/dreezy305/library-core-service/internal/model"
	"github.com/dreezy305/library-core-service/internal/payments/repository"
	"github.com/dreezy305/library-core-service/internal/types"
	"gorm.io/gorm"
)

type PaymentService struct {
	repo repository.PaymentRepository
	DB   *gorm.DB
}

func NewPaymentService(repo repository.PaymentRepository, db *gorm.DB) *PaymentService {
	return &PaymentService{
		repo: repo,
		DB:   db,
	}
}

func (s *PaymentService) InitializePayment(payload *types.InitiatePaymentPayload) error {
	payment := &model.PaymentEntity{
		OrderID:          payload.OrderID,
		Amount:           payload.Amount,
		PaymentGateway:   payload.PaymentGateway,
		Currency:         payload.Currency,
		Metadata:         json.RawMessage(payload.Metadata),
		WebhookProcessed: false,
		Status:           constants.PaymentPending,
		Reference:        payload.Reference,
	}
	return s.repo.InitializePayment(s.DB, payment)
}
