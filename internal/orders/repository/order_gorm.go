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
	return tx.Model(&model.OrderEntity{}).
		Where("id = ?", orderId).
		Updates(map[string]interface{}{
			"Status": status,
			"PaidAt": time.Now(),
		}).Error
}

func (r *GormOrderRepository) UpdateOrderItemStatus(tx *gorm.DB, orderId string, status string) error {
	return tx.Model(&model.OrderItemEntity{}).
		Where("order_id = ?", orderId).
		Update("Status", status).Error
}

func (r *GormOrderRepository) GetOrderByID(tx *gorm.DB, id string) (*model.OrderEntity, error) {
	var order model.OrderEntity
	err := tx.Preload("Items").Preload("User").Find(&order, "id = ?", id).Error
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
