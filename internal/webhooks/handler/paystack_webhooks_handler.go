package handler

import (
	"github.com/dreezy305/library-core-service/internal/webhooks/service"
	"github.com/gofiber/fiber/v3"
)

type PayStackWebhookHandler struct {
	service *service.PaystackWebhookService
}

func NewPayStackWebhookHandler(service *service.PaystackWebhookService) *PayStackWebhookHandler {
	return &PayStackWebhookHandler{service: service}
}

func (h *PayStackWebhookHandler) VerifySignature(c fiber.Ctx) error {
	signature := c.Get("x-paystack-signature")
	if signature == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing signature"})
	}

	body := c.Body()
	if err := h.service.VerifySignature(body, signature); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}
