package handler

import "github.com/gofiber/fiber/v3"

type OrderHandler struct {
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{}
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	return nil
}
