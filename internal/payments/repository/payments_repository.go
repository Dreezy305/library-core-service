package repository

import (
	"github.com/dreezy305/library-core-service/internal/model"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	gormRepo *GormPaymentRepository
}

func NewPaymentRepository(gormRepo *GormPaymentRepository) *PaymentRepository {
	return &PaymentRepository{
		gormRepo: gormRepo,
	}
}

func (r *PaymentRepository) InitializePayment(tx *gorm.DB, payment *model.PaymentEntity) error {
	return r.gormRepo.InitializePayment(tx, payment)
}
