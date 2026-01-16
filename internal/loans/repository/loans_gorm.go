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

func (r *GormLoanRepository) GetLoans(page int, limit int) ([]*types.LoanResponse, int, error) {
	var loans []*model.LoanEntity
	var response []*types.LoanResponse
	if page <= 0 || limit <= 0 {
		page = 1
		limit = 1
	}

	offset := (page - 1) * limit

	err := r.DB.Model(&model.LoanEntity{}).Find(&loans).Limit(limit).Offset(offset).Error
	if err != nil {
		return nil, 0, err
	}
	var total int64
	errr := r.DB.Model(&model.LoanEntity{}).Count(&total).Error

	if errr != nil {
		return nil, 0, errr
	}

	for _, v := range loans {
		response = append(response, &types.LoanResponse{
			ID:         v.ID,
			MemberID:   v.MemberID,
			BookID:     v.BookID,
			LoanDate:   v.LoanDate,
			DueDate:    v.DueDate,
			ReturnedAt: v.ReturnedAt,
			Status:     v.Status,
		})
	}

	return response, int(total), nil
}

func (r *GormLoanRepository) ReturnBook(loanId string, bookId string, memberId string) error {
	return nil
}

func (r *GormLoanRepository) GetMemberLoans(memberId string) ([]*types.LoanResponse, int64, error) {
	var loans []*model.LoanEntity
	var response []*types.LoanResponse

	err := r.DB.Preload("Member").Preload("Book").Where("member_id = ?", memberId).Find(&loans).Error
	if err != nil {
		return nil, 0, err
	}

	var total int64
	errr := r.DB.Model(&model.LoanEntity{}).Count(&total).Error

	if errr != nil {
		return nil, 0, errr
	}

	for _, v := range loans {
		response = append(response, &types.LoanResponse{
			ID:         v.ID,
			MemberID:   v.MemberID,
			BookID:     v.BookID,
			LoanDate:   v.LoanDate,
			DueDate:    v.DueDate,
			ReturnedAt: v.ReturnedAt,
			Status:     v.Status,
			Member: &types.UserResponse{
				ID:        v.Member.ID,
				FirstName: v.Member.FirstName,
				LastName:  v.Member.LastName,
				Email:     *v.Member.Email,
			},
			Book: &types.BookResponse{
				ID:              v.Book.ID,
				Title:           v.Book.Title,
				Description:     &v.Book.Description,
				AuthorID:        v.Book.AuthorID,
				ISBN:            v.Book.ISBN,
				PublishedYear:   &v.Book.PublishedYear,
				CopiesTotal:     v.Book.CopiesTotal,
				CopiesAvailable: v.Book.CopiesAvailable,
			},
		})
	}

	return response, total, nil
}
