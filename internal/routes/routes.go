package routes

import (
	AuthHandler "github.com/dreezy305/library-core-service/internal/auth/handler"
	AuthRepository "github.com/dreezy305/library-core-service/internal/auth/repository"
	AuthService "github.com/dreezy305/library-core-service/internal/auth/service"
	UserHandler "github.com/dreezy305/library-core-service/internal/users/handler"
	UserRepository "github.com/dreezy305/library-core-service/internal/users/repository"
	UserService "github.com/dreezy305/library-core-service/internal/users/service"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"gorm.io/gorm"
)

var authHandler *AuthHandler.AuthHandler
var authService *AuthService.AuthService
var authRepository *AuthRepository.AuthRepository
var gormRepo *AuthRepository.GormAuthRepository

var userHandler *UserHandler.UserHandler
var userService *UserService.UserService
var serviceRepository *UserRepository.UserRepository
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
	userGroup := app.Group("/users")
	userGroup.Get("/", userHandler.GetUsers)
	userGroup.Get("/me/:id", userHandler.GetUser)
	userGroup.Put("/update/:id", userHandler.UpdateUser)
}

// AUTHOR ROUTES
func AuthorRoutes(app fiber.Router, db *gorm.DB) {
	// Define author-related routes here
	// authorGroup := app.Group("/authors")
	// authorGroup.Get("/")
	// authorGroup.Get("/:id")
	// authorGroup.Post("/")
	// authorGroup.Put("/:id")
	// authorGroup.Delete("/:id")
}

// BOOK ROUTES
func BookRoutes(app fiber.Router, db *gorm.DB) {
	// Define book-related routes here
	// bookGroup := app.Group("/books")
	// bookGroup.Get("/")
	// bookGroup.Get("/:id")
	// bookGroup.Post("/books")
	// bookGroup.Put("/:id")
	// bookGroup.Delete("/:id")
	// bookGroup.Post("/:id/borrow")
	// bookGroup.Post("/:id/return")
	// bookGroup.Post("/:id/review")
}

// LOAN ROUTES
func LoanRoutes(app fiber.Router, db *gorm.DB) {
	// Define loan-related routes here
	// loanGroup := app.Group("/loans")
	// loanGroup.Get("/")
	// loanGroup.Get("/members/:id/loans")
	// loanGroup.Post("/")
	// loanGroup.Put("/:id")
	// loanGroup.Delete("/:id")
}

// BOOK CATEGORIES ROUTES
func BookCategoryRoutes(app fiber.Router) {
	// Define book category-related routes here
	// categoryGroup := app.Group("/categories")
	// categoryGroup.Get("/")
	// categoryGroup.Get("/:id")
	// categoryGroup.Post("/")
	// categoryGroup.Put("/:id")
	// categoryGroup.Delete("/:id")
}
