package repository

import (
	"github.com/dreezy305/library-core-service/internal/model"
	"gorm.io/gorm"
)

type GormOrderRepository struct {
	DB *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) *GormOrderRepository {
	return &GormOrderRepository{DB: db}
}

func (r *GormOrderRepository) CreateOrder(order *model.OrderEntity) error {
	return r.DB.Create(order).Error
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
