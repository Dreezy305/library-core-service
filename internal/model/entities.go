package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type UserEntity struct {
	ID           string         `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"` // Primary key field
	FirstName    string         `gorm:"not null"`
	LastName     string         `gorm:"not null"`
	Email        *string        `gorm:"not null;uniqueIndex"` // Unique email field
	MemberNumber sql.NullString `gorm:"uniqueIndex;->"`
	PasswordHash string         `gorm:"not null"`
	Role         string         `gorm:"not null"`
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
}

type BookEntity struct {
	ID              string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Title           string `gorm:"not null"`
	Description     string
	AuthorID        string  `gorm:"not null;index"`
	ISBN            *string `gorm:"uniqueIndex"`
	PublishedYear   int
	CopiesTotal     int `gorm:"not null;check:copies_total >= 0"`
	CopiesAvailable int `gorm:"not null;check:copies_available >= 0"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

type AuthorEntity struct {
	ID        string  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	FirstName string  `gorm:"not null"`
	LastName  string  `gorm:"not null"`
	Email     *string `gorm:"not null;uniqueIndex"` // Unique email field
	Bio       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type LoanEntity struct {
	ID       string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	BookID   string `gorm:"not null;index"`
	MemberID string `gorm:"not null;index"`

	LoanDate   time.Time `gorm:"not null"`
	DueDate    time.Time `gorm:"not null"`
	ReturnedAt *time.Time

	Status string `gorm:"type:varchar(20);not null"` // e.g. active, returned, overdue

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relationships (optional but recommended)
	Book   BookEntity `gorm:"foreignKey:BookID"`
	Member UserEntity `gorm:"foreignKey:MemberID"`
}

type PasswordResetTokenEntity struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID    string    `gorm:"not null;index"`
	Token     string    `gorm:"not null;uniqueIndex"`
	IsUsed    bool      `gorm:"not null;default:false"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User UserEntity `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
