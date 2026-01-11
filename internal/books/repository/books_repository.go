package repository

import "github.com/dreezy305/library-core-service/internal/model"

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

func (s *BookRepository) GetBooks() error {
	return s.gormRepo.GetBooks()
}

func (s *BookRepository) GetBook(bookId string) error {
	return s.gormRepo.GetBook(bookId)
}

func (s *BookRepository) UpdateBook() error {
	return s.gormRepo.UpdateBook()
}
