package handler

import (
	"github.com/dreezy305/library-core-service/internal/books/service"
	"github.com/gofiber/fiber/v3"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) GetBooks(c fiber.Ctx) error {
	return h.service.GetBooks()
}

func (h *BookHandler) GetBook(c fiber.Ctx) error {
	bookId := c.Params("id")
	return h.service.GetBook(bookId)
}

func (h *BookHandler) CreateBook(c fiber.Ctx) error {
	return h.service.CreateBook()
}

func (h *BookHandler) UpdateBook(c fiber.Ctx) error {
	return h.service.UpdateBook()
}
