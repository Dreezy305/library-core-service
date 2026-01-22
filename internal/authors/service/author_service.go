package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/dreezy305/library-core-service/internal/authors/repository"
	"github.com/dreezy305/library-core-service/internal/model"
	"github.com/dreezy305/library-core-service/internal/types"
	"github.com/dreezy305/library-core-service/internal/utils"
)

type AuthorService struct {
	repo repository.AuthorRepository
}

func NewAuthorService(repo repository.AuthorRepository) *AuthorService {
	return &AuthorService{repo: repo}
}

func (s *AuthorService) AuthorExist(email string) (bool, error) {
	return s.repo.AuthorExist(email)
}

func (s *AuthorService) CreateAuthor(payload types.AuthorPayload) error {
	exist, _ := s.AuthorExist(payload.Email)

	if exist {
		return errors.New("Author has already been created")
	}

	dob, err := time.Parse("2006-01-02", payload.DateOfBirth)
	if err != nil {
		return errors.New("dateOfBirth must be in YYYY-MM-DD format")
	}

	u := &model.AuthorEntity{
		FirstName:   payload.FirstName,
		LastName:    payload.LastName,
		Nationality: payload.Nationality,
		DateOfBirth: dob,
		Email:       &payload.Email,
		Bio:         *payload.Bio,
		PenName:     *payload.PenName,
		Website:     *payload.Website,
		Twitter:     *payload.Twitter,
		Facebook:    *payload.Facebook,
		Linkedln:    *payload.Linkedln,
	}

	errr := s.repo.CreateAuthor(u)

	if errr != nil {
		fmt.Println(errr)
		return errors.New("Failed to create author")
	}

	return nil
}

func (s *AuthorService) GetAuthors(page int, limit int, search *string, startDate *string, endDate *string) ([]*types.AuthorResponse, int64, error) {

	var startDatePtr *time.Time
	var endDatePtr *time.Time

	if startDate != nil && *startDate != "" {
		startDateParsed, err := utils.ParseDate(*startDate)
		if err != nil {
			return nil, 0, err
		}
		startDatePtr = &startDateParsed
	}

	if endDate != nil && *endDate != "" {
		endDateParsed, err := utils.ParseDate(*endDate)
		if err != nil {
			return nil, 0, err
		}
		endDatePtr = &endDateParsed
	}

	return s.repo.GetAuthors(page, limit, search, startDatePtr, endDatePtr)
}

func (s *AuthorService) GetAuthor(authorId string) (*types.AuthorResponse, error) {
	return s.repo.GetAuthor(authorId)
}

func (s *AuthorService) UpdateAuthor(authorId string, payload *types.UpdateAuthorPayload) error {
	return s.repo.UpdateAuthor(authorId, payload)
}

func (s *AuthorService) GetAuthorBooksByAuthorId(authorId string) error {
	return nil
}
