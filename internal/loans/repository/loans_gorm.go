package repository

import (
	"errors"
	"fmt"

	"github.com/dreezy305/library-core-service/internal/model"
	"github.com/dreezy305/library-core-service/internal/types"
	"gorm.io/gorm"
)

type GormLoanRepository struct {
	DB *gorm.DB
}

func NewGormLoanRepository(db *gorm.DB) *GormLoanRepository {
	return &GormLoanRepository{DB: db}
}

func (r *GormLoanRepository) CreateLoan(payload model.LoanEntity) error {
	// Dummy method to illustrate structure
	err := r.DB.Create(&payload).Error
	if err != nil {
		fmt.Println(err)
		return errors.New("Failed to create book")
	}
	return nil
}

func (r *GormLoanRepository) GetLoanByMemberAndBook(memberId string, bookId string) (*model.LoanEntity, error) {
	var loan model.LoanEntity
	err := r.DB.Where("member_id = ? AND book_id = ?", memberId, bookId).Find(&loan).Error
	if err != nil {
		return nil, err
	}
	return &loan, nil
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
