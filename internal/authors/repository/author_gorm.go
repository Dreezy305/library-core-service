package repository

import (
	"fmt"

	"github.com/dreezy305/library-core-service/internal/model"
	"github.com/dreezy305/library-core-service/internal/types"
	"gorm.io/gorm"
)

type GormAuthorRepository struct {
	DB *gorm.DB
}

func NewGormAuthorRepository(db *gorm.DB) *GormAuthorRepository {
	return &GormAuthorRepository{DB: db}
}

func (r *GormAuthorRepository) CreateAuthor(a *model.AuthorEntity) error {
	var author types.AuthorPayload
	err := r.DB.Create(author).Error
	if err != nil {
		fmt.Println("create author error:", err)
		return err
	}
	return nil
}

func (r *GormAuthorRepository) GetAuthors(page int, limit int) error {
	return nil
}

func (r *GormAuthorRepository) GetAuthor(authorId string) error {
	return nil
}

func (r *GormAuthorRepository) UpdateAuthor(payload *types.UpdateUser) error {
	return nil
}

func (r *GormAuthorRepository) GetAuthorBooksByAuthorId(authorId string) error {
	return nil
}
