package repository

import (
	"fmt"

	"github.com/dreezy305/library-core-service/internal/model"
	"gorm.io/gorm"
)

type GormCategoryRepository struct {
	DB *gorm.DB
}

func NewGormCategoryRepository(db *gorm.DB) *GormCategoryRepository {
	return &GormCategoryRepository{DB: db}
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

func (r *GormCategoryRepository) GetCategories() ([]string, error) {
	// Implementation for retrieving categories
	return nil, nil
}

func (r *GormCategoryRepository) DeleteCategory(name string) error {
	// Implementation for deleting a category
	return nil
}
