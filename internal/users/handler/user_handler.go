package handler

import (
	"fmt"
	"strconv"

	"github.com/dreezy305/library-core-service/internal/types"
	"github.com/dreezy305/library-core-service/internal/users/service"
	"github.com/dreezy305/library-core-service/internal/validators"
	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) GetUsers(c fiber.Ctx) error {
	queries := c.Queries()

	page, err := strconv.Atoi(queries["page"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid value for paramater page"})
	}

	limit, errr := strconv.Atoi(queries["limit"])

	if errr != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Invalid value for paramater limit"})
	}
	users, total, _ := h.Service.GetUsers(page, limit)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Users fetched successfully", "data": fiber.Map{"users": users, "meta": fiber.Map{"total": total, "page": page, "limit": limit}}})
}

func (h *UserHandler) GetUser(c fiber.Ctx) error {
	userId := c.Params("id")
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "user id parameter is missing"})
	}

	user, err := h.Service.GetUser(userId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "user not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User fetched successfully", "data": fiber.Map{"user": user}})
}

func (h *UserHandler) UpdateUser(c fiber.Ctx) error {
	userId := c.Params("id")
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "user id parameter is missing"})
	}
	var payload types.UpdateUser

	if err := c.Bind().Body(&payload); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusForbidden).JSON(validators.FormatValidationError(err))
	}

	err := h.Service.UpdateUser(userId, payload)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to update user"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User upated successfully"})
}
