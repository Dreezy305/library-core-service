package service

import (
	"github.com/dreezy305/library-core-service/internal/loans/repository"
	"github.com/dreezy305/library-core-service/internal/types"
	UserRepository "github.com/dreezy305/library-core-service/internal/users/repository"
)

const (
	MinLoanDays = 1
	MaxLoanDays = 14
)

type LoansService struct {
	loansRepo repository.LoansRepository
	userRepo  UserRepository.UserRepository
}

func NewLoansService(loansRepo repository.LoansRepository, userRepo UserRepository.UserRepository) *LoansService {
	return &LoansService{loansRepo: loansRepo, userRepo: userRepo}
}

func (s *LoansService) CreateLoan(memberId string, bookId string, payload types.LoanPayload) error {
	// check if member exists

	return s.loansRepo.CreateLoan(memberId, bookId, payload)
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
