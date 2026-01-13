package handler

import (
	"github.com/dreezy305/library-core-service/internal/loans/service"
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
	return h.Service.CreateLoan(memberId, bookId)
}

func (h *LoansHandler) GetLoans(c fiber.Ctx) error {
	return nil
}

func (h *LoansHandler) ReturnBook(loanId string, bookId string, memberId string) error {
	return nil
}

func (h *LoansHandler) GetMemberLoans(memberId string) error {
	return h.Service.GetMemberLoans(memberId)
}
