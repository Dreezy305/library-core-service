package repository

import (
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

func (r *GormAuthorRepository) CreateAuthor(a *model.AuthorEntity) (*types.AuthResponse, error) {
	var author types.AuthResponse

	return &author, nil
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
