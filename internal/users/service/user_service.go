package service

import (
	"github.com/dreezy305/library-core-service/internal/types"
	"github.com/dreezy305/library-core-service/internal/users/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsers(page int, limit int) ([]*types.UserResponse, int64, error) {
	return s.repo.GetUsers(page, limit)
}

func (s *UserService) GetUser(userId string) (*types.UserResponse, error) {
	return s.repo.GetUser(userId)
}

func (s *UserService) UpdateUser(userId string, payload types.UpdateUser) error {
	return s.repo.UpdateUser(userId, payload)
}
