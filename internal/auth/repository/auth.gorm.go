package repository

import (
	"fmt"
	"time"

	"github.com/dreezy305/library-core-service/internal/model"
	"github.com/dreezy305/library-core-service/internal/types"
	"gorm.io/gorm"
)

type GormAuthRepository struct {
	DB *gorm.DB
}

func NewGormAuthRepository(db *gorm.DB) *GormAuthRepository {
	return &GormAuthRepository{DB: db}
}

func (r *GormAuthRepository) EmailExists(email string) (bool, error) {
	fmt.Println(email)
	var count int64
	// Implementation to check if email exists in the database using GORM goes here
	err := r.DB.Model(&model.UserEntity{}).Where("email = ?", email).Count(&count).Error
	if count > 0 {
		return true, nil
	}
	if err != nil {
		return false, err
	}
	fmt.Println("email:", err)
	return false, nil
}

func (r *GormAuthRepository) CreateUser(u *model.UserEntity) error {
	// Implementation of user creation in the database using GORM goes here
	fmt.Println(u, "entity")
	err := r.DB.Create(u).Error
	if err != nil {
		fmt.Println("Create user error:", err)
		return err
	}
	return nil
}

func (r *GormAuthRepository) LoginUser(email string) (*model.UserEntity, error) {
	// Implementation of user login in the database using GORM goes here
	fmt.Println(email, "entity")
	var user model.UserEntity
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		fmt.Println("Login user error:", err)
		return nil, err
	}
	return &user, nil
}

func (r *GormAuthRepository) GetUserByEmail(email string) (*model.UserEntity, error) {
	var user model.UserEntity
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormAuthRepository) SaveTokenToDb(token string, u *model.UserEntity) error {
	otp := &model.PasswordResetTokenEntity{
		UserID:    u.ID,
		Token:     token,
		IsUsed:    false,
		ExpiresAt: time.Now().Add(time.Minute * 5),
	}
	err := r.DB.Create(otp).Error
	return err
}

func (r *GormAuthRepository) ResetPassword(u *types.ResetPassword, hashPassword string) error {
	err := r.DB.Model(&model.UserEntity{}).Where("email = ?", u.Email).Update("password_hash", hashPassword).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *GormAuthRepository) ConfirmResetToken(token string, userId string) error {
	var otp model.PasswordResetTokenEntity

	err := r.DB.Where("token = ? AND user_id = ?", token, userId).First(&otp).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *GormAuthRepository) DeleteResetEntity(token string, userId string) error {
	var otp model.PasswordResetTokenEntity

	err := r.DB.Where("token = ? AND user_id = ?", token, userId).Delete(&otp).Error
	if err != nil {
		return err
	}

	return nil
}
