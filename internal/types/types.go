package types

import (
	"encoding/json"
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
	Price         int64   `json:"price"`
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
	Price           int64     `json:"price"`

	Author     *AuthorResponse     `json:"author,omitempty"`
	Categories []*CategoryResponse `json:"categories,omitempty"`
}

type UpdateBookPayload struct {
	Title           *string   `json:"title,omitempty"`
	Description     *string   `json:"description,omitempty"`
	AuthorID        *string   `json:"authorId,omitempty"`
	ISBN            *string   `json:"isbn,omitempty"`
	PublishedYear   *int      `json:"publishedYear,omitempty"`
	CopiesTotal     *int      `json:"copiesTotal,omitempty"`
	CopiesAvailable *int      `json:"copiesAvailable,omitempty"`
	CategoryIds     []*string `json:"categoryIds,omitempty"`
	Price           *int64    `json:"price,omitempty"`
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

type InitiateOrderPayload struct {
	UserID string           `json:"userId"`
	Items  []OrderItemInput `json:"items"`
}

type OrderItemInput struct {
	BookID   string `json:"bookId"`
	Quantity int    `json:"quantity"`
}

type InitiatePaymentPayload struct {
	OrderID          string                  `json:"orderId"`
	PaymentMethod    string                  `json:"paymentMethod"`
	Amount           int64                   `json:"amount"`
	Metadata         string                  `json:"metadata"`
	WebhookProcessed bool                    `json:"webhookProcessed"`
	PaymentGateway   string                  `json:"paymentGateway"`
	Status           constants.PaymentStatus `json:"status"`
	Reference        string                  `json:"reference"`
	Currency         string                  `json:"currency"`
}

type PaystackVerifyResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ID              int64           `json:"id"`
		Domain          string          `json:"domain"`
		Status          string          `json:"status"`
		Reference       string          `json:"reference"`
		ReceiptNumber   *string         `json:"receipt_number"`
		Amount          int64           `json:"amount"`
		Message         *string         `json:"message"`
		GatewayResponse string          `json:"gateway_response"`
		PaidAt          string          `json:"paid_at"`
		CreatedAt       string          `json:"created_at"`
		Channel         string          `json:"channel"`
		Currency        string          `json:"currency"`
		IPAddress       string          `json:"ip_address"`
		Metadata        json.RawMessage `json:"metadata"`
		Log             struct {
			StartTime int64 `json:"start_time"`
			TimeSpent int64 `json:"time_spent"`
			Attempts  int   `json:"attempts"`
			Errors    int   `json:"errors"`
			Success   bool  `json:"success"`
			Mobile    bool  `json:"mobile"`
			Input     []any `json:"input"`
			History   []struct {
				Type    string `json:"type"`
				Message string `json:"message"`
				Time    int    `json:"time"`
			} `json:"history"`
		} `json:"log"`
		Fees          int64            `json:"fees"`
		FeesSplit     *json.RawMessage `json:"fees_split"`
		Authorization struct {
			AuthorizationCode string  `json:"authorization_code"`
			Bin               string  `json:"bin"`
			Last4             string  `json:"last4"`
			ExpMonth          string  `json:"exp_month"`
			ExpYear           string  `json:"exp_year"`
			Channel           string  `json:"channel"`
			CardType          string  `json:"card_type"`
			Bank              string  `json:"bank"`
			CountryCode       string  `json:"country_code"`
			Brand             string  `json:"brand"`
			Reusable          bool    `json:"reusable"`
			Signature         string  `json:"signature"`
			AccountName       *string `json:"account_name"`
		} `json:"authorization"`
		Customer struct {
			ID                       int64           `json:"id"`
			FirstName                *string         `json:"first_name"`
			LastName                 *string         `json:"last_name"`
			Email                    string          `json:"email"`
			CustomerCode             string          `json:"customer_code"`
			Phone                    *string         `json:"phone"`
			Metadata                 json.RawMessage `json:"metadata"`
			RiskAction               string          `json:"risk_action"`
			InternationalFormatPhone *string         `json:"international_format_phone"`
		} `json:"customer"`
		Plan               *json.RawMessage `json:"plan"`
		Split              map[string]any   `json:"split"`
		OrderID            *string          `json:"order_id"`
		PaidAtAlt          string           `json:"paidAt"`
		CreatedAtAlt       string           `json:"createdAt"`
		RequestedAmount    int64            `json:"requested_amount"`
		PosTransactionData *json.RawMessage `json:"pos_transaction_data"`
		Source             *json.RawMessage `json:"source"`
		FeesBreakdown      *json.RawMessage `json:"fees_breakdown"`
		Connect            *json.RawMessage `json:"connect"`
		TransactionDate    string           `json:"transaction_date"`
		PlanObject         map[string]any   `json:"plan_object"`
		Subaccount         map[string]any   `json:"subaccount"`
	} `json:"data"`
}

type PaystackWebhookEvent struct {
	Event string `json:"event"`
	Data  struct {
		Reference string          `json:"reference"`
		Amount    int64           `json:"amount"`
		Currency  string          `json:"currency"`
		ID        int64           `json:"id"`
		Metadata  json.RawMessage `json:"metadata"`
	} `json:"data"`
}

type PaymentMetadata struct {
	OrderID    string `json:"order_id"`
	UserID     string `json:"user_id"`
	OrderItems []struct {
		BookID   string `json:"book_id"`
		Quantity int    `json:"quantity"`
	} `json:"order_items"`
}

type UpdatePaymentPayload struct {
	Status         constants.PaymentStatus `gorm:"type:varchar(20);not null"`
	Reference      string                  `gorm:"not null;uniqueIndex"`
	PaymentGateway string                  `gorm:"not null"`
	PaymentMethod  string                  `gorm:"not null"`
	Currency       string                  `gorm:"not null"`

	Metadata      json.RawMessage `gorm:"type:jsonb"`
	TransactionID *string         `gorm:"uniqueIndex"`

	WebhookProcessed bool `gorm:"not null;default:false"`

	PaidAt *time.Time
}
