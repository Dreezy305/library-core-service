package repository

import (
	"errors"
	"fmt"

	"github.com/dreezy305/library-core-service/internal/model"
	"github.com/dreezy305/library-core-service/internal/types"
	"gorm.io/gorm"
)

type GormBookRepository struct {
	DB *gorm.DB
}

func NewGormBookRepository(db *gorm.DB) *GormBookRepository {
	return &GormBookRepository{DB: db}
}

func (r *GormBookRepository) BookExist(title string) (bool, error) {
	var count int64
	err := r.DB.Model(&model.BookEntity{}).Where("title = ?", title).Count(&count).Error
	if count > 0 {
		return true, nil
	}
	if err != nil {
		return false, err
	}
	return false, nil
}

func (r *GormBookRepository) CreateBook(b *model.BookEntity) error {
	err := r.DB.Create(b).Error
	if err != nil {
		fmt.Println(err)
		return errors.New("failed to create book")
	}
	return nil
}

func (r *GormBookRepository) GetBooks(page int, limit int, search *string) ([]*types.BookResponse, int64, error) {
	var total int64
	var books []*model.BookEntity

	if page <= 0 || limit <= 0 {
		page = 1
		limit = 1
	}

	offset := (page - 1) * limit

	query := r.DB.Preload("Author").Preload("Categories").Model(&model.BookEntity{})

	if search != nil {
		likeSearch := fmt.Sprintf("%%%s%%", *search)
		query = query.Where("title ILIKE ? OR description ILIKE ? OR isbn ILIKE ?", likeSearch, likeSearch, likeSearch)
	}

	err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&books).Error

	if err != nil {
		return nil, 0, err
	}

	errr := query.Count(&total).Error

	if errr != nil {
		return nil, 0, errr
	}

	var response []*types.BookResponse
	for _, book := range books {
		var categoryResponse []*types.CategoryResponse
		for _, cat := range book.Categories {
			categoryResponse = append(categoryResponse, &types.CategoryResponse{
				ID:          cat.ID,
				Name:        cat.Name,
				Description: cat.Description,
				Slug:        cat.Slug,
				IsActive:    cat.IsActive,
				CreatedAt:   cat.CreatedAt,
			})
		}
		response = append(response, &types.BookResponse{
			ID:              book.ID,
			Title:           book.Title,
			Description:     &book.Description,
			ISBN:            book.ISBN,
			PublishedYear:   &book.PublishedYear,
			CopiesTotal:     book.CopiesTotal,
			AuthorID:        book.AuthorID,
			CopiesAvailable: book.CopiesAvailable,
			CreatedAt:       book.CreatedAt,
			Author: &types.AuthorResponse{
				ID:          book.Author.ID,
				FirstName:   book.Author.FirstName,
				LastName:    book.Author.LastName,
				Email:       *book.Author.Email,
				Nationality: book.Author.Nationality,
				DateOfBirth: book.Author.DateOfBirth.Format("2006-01-02"),
			},
			Categories: categoryResponse,
		})
	}
	return response, total, nil
}

func (r *GormBookRepository) GetBook(bookId string) (*types.BookResponse, error) {
	var book model.BookEntity

	err := r.DB.Preload("Author").Where("id = ?", bookId).Find(&book).Error
	if err != nil {
		return nil, err
	}

	response := &types.BookResponse{
		ID:              book.ID,
		Title:           book.Title,
		Description:     &book.Description,
		ISBN:            book.ISBN,
		PublishedYear:   &book.PublishedYear,
		CopiesTotal:     book.CopiesTotal,
		CopiesAvailable: book.CopiesAvailable,
		CreatedAt:       book.CreatedAt,
		Author: &types.AuthorResponse{
			ID:          book.Author.ID,
			FirstName:   book.Author.FirstName,
			LastName:    book.Author.LastName,
			DateOfBirth: book.Author.DateOfBirth.Format("2006-01-02"),
			Nationality: book.Author.Nationality,
			Email:       *book.Author.Email,
		},
	}

	return response, nil
}

func (r *GormBookRepository) UpdateBook(bookId string, payload map[string]interface{}) error {
	

	err := r.DB.Model(&model.BookEntity{}).Where("id = ?", bookId).Updates(payload).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *GormBookRepository) DecrementAvailable(bookId string) error {
	err := r.DB.Model(&model.BookEntity{}).Where("id = ? AND copies_available > 0", bookId).Update("copies_available", gorm.Expr("copies_available - 1")).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *GormBookRepository) IncrementAvailable(bookId string) error {
	err := r.DB.Model(&model.BookEntity{}).Where("id = ? AND copies_available < copies_total", bookId).Update("copies_available", gorm.Expr("copies_available + 1")).Error
	if err != nil {
		return err
	}
	return nil
}
