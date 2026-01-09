package service

import (
	"github.com/dreezy305/library-core-service/internal/authors/repository"
	"github.com/dreezy305/library-core-service/internal/model"
	"github.com/dreezy305/library-core-service/internal/types"
)

type AuthorService struct {
	repo repository.AuthorRepository
}

func NewAuthorService(repo repository.AuthorRepository) *AuthorService {
	return &AuthorService{repo: repo}
}

func (s *AuthorService) AuthorExist(email string) (bool, error) {
	return s.repo.AuthorExist(email)
}

func (s *AuthorService) CreateAuthor(a *model.AuthorEntity) error {
	return s.repo.CreateAuthor(a)
}

func (s *AuthorService) GetAuthors(page int, limit int) ([]*types.AuthorResponse, int64, error) {
	return s.repo.GetAuthors(page, limit)
}

func (s *AuthorService) GetAuthor(authorId string) (*types.AuthorResponse, error) {
	return s.repo.GetAuthor(authorId)
}

func (s *AuthorService) UpdateAuthor(authorId string, payload *types.UpdateAuthorPayload) error {
	return s.repo.UpdateAuthor(authorId, payload)
}

func (s *AuthorService) GetAuthorBooksByAuthorId(authorId string) error {
	return nil
}
