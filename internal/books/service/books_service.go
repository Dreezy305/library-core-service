package service

import (
	"github.com/dreezy305/library-core-service/internal/books/repository"
	"github.com/dreezy305/library-core-service/internal/model"
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

func (s *BookService) GetBooks() error {
	return s.repo.GetBooks()
}

func (s *BookService) GetBook(bookId string) error {
	return s.repo.GetBook(bookId)
}

func (s *BookService) UpdateBook() error {
	return s.repo.UpdateBook()
}
