package handler

import (
	"fmt"

	"github.com/dreezy305/library-core-service/internal/categories/service"
	"github.com/dreezy305/library-core-service/internal/model"
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

	cModel := &model.CategoryEntity{
		Name: payload.Name,
	}

	if payload.Description != nil {
		cModel.Description = payload.Description
	}

	error := h.Service.CreateCategory(cModel)

	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create category"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Category created successfully"})
}

func (h *CategoryHandler) GetCategories() ([]string, error) {
	return h.Service.GetCategories()
}

func (h *CategoryHandler) DeleteCategory(name string) error {
	return h.Service.DeleteCategory(name)
}
