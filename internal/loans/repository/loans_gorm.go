package repository

import (
	"github.com/dreezy305/library-core-service/internal/types"
	"gorm.io/gorm"
)

type GormLoanRepository struct {
	DB *gorm.DB
}

func NewGormLoanRepository(db *gorm.DB) *GormLoanRepository {
	return &GormLoanRepository{DB: db}
}

func (r *GormLoanRepository) CreateLoan(memberId string, bookId string) error {
	// Dummy method to illustrate structure
	return nil
}

func (r *GormLoanRepository) GetLoans() ([]*types.LoanResponse, error) {
	return nil, nil
}

func (r *GormLoanRepository) ReturnBook(loanId string, bookId string, memberId string) error {
	return nil
}

func (r *GormLoanRepository) GetMemberLoans(memberId string) error {
	return nil
}
