package repository

import (
	"errors"
	"fmt"
	"time"

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

func (r *GormLoanRepository) GetLoans(page int, limit int, search *string, startDate *time.Time, endDate *time.Time) ([]*types.LoanResponse, int, error) {
	var loans []*model.LoanEntity
	var response []*types.LoanResponse
	if page <= 0 || limit <= 0 {
		page = 1
		limit = 1
	}

	query := r.DB.Model(&model.LoanEntity{})

	offset := (page - 1) * limit

	if search != nil && *search != "" {
		likeSearch := fmt.Sprintf("%%%s%%", *search)
		query = query.Where("first_name ILIKE ? OR last_name ILIKE ? OR email ILIKE ?", likeSearch, likeSearch, likeSearch)
	}

	fmt.Println(startDate, "start date")
	fmt.Println(endDate, "end date")

	if startDate != nil {
		query = query.Where("created_at >= ?", *startDate)
	}

	if endDate != nil {
		query = query.Where("created_at <= ?", *endDate)

	}

	err := query.Find(&loans).Limit(limit).Offset(offset).Error
	if err != nil {
		return nil, 0, err
	}
	var total int64
	errr := query.Count(&total).Error

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
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
		})
	}

	return response, int(total), nil
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

func (r *GormLoanRepository) ReturnBook(loanId string, memberId string, bookId string) error {

	err := r.DB.Model(&model.LoanEntity{}).Where("id = ? AND member_id = ? AND book_id = ?", loanId, memberId, bookId).Update("status", "returned").Error
	if err != nil {
		return err
	}

	return nil
}
