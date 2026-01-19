package service

import (
	"time"

	"github.com/dreezy305/library-core-service/internal/types"
	"github.com/dreezy305/library-core-service/internal/users/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsers(page int, limit int, search *string, startDate *string, endDate *string) ([]*types.UserResponse, int64, error) {
	var startDatePtr *time.Time
	var endDatePtr *time.Time

	if startDate != nil {
		startDateParsed, err := time.Parse("2006-01-02", *startDate)
		if err != nil {
			return nil, 0, err
		}
		startDatePtr = &startDateParsed
	}

	if endDate != nil {
		endDateParsed, err := time.Parse("2006-01-02", *endDate)
		if err != nil {
			return nil, 0, err
		}
		endDatePtr = &endDateParsed
	}
	return s.repo.GetUsers(page, limit, search, startDatePtr, endDatePtr)
}

func (s *UserService) GetUser(userId string) (*types.UserResponse, error) {
	return s.repo.GetUser(userId)
}

func (s *UserService) UpdateUser(userId string, payload types.UpdateUser) error {
	return s.repo.UpdateUser(userId, payload)
}
