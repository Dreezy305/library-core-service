package repository

import (
	"github.com/dreezy305/library-core-service/internal/model"
	"github.com/dreezy305/library-core-service/internal/types"
)

type AuthorRepository struct {
	gormRepo *GormAuthorRepository
}

func NewAuthorRepository(gormRepo *GormAuthorRepository) *AuthorRepository {
	return &AuthorRepository{gormRepo: gormRepo}
}

func (r *AuthorRepository) CreateAuthor(a *model.AuthorEntity) error {
	return r.gormRepo.CreateAuthor(a)
}

func (r *AuthorRepository) GetAuthors(page int, limit int) error {
	return nil
}

func (r *AuthorRepository) GetAuthor(authorId string) error {
	return nil
}

func (r *AuthorRepository) UpdateAuthor(payload *types.UpdateUser) error {
	return nil
}

func (r *AuthorRepository) GetAuthorBooksByAuthorId(authorId string) error {
	return nil
}
