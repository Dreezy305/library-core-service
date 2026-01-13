package repository

import "github.com/dreezy305/library-core-service/internal/types"

type LoansRepository struct {
	gormRepo *GormLoanRepository
}

func NewLoansRepository(gormRepo *GormLoanRepository) *LoansRepository {
	return &LoansRepository{gormRepo: gormRepo}
}

func (r *LoansRepository) CreateLoan(memberId string, bookId string, payload types.LoanPayload) error {
	return r.gormRepo.CreateLoan(memberId, bookId, payload)
}

func (r *LoansRepository) GetLoans() ([]*types.LoanResponse, error) {
	return r.gormRepo.GetLoans()
}

func (r *LoansRepository) ReturnBook(loanId string, bookId string, memberId string) error {
	return r.gormRepo.ReturnBook(loanId, bookId, memberId)
}

func (r *LoansRepository) GetMemberLoans(memberId string) error {
	return r.gormRepo.GetMemberLoans(memberId)
}
