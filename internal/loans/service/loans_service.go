package service

import (
	"github.com/dreezy305/library-core-service/internal/loans/repository"
	"github.com/dreezy305/library-core-service/internal/types"
)

type LoansService struct {
	loansRepo *repository.LoansRepository
}

func NewLoansService(loansRepo *repository.LoansRepository) *LoansService {
	return &LoansService{loansRepo: loansRepo}
}

func (s *LoansService) CreateLoan(memberId string, bookId string) error {
	return s.loansRepo.CreateLoan(memberId, bookId)
}

func (s *LoansService) GetLoans() ([]*types.LoanResponse, error) {
	return s.loansRepo.GetLoans()
}

func (s *LoansService) ReturnBook(loanId string, bookId string, memberId string) error {
	return s.loansRepo.ReturnBook(loanId, bookId, memberId)
}

func (s *LoansService) GetMemberLoans(memberId string) error {
	return s.loansRepo.GetMemberLoans(memberId)
}
