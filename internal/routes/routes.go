package routes

import (
	AuthHandler "github.com/dreezy305/library-core-service/internal/auth/handler"
	AuthRepository "github.com/dreezy305/library-core-service/internal/auth/repository"
	AuthService "github.com/dreezy305/library-core-service/internal/auth/service"
	AuthorHandler "github.com/dreezy305/library-core-service/internal/authors/handler"
	AuthorRepository "github.com/dreezy305/library-core-service/internal/authors/repository"
	AuthorService "github.com/dreezy305/library-core-service/internal/authors/service"
	BookHandler "github.com/dreezy305/library-core-service/internal/books/handler"
	BookRepository "github.com/dreezy305/library-core-service/internal/books/repository"
	BookService "github.com/dreezy305/library-core-service/internal/books/service"
	CategoryHandler "github.com/dreezy305/library-core-service/internal/categories/handler"
	CategoryRepository "github.com/dreezy305/library-core-service/internal/categories/repository"
	CategoryService "github.com/dreezy305/library-core-service/internal/categories/service"
	LoansHandler "github.com/dreezy305/library-core-service/internal/loans/handler"
	LoansRepository "github.com/dreezy305/library-core-service/internal/loans/repository"
	LoansService "github.com/dreezy305/library-core-service/internal/loans/service"
	"github.com/dreezy305/library-core-service/internal/middleware"
	UserHandler "github.com/dreezy305/library-core-service/internal/users/handler"
	UserRepository "github.com/dreezy305/library-core-service/internal/users/repository"
	UserService "github.com/dreezy305/library-core-service/internal/users/service"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"gorm.io/gorm"
)

var gormRepo *AuthRepository.GormAuthRepository
var userGormRepo *UserRepository.GormUserRepository

// health check route
func HealthCheckRoute(app fiber.Router) {
	app.Get("/health", healthcheck.New())
}

// USER ROUTES
func AuthRoutes(app fiber.Router, db *gorm.DB) {
	gormRepo = AuthRepository.NewGormAuthRepository(db)
	authRepo := AuthRepository.NewAuthRepository(gormRepo)
	authService := AuthService.NewAuthService(*authRepo)
	authHandler := AuthHandler.NewAuthHandler(authService)

	// scripts.SeedUsers(authService, 100)

	// Define user-related routes here
	authGroup := app.Group("/auth")
	authGroup.Post("/register", authHandler.RegisterUserHandler)
	authGroup.Post("/login", authHandler.LoginUserHandler)
	authGroup.Post("/verify-email", authHandler.VerifyEmailHandler)
	authGroup.Post("/forgot-password", authHandler.ForgotPasswordHandler)
	authGroup.Post("/reset-password", authHandler.ResetPasswordHandler)
}

func UserRoutes(app fiber.Router, db *gorm.DB) {
	userGormRepo = UserRepository.NewGormUserRepository(db)
	userRepo := UserRepository.NewUserRepository(userGormRepo)
	userService := UserService.NewUserService(*userRepo)
	userHandler := UserHandler.NewUserHandler(userService)

	// Define user-related routes here
	userGroup := app.Group("/users", middleware.JWTProtected())
	userGroup.Get("/", userHandler.GetUsers)
	userGroup.Get("/me/:id", userHandler.GetUser)
	userGroup.Put("/update/:id", userHandler.UpdateUser)
}

// AUTHOR ROUTES
func AuthorRoutes(app fiber.Router, db *gorm.DB) {
	authorGormRepo := AuthorRepository.NewGormAuthorRepository(db)
	authorRepo := AuthorRepository.NewAuthorRepository(authorGormRepo)
	authorService := AuthorService.NewAuthorService(*authorRepo)
	authorHandler := AuthorHandler.NewAuthorHandler(authorService)

	// Define author-related routes here
	authorGroup := app.Group("/authors", middleware.JWTProtected())
	authorGroup.Get("/", authorHandler.GetAuthors)
	authorGroup.Get("/:id", authorHandler.GetAuthor)
	authorGroup.Post("/create", authorHandler.CreateAuthor)
	authorGroup.Put("/:id", authorHandler.UpdateAuthor)
}

// BOOK ROUTES
func BookRoutes(app fiber.Router, db *gorm.DB) {
	bookGormRepo := BookRepository.NewGormBookRepository(db)
	bookRepo := BookRepository.NewBookRepository(bookGormRepo)
	bookService := BookService.NewBookService(*bookRepo)
	bookHandler := BookHandler.NewBookHandler(bookService)

	// Define book-related routes here
	bookGroup := app.Group("/books", middleware.JWTProtected())
	bookGroup.Get("/", bookHandler.GetBooks)
	bookGroup.Get("/:id", bookHandler.GetBook)
	bookGroup.Post("/create", bookHandler.CreateBook)
	bookGroup.Put("/:id", bookHandler.UpdateBook)

}

// BOOK CATEGORIES ROUTES
func BookCategoryRoutes(app fiber.Router, db *gorm.DB) {
	categoryGormRepo := CategoryRepository.NewGormCategoryRepository(db)
	categoryRepo := CategoryRepository.NewCategoryRepository(categoryGormRepo)
	categoryService := CategoryService.NewCategoryService(*categoryRepo)
	categoryHandler := CategoryHandler.NewCategoryHandler(categoryService)

	// Define book category-related routes here
	categoryGroup := app.Group("/categories", middleware.JWTProtected())
	categoryGroup.Get("/", categoryHandler.GetCategories)
	categoryGroup.Post("/create", categoryHandler.CreateCategory)
	categoryGroup.Delete("/delete/:id", categoryHandler.DeleteCategory)
}

// LOAN ROUTES
func LoanRoutes(app fiber.Router, db *gorm.DB) {
	userGormRepo = UserRepository.NewGormUserRepository(db)
	userRepo := UserRepository.NewUserRepository(userGormRepo)

	bookGormRepo := BookRepository.NewGormBookRepository(db)
	bookRepo := BookRepository.NewBookRepository(bookGormRepo)

	loanGormRepo := LoansRepository.NewGormLoanRepository(db)
	loanRepo := LoansRepository.NewLoansRepository(loanGormRepo)
	loanService := LoansService.NewLoansService(*loanRepo, *userRepo, *bookRepo)
	loanHandler := LoansHandler.NewLoansHandler(loanService)

	// Define loan-related routes here
	loanGroup := app.Group("/loans", middleware.JWTProtected())
	loanGroup.Get("/", loanHandler.GetLoans)
	loanGroup.Get("/member/:memberId", loanHandler.GetMemberLoans)
	loanGroup.Post("/:memberId/:bookId", loanHandler.CreateLoan)
	loanGroup.Put("/:loanId/:memberId/:bookId", loanHandler.ReturnBook)
}
