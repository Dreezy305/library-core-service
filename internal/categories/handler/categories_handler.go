package handler

import "github.com/dreezy305/library-core-service/internal/categories/service"

type CategoryHandler struct {
	Service service.CategoryService
}

func NewCategoryHandler(s *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{Service: *s}
}

func (h *CategoryHandler) CreateCategory(name string) error {
	return h.Service.CreateCategory(name)
}

func (h *CategoryHandler) GetCategories() ([]string, error) {
	return h.Service.GetCategories()
}

func (h *CategoryHandler) DeleteCategory(name string) error {
	return h.Service.DeleteCategory(name)
}
