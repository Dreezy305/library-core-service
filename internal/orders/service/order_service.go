package service

import (
	"github.com/dreezy305/library-core-service/internal/model"
	"github.com/dreezy305/library-core-service/internal/orders/repository"
)

type OrderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order *model.OrderEntity) error {
	return s.repo.CreateOrder(order)
}

func (s *OrderService) GetOrderByID(id string) (*model.OrderEntity, error) {
	return s.repo.GetOrderByID(id)
}

func (s *OrderService) ListOrdersByUserID(userID string) ([]*model.OrderEntity, error) {
	return s.repo.ListOrdersByUserID(userID)
}
