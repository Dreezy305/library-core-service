package service

import "github.com/dreezy305/library-core-service/internal/categories/repository"

type CategoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(name string) error {
	return s.repo.CreateCategory(name)
}

func (s *CategoryService) GetCategories() ([]string, error) {
	return s.repo.GetCategories()
}

func (s *CategoryService) DeleteCategory(name string) error {
	return s.repo.DeleteCategory(name)
}
