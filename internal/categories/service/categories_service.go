package service

import (
	"github.com/dreezy305/library-core-service/internal/categories/repository"
	"github.com/dreezy305/library-core-service/internal/model"
	"github.com/dreezy305/library-core-service/internal/types"
)

type CategoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CategoryExists(name string) (bool, error) {
	return s.repo.CategoryExists(name)
}

func (s *CategoryService) CreateCategory(c *model.CategoryEntity) error {
	return s.repo.CreateCategory(c)
}

func (s *CategoryService) GetCategories() ([]*types.CategoryResponse, error) {
	return s.repo.GetCategories()
}

func (s *CategoryService) DeleteCategory(categoryId string) error {
	return s.repo.DeleteCategory(categoryId)
}
