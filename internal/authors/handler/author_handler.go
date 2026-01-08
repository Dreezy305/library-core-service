package handler

import (
	"github.com/dreezy305/library-core-service/internal/authors/service"
	"github.com/gofiber/fiber/v3"
)

type AuthorHandler struct {
	Service *service.AuthorService
}

func NewAuthorHandler(s *service.AuthorService) *AuthorHandler {
	return &AuthorHandler{Service: s}
}

func (h *AuthorHandler) CreateAuthor(c fiber.Ctx) error {

	return nil
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
