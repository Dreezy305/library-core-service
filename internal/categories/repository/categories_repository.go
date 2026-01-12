package repository

import (
	"github.com/dreezy305/library-core-service/internal/model"
	"github.com/dreezy305/library-core-service/internal/types"
)

type CategoryRepository struct {
	gormRepo *GormCategoryRepository
}

func NewCategoryRepository(gormRepo *GormCategoryRepository) *CategoryRepository {
	return &CategoryRepository{gormRepo: gormRepo}
}

func (r *CategoryRepository) CategoryExists(name string) (bool, error) {
	return r.gormRepo.CategoryExists(name)
}

func (r *CategoryRepository) CreateCategory(c *model.CategoryEntity) error {
	return r.gormRepo.CreateCategory(c)
}

func (r *CategoryRepository) GetCategories() ([]*types.CategoryResponse, error) {
	return r.gormRepo.GetCategories()
}

func (r *CategoryRepository) DeleteCategory(categoryId string) error {
	return r.gormRepo.DeleteCategory(categoryId)
}
