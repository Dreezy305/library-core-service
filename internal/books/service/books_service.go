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
	if payload == nil {
		return errors.New("no fields provided for update")
	}

	updates := map[string]interface{}{}

	if payload.Title != nil {
		updates["title"] = *payload.Title
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
	if payload.CopiesTotal != nil {
		updates["copies_total"] = *payload.CopiesTotal
	}
	if payload.CopiesAvailable != nil {
		updates["copies_available"] = *payload.CopiesAvailable
	}
	if payload.AuthorID != nil {
		updates["author_id"] = *payload.AuthorID
	}
	if payload.CategoryIds != nil {
		updates["categories"] = payload.CategoryIds
	}

	return s.repo.UpdateBook(bookId, updates)
}

func (s *BookService) DecrementAvailable(bookId string) error {
	return s.repo.DecrementAvailable(bookId)
}

func (s *BookService) IncrementAvailable(bookId string) error {
	return s.repo.IncrementAvailable(bookId)
}
