package repository

import (
	"fmt"
	"time"

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

func (r *GormAuthorRepository) GetAuthors(page int, limit int, search *string, startDate *time.Time, endDate *time.Time) ([]*types.AuthorResponse, int64, error) {
	var total int64
	var authors []*model.AuthorEntity

	if page <= 0 || limit <= 0 {
		page = 1
		limit = 1
	}

	query := r.DB.Model(&model.AuthorEntity{})

	if search != nil {
		likeSearch := fmt.Sprintf("%%%s%%", *search)
		query = query.Where("first_name ILIKE ? OR last_name ILIKE ? OR email ILIKE ? OR pen_name ILIKE ?", likeSearch, likeSearch, likeSearch, likeSearch)
	}

	offset := (page - 1) * limit

	err := query.Find(&authors).Offset(offset).Limit(limit).Error
	if err != nil {
		return nil, 0, err
	}

	if startDate != nil {
		query = query.Where("created_at >= ?", *startDate)
	}

	if endDate != nil {
		query = query.Where("created_at <= ?", *endDate)

	}

	errr := query.Count(&total).Error

	if errr != nil {
		return nil, 0, errr
	}

	var response []*types.AuthorResponse
	for _, author := range authors {
		response = append(response, &types.AuthorResponse{
			ID:          author.ID,
			FirstName:   author.FirstName,
			LastName:    author.LastName,
			Bio:         &author.Bio,
			Email:       *author.Email,
			DateOfBirth: author.DateOfBirth.Format("2006-01-02"),
			Nationality: author.Nationality,
			Website:     &author.Website,
			Twitter:     &author.Twitter,
			Facebook:    &author.Facebook,
			Linkedln:    &author.Linkedln,
			PenName:     &author.PenName,
		})
	}

	return response, total, nil
}

func (r *GormAuthorRepository) GetAuthor(authorId string) (*types.AuthorResponse, error) {
	var author model.AuthorEntity

	err := r.DB.Where("id = ?", authorId).Find(&author).Error

	if err != nil {
		return nil, err
	}

	response := &types.AuthorResponse{
		ID:          author.ID,
		FirstName:   author.FirstName,
		LastName:    author.LastName,
		Bio:         &author.Bio,
		Email:       *author.Email,
		DateOfBirth: author.DateOfBirth.Format("2006-01-02"),
		Nationality: author.Nationality,
		Website:     &author.Website,
		Twitter:     &author.Twitter,
		Facebook:    &author.Facebook,
		Linkedln:    &author.Linkedln,
		PenName:     &author.PenName,
	}
	return response, nil
}

func (r *GormAuthorRepository) UpdateAuthor(authorId string, payload *types.UpdateAuthorPayload) error {
	updates := map[string]interface{}{}

	if payload.FirstName != nil {
		updates["first_name"] = payload.FirstName
	}
	if payload.LastName != nil {
		updates["last_name"] = payload.LastName
	}
	if payload.DateOfBirth != nil {
		updates["date_of_birth"] = payload.DateOfBirth
	}
	if payload.Nationality != nil {
		updates["nationality"] = *payload.Nationality
	}
	if payload.Bio != nil {
		updates["bio"] = *payload.Bio
	}
	if payload.Website != nil {
		updates["website"] = *payload.Website
	}
	if payload.Twitter != nil {
		updates["twitter"] = *payload.Twitter
	}
	if payload.Facebook != nil {
		updates["facebook"] = *payload.Facebook
	}
	if payload.Linkedln != nil {
		updates["linkedln"] = *payload.Linkedln
	}
	if payload.PenName != nil {
		updates["pen_name"] = *payload.PenName
	}

	// Nothing to update
	if len(updates) == 0 {
		return nil
	}

	err := r.DB.Model(&model.AuthorEntity{}).Where("id = ?", authorId).Updates(updates).Error

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r *GormAuthorRepository) GetAuthorBooksByAuthorId(authorId string) error {

	return nil
}
