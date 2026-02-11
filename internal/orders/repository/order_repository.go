package repository

import "github.com/dreezy305/library-core-service/internal/model"

type OrderRepository struct {
	gormRepo *GormOrderRepository
}

func NewOrderRepository(gormRepo *GormOrderRepository) *OrderRepository {
	return &OrderRepository{gormRepo: gormRepo}
}

func (s *OrderRepository) CreateOrder(order *model.OrderEntity) error {
	return s.gormRepo.CreateOrder(order)
}

func (s *OrderRepository) GetOrderByID(id string) (*model.OrderEntity, error) {
	return s.gormRepo.GetOrderByID(id)
}

func (s *OrderRepository) ListOrdersByUserID(userID string) ([]*model.OrderEntity, error) {
	return s.gormRepo.ListOrdersByUserID(userID)
}
