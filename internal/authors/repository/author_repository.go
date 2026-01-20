package repository

import (
	"time"

	"github.com/dreezy305/library-core-service/internal/model"
	"github.com/dreezy305/library-core-service/internal/types"
)

type AuthorRepository struct {
	gormRepo *GormAuthorRepository
}

func NewAuthorRepository(gormRepo *GormAuthorRepository) *AuthorRepository {
	return &AuthorRepository{gormRepo: gormRepo}
}

func (r *AuthorRepository) AuthorExist(email string) (bool, error) {
	return r.gormRepo.AuthorExist(email)
}

func (r *AuthorRepository) CreateAuthor(a *model.AuthorEntity) error {
	return r.gormRepo.CreateAuthor(a)
}

func (r *AuthorRepository) GetAuthors(page int, limit int, search *string, startDate *time.Time, endDate *time.Time) ([]*types.AuthorResponse, int64, error) {
	return r.gormRepo.GetAuthors(page, limit, search, startDate, endDate)
}

func (r *AuthorRepository) GetAuthor(authorId string) (*types.AuthorResponse, error) {
	return r.gormRepo.GetAuthor(authorId)
}

func (r *AuthorRepository) UpdateAuthor(authorId string, payload *types.UpdateAuthorPayload) error {
	return r.gormRepo.UpdateAuthor(authorId, payload)
}

func (r *AuthorRepository) GetAuthorBooksByAuthorId(authorId string) error {
	return nil
}
