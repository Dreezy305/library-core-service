package types

import (
	"time"

	"github.com/dreezy305/library-core-service/internal/constants"
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
	Status     constants.LoanStatus
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
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	CreatedAt    time.Time `json:"createdAt"`
	MemberNumber string    `json:"memberNumber"`
}

type UpdateUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type AuthorPayload struct {
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	DateOfBirth string  `json:"dateOfBirth"`
	Email       string  `json:"email"`
	Nationality string  `json:"nationality"`
	Bio         *string `json:"bio,omitempty"`
	Website     *string `json:"website,omitempty"`
	Twitter     *string `json:"twitter,omitempty"`
	Facebook    *string `json:"facebook,omitempty"`
	Linkedln    *string `json:"linkedln,omitempty"`
	PenName     *string `json:"penName,omitempty"`
}

type AuthorResponse struct {
	ID          string  `json:"id"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	DateOfBirth string  `json:"dateOfBirth"`
	Email       string  `json:"email"`
	Nationality string  `json:"nationality"`
	Bio         *string `json:"bio,omitempty"`
	Website     *string `json:"website,omitempty"`
	Twitter     *string `json:"twitter,omitempty"`
	Facebook    *string `json:"facebook,omitempty"`
	Linkedln    *string `json:"linkedln,omitempty"`
	PenName     *string `json:"penName,omitempty"`
}

type UpdateAuthorPayload struct {
	FirstName   *string `json:"firstName"`
	LastName    *string `json:"lastName"`
	DateOfBirth *string `json:"dateOfBirth"`
	Nationality *string `json:"nationality"`
	Bio         *string `json:"bio,omitempty"`
	Website     *string `json:"website,omitempty"`
	Twitter     *string `json:"twitter,omitempty"`
	Facebook    *string `json:"facebook,omitempty"`
	Linkedln    *string `json:"linkedln,omitempty"`
	PenName     *string `json:"penName,omitempty"`
}

type BookPayload struct {
	Title         string  `json:"title"`
	Description   *string `json:"description,omitempty"`
	AuthorID      string  `json:"authorId"`
	ISBN          *string `json:"isbn,omitempty"`
	PublishedYear *int    `json:"publishedYear,omitempty"`
	CopiesTotal   int     `json:"copiesTotal"`
}

type BookResponse struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Description     *string   `json:"description,omitempty"`
	AuthorID        string    `json:"authorId"`
	ISBN            *string   `json:"isbn,omitempty"`
	PublishedYear   *int      `json:"publishedYear,omitempty"`
	CopiesTotal     int       `json:"copiesTotal"`
	CopiesAvailable int       `json:"copiesAvailable"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`

	Author     *AuthorResponse     `json:"author,omitempty"`
	Categories []*CategoryResponse `json:"categories,omitempty"`
}

type UpdateBookPayload struct {
	Title           *string `json:"title,omitempty"`
	Description     *string `json:"description,omitempty"`
	AuthorID        *string `json:"authorId,omitempty"`
	ISBN            *string `json:"isbn,omitempty"`
	PublishedYear   *int    `json:"publishedYear,omitempty"`
	CopiesTotal     *int    `json:"copiesTotal,omitempty"`
	CopiesAvailable *int    `json:"copiesAvailable,omitempty"`
}

type CategoryPayload struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type CategoryResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	Slug        string    `json:"slug"`
	IsActive    bool      `json:"isActive"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type UpdateCategoryPayload struct {
	Name *string `json:"name,omitempty"`
}

type LoanResponse struct {
	ID         string        `json:"id"`
	MemberID   string        `json:"memberId"`
	BookID     string        `json:"bookId"`
	Status     string        `json:"status"`
	LoanDate   time.Time     `json:"loanDate"`
	DueDate    time.Time     `json:"dueDate"`
	ReturnedAt *time.Time    `json:"returnedAt,omitempty"`
	Member     *UserResponse `json:"member,omitempty"`
	Book       *BookResponse `json:"book,omitempty"`
	CreatedAt  time.Time     `json:"createdAt"`
	UpdatedAt  time.Time     `json:"updatedAt"`
}

type LoanPayload struct {
	DurationInDays int `json:"duration"`
}

type MemberLoansResponse struct {
	Loans []*LoanResponse `json:"loans"`
}
