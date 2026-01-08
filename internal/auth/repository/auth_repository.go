package repository

import (
	"github.com/dreezy305/library-core-service/internal/model"
	"github.com/dreezy305/library-core-service/internal/types"
)

type AuthRepository struct {
	gormRepo *GormAuthRepository
}

func NewAuthRepository(gormRepo *GormAuthRepository) *AuthRepository {
	return &AuthRepository{gormRepo: gormRepo}
}

func (r *AuthRepository) EmailExists(email string) (bool, error) {
	// Implementation to check if email exists in the database goes here
	return r.gormRepo.EmailExists(email)
}

func (r *AuthRepository) CreateUser(u *model.UserEntity) error {
	// Implementation of user creation in the database goes here
	return r.gormRepo.CreateUser(u)
}

func (r *AuthRepository) GetUserByEmail(email string) (*model.UserEntity, error) {
	return r.gormRepo.GetUserByEmail(email)
}

func (r *AuthRepository) SaveTokenToDb(token string, u *model.UserEntity) error {
	return r.gormRepo.SaveTokenToDb(token, u)
}

func (r *AuthRepository) ConfirmResetToken(token string, userId string) error {
	return r.gormRepo.ConfirmResetToken(token, userId)
}

func (r *AuthRepository) ResetPassword(u *types.ResetPassword, hashPassword string) error {
	return r.gormRepo.ResetPassword(u, hashPassword)
}

func (r *AuthRepository) DeleteResetEntity(token string, userId string) error {
	return r.gormRepo.DeleteResetEntity(token, userId)
}
