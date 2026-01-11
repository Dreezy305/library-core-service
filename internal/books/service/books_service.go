package service

import "github.com/dreezy305/library-core-service/internal/books/repository"

type BookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetBooks() error {
	return s.repo.GetBooks()
}

func (s *BookService) GetBook(bookId string) error {
	return s.repo.GetBook(bookId)
}

func (s *BookService) CreateBook() error {
	return s.repo.CreateBook()
}

func (s *BookService) UpdateBook() error {
	return s.repo.UpdateBook()
}
