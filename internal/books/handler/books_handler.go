package handler

import (
	"fmt"
	"strconv"

	"github.com/dreezy305/library-core-service/internal/books/service"
	"github.com/dreezy305/library-core-service/internal/model"
	"github.com/dreezy305/library-core-service/internal/types"
	"github.com/dreezy305/library-core-service/internal/validators"
	"github.com/gofiber/fiber/v3"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) GetBooks(c fiber.Ctx) error {
	queries := c.Queries()

	page, err := strconv.Atoi(queries["page"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid value for paramater page"})
	}

	limit, errr := strconv.Atoi(queries["limit"])

	if errr != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Invalid value for paramater limit"})
	}

	search := queries["search"]

	books, count, err := h.service.GetBooks(page, limit, &search)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to retrieve books"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Books fetched successfully", "data": fiber.Map{"books": books, "meta": fiber.Map{"total": count, "page": page, "limit": limit}}})
}

func (h *BookHandler) GetBook(c fiber.Ctx) error {
	bookId := c.Params("id")
	if bookId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "id parameter is missing"})
	}
	book, err := h.service.GetBook(bookId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to retrieve book"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Book fetched successfully", "data": fiber.Map{"book": book}})
}

func (h *BookHandler) CreateBook(c fiber.Ctx) error {
	var payload types.BookPayload

	if err := c.Bind().Body(&payload); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusForbidden).JSON(validators.FormatValidationError(err))
	}

	fmt.Println(payload)

	errs := validators.ValidateStruct(payload)
	if errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validators.FormatValidationError(errs))
	}

	exist, _ := h.service.BookExists(payload.Title)

	if exist {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Book has already been created"})
	}

	b := &model.BookEntity{
		Title:       payload.Title,
		ISBN:        payload.ISBN,
		CopiesTotal: payload.CopiesTotal,
		AuthorID:    payload.AuthorID,
	}

	if payload.Description != nil {
		b.Description = *payload.Description
	}

	if payload.PublishedYear != nil {
		b.PublishedYear = *payload.PublishedYear
	}

	error := h.service.CreateBook(b)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create book"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Book created successfully"})
}

func (h *BookHandler) UpdateBook(c fiber.Ctx) error {
	var payload types.UpdateBookPayload

	bookId := c.Params("id")
	if bookId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "id parameter is missing"})
	}

	if err := c.Bind().Body(&payload); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusForbidden).JSON(validators.FormatValidationError(err))
	}

	errs := validators.ValidateStruct(payload)
	if errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validators.FormatValidationError(errs))
	}
	er := h.service.UpdateBook(bookId, &payload)
	if er != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update book"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Book updated successfully"})
}
