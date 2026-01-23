package service

import (
	"errors"

	"github.com/dreezy305/library-core-service/internal/books/repository"
	"github.com/dreezy305/library-core-service/internal/model"
	"github.com/dreezy305/library-core-service/internal/types"
)

type BookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) BookExists(title string) (bool, error) {
	return s.repo.BookExists(title)
}

func (s *BookService) CreateBook(payload types.BookPayload) error {
	exist, _ := s.BookExists(payload.Title)

	if exist {
		return errors.New("Book has already been created")
	}

	b := &model.BookEntity{
		Title:       payload.Title,
		ISBN:        payload.ISBN,
		CopiesTotal: payload.CopiesTotal,
		AuthorID:    payload.AuthorID,
	}

	if payload.Description != nil {
		b.Description = *payload.Description
	}

	if payload.PublishedYear != nil {
		b.PublishedYear = *payload.PublishedYear
	}
	return s.repo.CreateBook(b)
}

func (s *BookService) GetBooks(page int, limit int, search *string) ([]*types.BookResponse, int64, error) {
	return s.repo.GetBooks(page, limit, search)
}

func (s *BookService) GetBook(bookId string) (*types.BookResponse, error) {
	return s.repo.GetBook(bookId)
}

func (s *BookService) UpdateBook(bookId string, payload *types.UpdateBookPayload) error {
	return s.repo.UpdateBook(bookId, payload)
}

func (s *BookService) DecrementAvailable(bookId string) error {
	return s.repo.DecrementAvailable(bookId)
}

func (s *BookService) IncrementAvailable(bookId string) error {
	return s.repo.IncrementAvailable(bookId)
}
