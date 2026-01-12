package repository

import "gorm.io/gorm"

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

func (r *GormLoanRepository) GetLoans() {}

func (r *GormLoanRepository) ReturnBook() {}

func (r *GormLoanRepository) GetMemberLoans() {}
