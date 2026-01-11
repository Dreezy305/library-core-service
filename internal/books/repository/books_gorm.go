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

func (r *GormBookRepository) GetBooks(page int, limit int) ([]*types.BookResponse, int64, error) {
	var total int64
	var books []*model.BookEntity

	if page <= 0 || limit <= 0 {
		page = 1
		limit = 1
	}

	offset := (page - 1) * limit

	err := r.DB.Preload("Author").Model(&model.BookEntity{}).Find(&books).Offset(offset).Limit(limit).Error

	if err != nil {
		return nil, 0, err
	}

	errr := r.DB.Model(&model.BookEntity{}).Count(&total).Error

	if errr != nil {
		return nil, 0, errr
	}

	var response []*types.BookResponse
	for i, book := range books {

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
		})
		books[i] = book
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

func (r *GormBookRepository) UpdateBook(bookId string, payload *types.BookPayload) error {
	updates := map[string]interface{}{}

	if payload.Title != "" {
		updates["title"] = payload.Title
	}
	if payload.Description != nil {
		updates["description"] = *payload.Description
	}
	if payload.ISBN != nil {
		updates["isbn"] = *payload.ISBN
	}
	if payload.PublishedYear != nil {
		updates["published_year"] = *payload.PublishedYear
	}
	if payload.CopiesTotal != 0 {
		updates["copies_total"] = payload.CopiesTotal
	}
	if payload.AuthorID != "" {
		updates["author_id"] = payload.AuthorID
	}

	err := r.DB.Model(&model.BookEntity{}).Where("id = ?", bookId).Updates(updates).Error
	if err != nil {
		return err
	}
	return nil
}
