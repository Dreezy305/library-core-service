package handler

import (
	"fmt"

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
	return nil
}

func (h *LoansHandler) ReturnBook(c fiber.Ctx) error {
	loanId := c.Params("loanId")
	bookId := c.Params("bookId")
	memberId := c.Params("memberId")
	return h.Service.ReturnBook(loanId, bookId, memberId)
}

func (h *LoansHandler) GetMemberLoans(c fiber.Ctx) error {
	memberId := c.Params("memberId")
	return h.Service.GetMemberLoans(memberId)
}
