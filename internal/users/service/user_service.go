package service

import (
	"github.com/dreezy305/library-core-service/internal/types"
	"github.com/dreezy305/library-core-service/internal/users/repository"
	"github.com/dreezy305/library-core-service/internal/utils"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsers(page int, limit int, search *string, startDate *string, endDate *string) ([]*types.UserResponse, int64, error) {
	startDatePtr, err := utils.ParseDate(*startDate)

	if err != nil {
		return nil, 0, err
	}

	endDatePtr, err := utils.ParseDate(*endDate)

	if err != nil {
		return nil, 0, err
	}

	return s.repo.GetUsers(page, limit, search, &startDatePtr, &endDatePtr)
}

func (s *UserService) GetUser(userId string) (*types.UserResponse, error) {
	return s.repo.GetUser(userId)
}

func (s *UserService) UpdateUser(userId string, payload types.UpdateUser) error {
	return s.repo.UpdateUser(userId, payload)
}
