package repository

import "gorm.io/gorm"

type GormBookRepository struct {
	DB *gorm.DB
}

func NewGormBookRepository(db *gorm.DB) *GormBookRepository {
	return &GormBookRepository{DB: db}
}

func (r *GormBookRepository) GetBooks() error {
	return nil
}

func (r *GormBookRepository) GetBook(bookId string) error {
	return nil
}

func (r *GormBookRepository) CreateBook() error {
	return nil
}

func (r *GormBookRepository) UpdateBook() error {
	return nil
}
