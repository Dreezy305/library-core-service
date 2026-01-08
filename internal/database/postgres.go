package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/dreezy305/library-core-service/internal/config"
	"github.com/dreezy305/library-core-service/internal/model"
)

func Connect(config *config.AppConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DatabaseURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	db.AutoMigrate(&model.UserEntity{}, &model.BookEntity{}, &model.AuthorEntity{}, &model.LoanEntity{}, &model.PasswordResetTokenEntity{})
	return db, nil
}
