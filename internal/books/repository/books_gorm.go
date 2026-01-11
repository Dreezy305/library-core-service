package repository

import (
	"errors"
	"fmt"

	"github.com/dreezy305/library-core-service/internal/model"
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

func (r *GormBookRepository) GetBooks() error {
	return nil
}

func (r *GormBookRepository) GetBook(bookId string) error {
	return nil
}

func (r *GormBookRepository) UpdateBook() error {
	return nil
}
