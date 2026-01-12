package repository

import "github.com/dreezy305/library-core-service/internal/model"

type CategoryRepository struct {
	gormRepo *GormCategoryRepository
}

func NewCategoryRepository(gormRepo *GormCategoryRepository) *CategoryRepository {
	return &CategoryRepository{gormRepo: gormRepo}
}

func (r *CategoryRepository) CreateCategory(c *model.CategoryEntity) error {
	return r.gormRepo.CreateCategory(c)
}

func (r *CategoryRepository) GetCategories() ([]string, error) {
	return r.gormRepo.GetCategories()
}

func (r *CategoryRepository) DeleteCategory(name string) error {
	return r.gormRepo.DeleteCategory(name)
}
