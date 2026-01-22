package handler

import (
	"fmt"

	"github.com/dreezy305/library-core-service/internal/categories/service"
	"github.com/dreezy305/library-core-service/internal/types"
	"github.com/dreezy305/library-core-service/internal/validators"
	"github.com/gofiber/fiber/v3"
)

type CategoryHandler struct {
	Service service.CategoryService
}

func NewCategoryHandler(s *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{Service: *s}
}

func (h *CategoryHandler) CreateCategory(c fiber.Ctx) error {
	var payload types.CategoryPayload
	if err := c.Bind().Body(&payload); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusForbidden).JSON(validators.FormatValidationError(err))
	}

	fmt.Println(payload)

	errs := validators.ValidateStruct(payload)
	if errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validators.FormatValidationError(errs))
	}

	error := h.Service.CreateCategory(payload)

	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to create category"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Category created successfully"})
}

func (h *CategoryHandler) GetCategories(c fiber.Ctx) error {
	categories, err := h.Service.GetCategories()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to get categories"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": fiber.Map{"categories": categories}})
}

func (h *CategoryHandler) DeleteCategory(c fiber.Ctx) error {
	categoryId := c.Params("id")
	if categoryId == "" {
		return fmt.Errorf("id parameter is missing")
	}
	error := h.Service.DeleteCategory(categoryId)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to delete category"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Category deleted successfully"})
}
