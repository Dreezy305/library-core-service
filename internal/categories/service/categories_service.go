package service

import (
	"errors"
	"fmt"
	"strings"

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

func (s *CategoryService) CreateCategory(payload types.CategoryPayload) error {
	exists, _ := s.CategoryExists(payload.Name)

	if exists {
		return errors.New("Category already exists")
	}

	fmt.Println(payload, "payload")

	slug := strings.ToLower(strings.ReplaceAll(payload.Name, " ", "-"))
	fmt.Println(slug, "slug")

	cModel := &model.CategoryEntity{
		Name: payload.Name,
		Slug: slug,
	}

	if payload.Description != nil {
		cModel.Description = payload.Description
	}

	fmt.Println(cModel, "model payload")

	return s.repo.CreateCategory(cModel)
}

func (s *CategoryService) GetCategories() ([]*types.CategoryResponse, error) {
	return s.repo.GetCategories()
}

func (s *CategoryService) DeleteCategory(categoryId string) error {
	return s.repo.DeleteCategory(categoryId)
}


