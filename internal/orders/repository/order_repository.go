package repository

import (
	"github.com/dreezy305/library-core-service/internal/model"
	"gorm.io/gorm"
)

type OrderRepository struct {
	gormRepo *GormOrderRepository
}

func NewOrderRepository(gormRepo *GormOrderRepository) *OrderRepository {
	return &OrderRepository{gormRepo: gormRepo}
}

func (s *OrderRepository) CreateOrder(tx *gorm.DB, order *model.OrderEntity) error {
	return s.gormRepo.CreateOrder(tx, order)
}

func (s *OrderRepository) CreateOrderItems(tx *gorm.DB, orderItems []*model.OrderItemEntity) error {

	return s.gormRepo.CreateOrderItems(tx, orderItems)
}

func (s *OrderRepository) GetOrderByID(id string) (*model.OrderEntity, error) {
	return s.gormRepo.GetOrderByID(id)
}

func (s *OrderRepository) UpdateOrderStatus(tx *gorm.DB, orderId string, status string) error {
	return s.gormRepo.UpdateOrderStatus(tx, orderId, status)
}

func (s *OrderRepository) UpdateOrderItems(tx *gorm.DB, ItemId string, status string) error {
	return s.gormRepo.UpdateOrderItems(tx, ItemId, status)
}

func (s *OrderRepository) ListOrdersByUserID(userID string) ([]*model.OrderEntity, error) {
	return s.gormRepo.ListOrdersByUserID(userID)
}
