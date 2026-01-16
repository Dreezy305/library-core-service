package handler

import (
	"fmt"
	"strconv"

	"github.com/dreezy305/library-core-service/internal/loans/service"
	"github.com/dreezy305/library-core-service/internal/types"
	"github.com/dreezy305/library-core-service/internal/validators"
	"github.com/gofiber/fiber/v3"
)

type LoansHandler struct {
	Service *service.LoansService
}

func NewLoansHandler(service *service.LoansService) *LoansHandler {
	return &LoansHandler{Service: service}
}

func (h *LoansHandler) CreateLoan(c fiber.Ctx) error {
	memberId := c.Params("memberId")
	bookId := c.Params("bookId")
	if memberId == "" || bookId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "memberId and bookId are required"})
	}
	var payload types.LoanPayload
	if err := c.Bind().Body(&payload); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusForbidden).JSON(validators.FormatValidationError(err))
	}

	fmt.Println(payload)

	errs := validators.ValidateStruct(payload)
	if errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validators.FormatValidationError(errs))
	}

	err := h.Service.CreateLoan(memberId, bookId, payload)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Loan created successfully"})
}

func (h *LoansHandler) GetLoans(c fiber.Ctx) error {
	queries := c.Queries()

	page, err := strconv.Atoi(queries["page"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid value for paramater page"})
	}

	limit, errr := strconv.Atoi(queries["limit"])

	if errr != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Invalid value for paramater limit"})
	}

	loans, total, errrr := h.Service.GetLoans(page, limit)

	if errrr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to fetch loans"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Loans fetched successfully", "data": fiber.Map{"loans": loans, "meta": fiber.Map{"page": page, "limit": limit, "total": total}}})

}

func (h *LoansHandler) ReturnBook(c fiber.Ctx) error {
	loanId := c.Params("loanId")
	bookId := c.Params("bookId")
	memberId := c.Params("memberId")
	return h.Service.ReturnBook(loanId, bookId, memberId)
}

func (h *LoansHandler) GetMemberLoans(c fiber.Ctx) error {
	memberId := c.Params("memberId")
	if memberId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "memberId is required"})
	}

	member, total, err := h.Service.GetMemberLoans(memberId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to fetch member loans"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Member loans fetched successfully", "data": fiber.Map{"loans": member, "meta": fiber.Map{"total": total}}})
}
