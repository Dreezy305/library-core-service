package handler

import (
	"github.com/dreezy305/library-core-service/internal/orders/service"
	"github.com/dreezy305/library-core-service/internal/types"
	"github.com/dreezy305/library-core-service/internal/validators"
	"github.com/gofiber/fiber/v3"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) CreateOrder(c fiber.Ctx) error {
	var payload types.InitiateOrderPayload
	if err := c.Bind().Body(&payload); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(validators.FormatValidationError(err))
	}

	errs := validators.ValidateStruct(payload)
	if errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validators.FormatValidationError(errs))
	}

	error := h.service.CreateOrder(payload)

	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create order"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Create order endpoint"})
}

func (h *OrderHandler) MarkOrderAsPaid(c fiber.Ctx) error {
	orderId := c.Params("orderId")
	if orderId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "id parameter is missing"})
	}
	return c.JSON(fiber.Map{"message": "Update order endpoint"})
}

func (h *OrderHandler) GetOrder(c fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get order endpoint"})
}

func (h *OrderHandler) ListOrdersByUserID(c fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "List orders by user ID endpoint"})
}
