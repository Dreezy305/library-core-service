package handler

import (
	"fmt"

	"github.com/dreezy305/library-core-service/internal/authors/service"
	"github.com/dreezy305/library-core-service/internal/model"
	"github.com/dreezy305/library-core-service/internal/types"
	"github.com/dreezy305/library-core-service/internal/validators"
	"github.com/gofiber/fiber/v3"
)

type AuthorHandler struct {
	Service *service.AuthorService
}

func NewAuthorHandler(s *service.AuthorService) *AuthorHandler {
	return &AuthorHandler{Service: s}
}

func (h *AuthorHandler) CreateAuthor(c fiber.Ctx) error {
	var payload types.AuthorType

	if err := c.Bind().Body(&payload); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusForbidden).JSON(validators.FormatValidationError(err))
	}
	
	fmt.Println(payload)

	errs := validators.ValidateStruct(payload)
	if errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validators.FormatValidationError(errs))
	}

	u := &model.AuthorEntity{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}

	error := h.Service.CreateAuthor(u)

	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to create author"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Author created successfully"})
}

func (h *AuthorHandler) GetAuthors(c fiber.Ctx) error {
	return nil
}

func (h *AuthorHandler) GetAuthor(c fiber.Ctx) error {
	return nil
}

func (h *AuthorHandler) UpdateAuthor(c fiber.Ctx) error {
	return nil
}

func (h *AuthorHandler) GetAuthorBooksByAuthorId(c fiber.Ctx) error {
	return nil
}
