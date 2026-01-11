package repository

import "gorm.io/gorm"

type GormCategoryRepository struct {
	DB *gorm.DB
}

func NewGormCategoryRepository(db *gorm.DB) *GormCategoryRepository {
	return &GormCategoryRepository{DB: db}
}

// Implement methods for category repository

func (r *GormCategoryRepository) CreateCategory(name string) error {
	// Implementation for creating a category
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
