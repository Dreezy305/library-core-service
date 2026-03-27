package repository

import (
	"github.com/dreezy305/library-core-service/internal/model"
	"gorm.io/gorm"
)

type GormPaymentRepository struct {
	DB *gorm.DB
}

func NewGormPaymentRepository(db *gorm.DB) *GormPaymentRepository {
	return &GormPaymentRepository{
		DB: db,
	}
}

func (r *GormPaymentRepository) InitializePayment(tx *gorm.DB, payment *model.PaymentEntity) error {
	return tx.Create(payment).Error
}

func (r *GormPaymentRepository) UpdatePaymentStatus(tx *gorm.DB, paymentId string, status string) error {
	return tx.Model(&model.PaymentEntity{}).
		Where("id = ?", paymentId).
		Update("Status", status).Error
}

func (r *GormPaymentRepository) UpdatePaymentInfo(tx *gorm.DB, paymentId string, payload *model.PaymentEntity) error {
	return tx.Model(&model.PaymentEntity{}).
		Where("id = ?", paymentId).
		Updates(payload).Error
}
