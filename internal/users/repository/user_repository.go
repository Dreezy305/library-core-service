package repository

import (
	"time"

	"github.com/dreezy305/library-core-service/internal/types"
)

type UserRepository struct {
	gormRepo *GormUserRepository
}

func NewUserRepository(gormRepo *GormUserRepository) *UserRepository {
	return &UserRepository{gormRepo: gormRepo}
}

func (s *UserRepository) GetUsers(page int, limit int, search *string, startDate *time.Time, endDate *time.Time) ([]*types.UserResponse, int64, error) {
	return s.gormRepo.GetUsers(page, limit, search, startDate, endDate)
}
func (s *UserRepository) GetUser(userId string) (*types.UserResponse, error) {
	return s.gormRepo.GetUser(userId)
}
func (s *UserRepository) UpdateUser(userId string, payload types.UpdateUser) error {
	return s.gormRepo.UpdateUser(userId, payload)
}
