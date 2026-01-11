package service

import (
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

func (s *BookService) CreateBook(b *model.BookEntity) error {
	return s.repo.CreateBook(b)
}

func (s *BookService) GetBooks(page int, limit int) ([]*types.BookResponse, int64, error) {
	return s.repo.GetBooks(page, limit)
}

func (s *BookService) GetBook(bookId string) (*types.BookResponse, error) {
	return s.repo.GetBook(bookId)
}

func (s *BookService) UpdateBook(bookId string, payload *types.BookPayload) error {
	return s.repo.UpdateBook(bookId, payload)
}
