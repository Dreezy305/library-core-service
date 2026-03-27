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

func (s *PaymentService) UpdatePaymentStatus(paymentId string, status string) error {
	return s.repo.UpdatePaymentStatus(s.DB, paymentId, status)
}

func (s *PaymentService) UpdatePaymentInfoTx(tx *gorm.DB, paymentId string, payload *types.UpdatePaymentPayload) error {
	var status constants.PaymentStatus
	switch payload.Status {
	case "success":
		status = constants.PaymentPaid
	case "failed":
		status = constants.PaymentFailed
	default:
		status = constants.PaymentPending
	}
	payment := &model.PaymentEntity{
		Status:           status,
		Reference:        payload.Reference,
		PaymentGateway:   payload.PaymentGateway,
		WebhookProcessed: true,
		PaymentMethod:    payload.PaymentGateway,
		Currency:         payload.Currency,
		Metadata:         json.RawMessage(payload.Metadata),
	}
	return s.repo.UpdatePaymentInfo(tx, paymentId, payment)
}

func (s *PaymentService) UpdatePaymentInfo(paymentId string, payload *types.UpdatePaymentPayload) error {

	var status constants.PaymentStatus
	switch payload.Status {
	case "success":
		status = constants.PaymentPaid
	case "failed":
		status = constants.PaymentFailed
	default:
		status = constants.PaymentPending
	}

	payment := &model.PaymentEntity{
		Status:           status,
		Reference:        payload.Reference,
		PaymentGateway:   payload.PaymentGateway,
		WebhookProcessed: true,
		PaymentMethod:    payload.PaymentGateway,
		Currency:         payload.Currency,
		Metadata:         json.RawMessage(payload.Metadata),
	}

	return s.repo.UpdatePaymentInfo(s.DB, paymentId, payment)
}
