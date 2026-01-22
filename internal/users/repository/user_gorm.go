package repository

import (
	"fmt"
	"time"

	"github.com/dreezy305/library-core-service/internal/model"
	"github.com/dreezy305/library-core-service/internal/types"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	DB *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{DB: db}
}

func (r *GormUserRepository) GetUsers(page int, limit int, search *string, startDate *time.Time, endDate *time.Time) ([]*types.UserResponse, int64, error) {
	var total int64
	var users []*model.UserEntity

	if page <= 0 || limit <= 0 {
		page = 1
		limit = 1
	}

	offset := (page - 1) * limit

	query := r.DB.Model(&model.UserEntity{})

	if search != nil && *search != "" {
		likeSearch := fmt.Sprintf("%%%s%%", *search)
		query = query.Where("first_name ILIKE ? OR last_name ILIKE ? OR email ILIKE ?", likeSearch, likeSearch, likeSearch)
	}

	fmt.Println(startDate, "start date")
	fmt.Println(endDate, "end date")

	if startDate != nil {
		query = query.Where("created_at >= ?", *startDate)
	}

	if endDate != nil {
		query = query.Where("created_at <= ?", *endDate)

	}

	err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&users).Error

	if err != nil {
		return nil, 0, err
	}

	errr := query.Count(&total).Error

	if errr != nil {
		return nil, 0, errr
	}

	fmt.Println(users)

	var response []*types.UserResponse
	for _, v := range users {
		fmt.Println(v, "vvv")
		response = append(response, &types.UserResponse{
			ID:           v.ID,
			FirstName:    v.FirstName,
			LastName:     v.LastName,
			Email:        *v.Email,
			MemberNumber: v.MemberNumber,
			CreatedAt:    v.CreatedAt,
		})
	}

	return response, total, nil
}

func (r *GormUserRepository) GetUser(userId string) (*types.UserResponse, error) {
	var user model.UserEntity

	err := r.DB.Where("id = ?", userId).Find(&user).Error

	if err != nil {
		return nil, err
	}

	response := &types.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
		Email:     *user.Email,
	}

	return response, nil
}

func (r *GormUserRepository) UpdateUser(userId string, payload types.UpdateUser) error {
	updates := map[string]interface{}{}

	if payload.FirstName != "" {
		updates["first_name"] = payload.FirstName
	}

	if payload.LastName != "" {
		updates["last_name"] = payload.LastName
	}

	if len(updates) == 0 {
		return nil
	}

	err := r.DB.Model(&model.UserEntity{}).Where("id = ?", userId).Updates(updates).Error

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
