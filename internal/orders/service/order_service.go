package service

import (
	"errors"
	"fmt"

	BookRepository "github.com/dreezy305/library-core-service/internal/books/repository"
	"github.com/dreezy305/library-core-service/internal/constants"
	"github.com/dreezy305/library-core-service/internal/model"
	"github.com/dreezy305/library-core-service/internal/orders/repository"
	"github.com/dreezy305/library-core-service/internal/types"
	UserRepository "github.com/dreezy305/library-core-service/internal/users/repository"
	"gorm.io/gorm"
)

type OrderService struct {
	repo     repository.OrderRepository
	DB       *gorm.DB
	bookRepo *BookRepository.BookRepository
	userRepo *UserRepository.UserRepository
}

func NewOrderService(repo repository.OrderRepository, bookRepo *BookRepository.BookRepository, userRepo *UserRepository.UserRepository, db *gorm.DB) *OrderService {
	return &OrderService{repo: repo, bookRepo: bookRepo, userRepo: userRepo, DB: db}
}

func (s *OrderService) CreateOrder(payload types.InitiateOrderPayload) error {

	return s.DB.Transaction(func(tx *gorm.DB) error {

		// 1️⃣ Validate user exists
		user, err := s.userRepo.GetUser(payload.UserID)
		if err != nil {
			return err
		}
		if user == nil {
			return errors.New("user not found")
		}

		var orderTotal int64
		orderItems := make([]*model.OrderItemEntity, 0, len(payload.Items))

		// 2️⃣ Loop through requested items
		for _, item := range payload.Items {

			book, err := s.bookRepo.GetBook(item.BookID)
			if err != nil {
				return err
			}
			if book == nil {
				return fmt.Errorf("book %s not found", item.BookID)
			}

			// Optional: check stock
			if book.CopiesAvailable < item.Quantity {
				return fmt.Errorf("not enough copies for book %s", book.Title)
			}

			unitPrice := book.Price
			totalPrice := unitPrice * int64(item.Quantity)

			orderTotal += totalPrice

			orderItems = append(orderItems, &model.OrderItemEntity{
				BookID:     book.ID,
				Quantity:   item.Quantity,
				UnitPrice:  unitPrice,
				TotalPrice: totalPrice,
				Status:     string(constants.OrderPending),
			})
		}

		// 3️⃣ Create order
		order := &model.OrderEntity{
			UserID:      payload.UserID,
			Status:      string(constants.OrderPending),
			TotalAmount: orderTotal,
		}

		if err := s.repo.CreateOrder(tx, order); err != nil {
			return err
		}

		// 4️⃣ Assign OrderID to each item
		for _, item := range orderItems {
			item.OrderID = order.ID
		}

		// 5️⃣ Batch insert
		if err := s.repo.CreateOrderItems(tx, orderItems); err != nil {
			return err
		}

		return nil
	})
}

func (s *OrderService) MarkOrderAsPaid(orderId string) error {
	return s.DB.Transaction(func(tx *gorm.DB) error {
		order, err := s.repo.GetOrderByID(tx, orderId)
		if err != nil {
			return err
		}
		if order == nil {
			return errors.New("order not found")
		}

		if order.Status != string(constants.OrderPending) {
			return errors.New("order has already been paid or is in an invalid state")
		}

		// UPDATE ORDER STATUS
		if err := s.repo.UpdateOrderStatus(tx, orderId, string(constants.OrderPaid)); err != nil {
			return err
		}

		orderItems := order.Items
		for _, item := range orderItems {
			if err := s.repo.UpdateOrderItemStatus(tx, item.ID, string(constants.OrderPaid)); err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *OrderService) GetOrderByID(id string) (*model.OrderEntity, error) {
	return s.repo.GetOrderByID(s.DB, id)
}

func (s *OrderService) ListOrdersByUserID(userID string) ([]*model.OrderEntity, error) {
	return s.repo.ListOrdersByUserID(userID)
}
