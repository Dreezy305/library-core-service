package repository

import (
	"fmt"

	"github.com/dreezy305/library-core-service/internal/model"
	"github.com/dreezy305/library-core-service/internal/types"
	"gorm.io/gorm"
)

type GormCategoryRepository struct {
	DB *gorm.DB
}

func NewGormCategoryRepository(db *gorm.DB) *GormCategoryRepository {
	return &GormCategoryRepository{DB: db}
}

func (r *GormCategoryRepository) CategoryExists(name string) (bool, error) {
	// Implementation for checking if a category exists
	var count int64
	err := r.DB.Model(&model.CategoryEntity{}).Where("name = ?", name).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *GormCategoryRepository) CreateCategory(c *model.CategoryEntity) error {
	// Implementation for creating a category
	err := r.DB.Create(c).Error
	if err != nil {
		fmt.Println("create category error:", err)
		return err
	}
	return nil
}

func (r *GormCategoryRepository) GetCategories() ([]*types.CategoryResponse, error) {
	// Implementation for retrieving categories
	var categories []model.CategoryEntity

	var response []*types.CategoryResponse

	err := r.DB.Model(&model.CategoryEntity{}).Find(&categories).Error

	for _, category := range categories {
		response = append(response, &types.CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			Slug:        category.Slug,
		})
	}
	return response, err
}

func (r *GormCategoryRepository) DeleteCategory(categoryId string) error {
	// Implementation for deleting a category
	return r.DB.Where("id = ?", categoryId).Delete(&model.CategoryEntity{}).Error
}
