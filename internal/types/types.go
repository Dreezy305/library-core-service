package types

import (
	"time"

	"github.com/dreezy305/library-core-service/internal/enums"
)

type UserType struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	Role      string `json:"role" validate:"required"`
}

type LoginUserType struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type ForgotPassword struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPassword struct {
	Email    string `json:"email" validate:"required,email"`
	Token    string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
	// ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=password"`
}

type Book struct {
	Title           string
	Description     string
	AuthorID        string
	ISBN            *string
	PublishedYear   int
	CopiesTotal     int
	CopiesAvailable int
}

type AuthorType struct {
	FirstName string
	LastName  string
	Bio       string
}

type Loan struct {
	ID         string
	UserID     uint
	BookID     string
	Status     enums.LoanStatus
	BorrowedAt time.Time
	DueAt      time.Time
	ReturnedAt *time.Time
}

type CreateLoan struct {
	UserID uint
	BookID string
	DueAt  time.Time
}

type ReturnLoan struct {
	LoanID string
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type SendEmail struct {
	HtmlBody     string `json:"htmlBody"`
	Name         string `json:"name"`
	Subject      string `json:"subject"`
	EmailAddress string `json:"emailAddress"`
}

type ResetTokenHtmlBodyStruct struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
}

type UpdateUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type AuthResponse struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
