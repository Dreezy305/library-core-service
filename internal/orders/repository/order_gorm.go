package repository

import (
	"time"

	"github.com/dreezy305/library-core-service/internal/model"
	"gorm.io/gorm"
)

type GormOrderRepository struct {
	DB *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) *GormOrderRepository {
	return &GormOrderRepository{DB: db}
}

func (r *GormOrderRepository) CreateOrder(tx *gorm.DB, order *model.OrderEntity) error {
	return tx.Create(order).Error
}

func (r *GormOrderRepository) CreateOrderItems(tx *gorm.DB, items []*model.OrderItemEntity) error {
	return tx.Create(&items).Error // batch insert
}

func (r *GormOrderRepository) UpdateOrderStatus(tx *gorm.DB, orderId string, status string) error {
	var order model.OrderEntity
	err := tx.Preload("Items").Find(&order, "id = ?", orderId).Error
	if err != nil {
		return err
	}
	if order.ID == "" {
		return gorm.ErrRecordNotFound
	}

	if status != "" {
		order.Status = status
		now := time.Now()
		order.PaidAt = &now
	}

	return tx.Save(&order).Error
}

func (r *GormOrderRepository) UpdateOrderItems(tx *gorm.DB, ItemId string, status string) error {
	var item model.OrderItemEntity
	err := tx.Find(&item, "id = ?", ItemId).Error
	if err != nil {
		return err
	}
	if item.ID == "" {
		return gorm.ErrRecordNotFound
	}

	if status != "" {
		item.Status = status
	}

	return tx.Save(&item).Error
}

func (r *GormOrderRepository) GetOrderByID(id string) (*model.OrderEntity, error) {
	var order model.OrderEntity
	err := r.DB.Preload("Items").Preload("User").First(&order, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *GormOrderRepository) ListOrdersByUserID(userID string) ([]*model.OrderEntity, error) {
	var orders []*model.OrderEntity
	err := r.DB.Preload("Items").Preload("User").Where("user_id = ?", userID).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
