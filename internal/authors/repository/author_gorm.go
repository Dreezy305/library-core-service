package repository

import (
	"fmt"

	"github.com/dreezy305/library-core-service/internal/model"
	"github.com/dreezy305/library-core-service/internal/types"
	"gorm.io/gorm"
)

type GormAuthorRepository struct {
	DB *gorm.DB
}

func NewGormAuthorRepository(db *gorm.DB) *GormAuthorRepository {
	return &GormAuthorRepository{DB: db}
}

func (r *GormAuthorRepository) AuthorExist(email string) (bool, error) {
	var count int64
	err := r.DB.Model(&model.AuthorEntity{}).Where("email = ?", email).Count(&count).Error
	if count > 0 {
		return true, nil
	}
	if err != nil {
		return false, err
	}
	return false, nil
}

func (r *GormAuthorRepository) CreateAuthor(a *model.AuthorEntity) error {
	err := r.DB.Create(a).Error
	if err != nil {
		fmt.Println("create author error:", err)
		return err
	}
	return nil
}

func (r *GormAuthorRepository) GetAuthors(page int, limit int) ([]*types.AuthorResponse, int64, error) {
	var total int64
	var authors []*model.AuthorEntity

	if page <= 0 || limit <= 0 {
		page = 1
		limit = 1
	}

	offset := (page - 1) * limit

	err := r.DB.Model(&model.AuthorEntity{}).Find(&authors).Offset(offset).Limit(limit).Error

	if err != nil {
		return nil, 0, err
	}

	errr := r.DB.Model(&model.AuthorEntity{}).Count(&total).Error

	if errr != nil {
		return nil, 0, errr
	}

	var response []*types.AuthorResponse
	for _, author := range authors {
		response = append(response, &types.AuthorResponse{
			ID: author.ID,
			FirstName: author.FirstName,
			LastName:  author.LastName,
			Bio:       &author.Bio,
			Email:     *author.Email,
		})
	}

	return response, total, nil
}

func (r *GormAuthorRepository) GetAuthor(authorId string) error {
	return nil
}

func (r *GormAuthorRepository) UpdateAuthor(payload *types.UpdateUser) error {
	return nil
}

func (r *GormAuthorRepository) GetAuthorBooksByAuthorId(authorId string) error {
	return nil
}
