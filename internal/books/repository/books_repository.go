package repository

import (
	"github.com/dreezy305/library-core-service/internal/model"
	"github.com/dreezy305/library-core-service/internal/types"
)

type BookRepository struct {
	gormRepo *GormBookRepository
}

func NewBookRepository(gormRepo *GormBookRepository) *BookRepository {
	return &BookRepository{gormRepo: gormRepo}
}

func (s *BookRepository) BookExists(title string) (bool, error) {
	return s.gormRepo.BookExist(title)
}

func (s *BookRepository) CreateBook(b *model.BookEntity) error {
	return s.gormRepo.CreateBook(b)
}

func (s *BookRepository) GetBooks(page int, limit int, search *string) ([]*types.BookResponse, int64, error) {
	return s.gormRepo.GetBooks(page, limit, search)
}

func (s *BookRepository) GetBook(bookId string) (*types.BookResponse, error) {
	return s.gormRepo.GetBook(bookId)
}

func (s *BookRepository) UpdateBook(bookId string, payload *types.UpdateBookPayload) error {
	updates := &types.UpdateBookPayload{}

	if payload.Title != nil {
		updates.Title = payload.Title
	}
	if payload.Description != nil {
		updates.Description = payload.Description
	}
	if payload.ISBN != nil {
		updates.ISBN = payload.ISBN
	}
	if payload.PublishedYear != nil {
		updates.PublishedYear = payload.PublishedYear
	}
	if payload.CopiesTotal != nil {
		updates.CopiesTotal = payload.CopiesTotal
	}
	if payload.CopiesAvailable != nil {
		updates.CopiesAvailable = payload.CopiesAvailable
	}
	if payload.AuthorID != nil {
		updates.AuthorID = payload.AuthorID
	}
	return s.gormRepo.UpdateBook(bookId, updates)
}

func (s *BookRepository) DecrementAvailable(bookId string) error {
	return s.gormRepo.DecrementAvailable(bookId)
}

func (s *BookRepository) IncrementAvailable(bookId string) error {
	return s.gormRepo.IncrementAvailable(bookId)
}
