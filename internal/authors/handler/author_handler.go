package handler

import (
	"fmt"
	"strconv"
	"time"

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
	var payload types.AuthorPayload

	if err := c.Bind().Body(&payload); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusForbidden).JSON(validators.FormatValidationError(err))
	}

	fmt.Println(payload)

	errs := validators.ValidateStruct(payload)
	if errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validators.FormatValidationError(errs))
	}

	exist, _ := h.Service.AuthorExist(payload.Email)

	if exist {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Author has already been created"})
	}

	dob, err := time.Parse("2006-01-02", payload.DateOfBirth)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "dateOfBirth must be in YYYY-MM-DD format"})
	}

	u := &model.AuthorEntity{
		FirstName:   payload.FirstName,
		LastName:    payload.LastName,
		Nationality: payload.Nationality,
		DateOfBirth: dob,
		Email:       &payload.Email,
		Bio:         *payload.Bio,
		PenName:     *payload.PenName,
		Website:     *payload.Website,
		Twitter:     *payload.Twitter,
		Facebook:    *payload.Facebook,
		Linkedln:    *payload.Linkedln,
	}

	error := h.Service.CreateAuthor(u)

	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to create author"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Author created successfully"})
}

func (h *AuthorHandler) GetAuthors(c fiber.Ctx) error {
	queries := c.Queries()

	page, err := strconv.Atoi(queries["page"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid value for paramater page"})
	}

	limit, errr := strconv.Atoi(queries["limit"])

	if errr != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Invalid value for paramater limit"})
	}

	authors, total, errrr := h.Service.GetAuthors(page, limit)

	if errrr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to fetch authors"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Authors fetched successfully", "data": fiber.Map{"authors": authors, "meta": fiber.Map{"total": total, "page": page, "limit": limit}}})
}

func (h *AuthorHandler) GetAuthor(c fiber.Ctx) error {
	authorId := c.Params("id")
	if authorId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "id parameter is missing"})
	}

	author, err := h.Service.GetAuthor(authorId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "author not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Author fetched successfully", "data": fiber.Map{"author": author}})
}

func (h *AuthorHandler) UpdateAuthor(c fiber.Ctx) error {
	authorId := c.Params("id")
	if authorId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "id parameter is missing"})
	}

	var payload types.UpdateAuthorPayload
	if err := c.Bind().Body(&payload); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusForbidden).JSON(validators.FormatValidationError(err))
	}

	err := h.Service.UpdateAuthor(authorId, &payload)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to update author"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Author upated successfully"})
}

func (h *AuthorHandler) GetAuthorBooksByAuthorId(c fiber.Ctx) error {
	return nil
}
