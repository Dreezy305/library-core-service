package repository

import (
	"time"

	"github.com/dreezy305/library-core-service/internal/model"
	"github.com/dreezy305/library-core-service/internal/types"
)

type LoansRepository struct {
	gormRepo *GormLoanRepository
}

func NewLoansRepository(gormRepo *GormLoanRepository) *LoansRepository {
	return &LoansRepository{gormRepo: gormRepo}
}

func (r *LoansRepository) CreateLoan(payload model.LoanEntity) error {
	return r.gormRepo.CreateLoan(payload)
}

func (r *LoansRepository) GetLoanByMemberAndBook(memberId string, bookId string) (*model.LoanEntity, error) {
	return r.gormRepo.GetLoanByMemberAndBook(memberId, bookId)
}

func (r *LoansRepository) GetLoans(page int, limit int, search *string, startDate *time.Time, endDate *time.Time) ([]*types.LoanResponse, int, error) {
	return r.gormRepo.GetLoans(page, limit, search, startDate, endDate)
}

func (r *LoansRepository) ReturnBook(loanId string, memberId string, bookId string) error {
	return r.gormRepo.ReturnBook(loanId, memberId, bookId)
}

func (r *LoansRepository) GetMemberLoans(memberId string) ([]*types.LoanResponse, int64, error) {
	return r.gormRepo.GetMemberLoans(memberId)
}
