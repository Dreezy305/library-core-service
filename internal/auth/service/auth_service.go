package service

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/dreezy305/library-core-service/internal/auth/repository"
	"github.com/dreezy305/library-core-service/internal/config"
	"github.com/dreezy305/library-core-service/internal/mailer"
	"github.com/dreezy305/library-core-service/internal/model"
	"github.com/dreezy305/library-core-service/internal/types"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) RegisterUserService(u *types.UserType) error {
	fmt.Println("In service:", u)
	// Implementation of user registration service goes here
	// check if email exists
	exists, err := s.repo.EmailExists(u.Email)
	if err != nil {
		return err
	}
	if exists {
		fmt.Println(exists, "exists")
		return errors.New("email already in use")
	}
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(u.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user := &model.UserEntity{
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Email:        &u.Email,
		PasswordHash: string(hashedPassword),
		Role:         "member", // default role
	}
	// fmt.Println(hashedPassword)
	fmt.Println("User entity:", user)
	return s.repo.CreateUser(user)
}

func (s *AuthService) LoginUserService(u *types.LoginUserType) (string, error) {
	// Implementation of user login service goes here
	user, err := s.repo.GetUserByEmail(u.Email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(u.Password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	cfg := config.Load().JWTSecret
	jwtSecret := cfg
	claims := jwt.MapClaims{
		"id":        user.ID,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"email":     *user.Email,
		"role":      user.Role,
		"exp":       time.Now().Add(time.Hour * 72).Unix(), // 3 days
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))

}

func (s *AuthService) VerifyEmailService() {
	// Implementation of email verification service goes here
}

func (s *AuthService) ForgotPasswordService(u *types.ForgotPassword) error {
	// Implementation of forgot password service goes here
	user, err := s.repo.GetUserByEmail(u.Email)
	if err != nil {
		return err
	}

	var token string

	for {
		token, err = s.GenerateSecureRandom6DigitToken()
		if err != nil {
			return err
		}

		errrs := s.repo.SaveTokenToDb(token, user)
		if errrs == nil {
			// Save the token to DB
			break
		}
	}

	TokenPayload := &types.ResetTokenHtmlBodyStruct{
		Name:  user.FirstName + " " + user.LastName,
		Token: token,
	}

	htmlBody, errrr := mailer.ForgotPasswordtemplate(TokenPayload)
	if errrr != nil {
		return errrr
	}

	// send email containing token
	SendEmailPayload := &types.SendEmail{
		HtmlBody:     htmlBody,
		Name:         user.FirstName + " " + user.LastName,
		Subject:      "Password reset request 🔐",
		EmailAddress: u.Email,
	}

	mailer.SendEmail(SendEmailPayload)

	return nil
}

func (s *AuthService) ResetPasswordService(u *types.ResetPassword) error {
	// Implementation of reset password service goes here
	_, err := s.repo.GetUserByEmail(u.Email)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) GenerateToken(userID, firstName, lastName, email, role string) (string, error) {

	cfg := config.Load().JWTSecret
	jwtSecret := cfg
	claims := jwt.MapClaims{
		"id":        userID,
		"firstName": firstName,
		"lastName":  lastName,
		"email":     email,
		"role":      role,
		"exp":       time.Now().Add(time.Hour * 72).Unix(), // 3 days
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func (s *AuthService) ValidateToken(tokenString string) (*jwt.Token, error) {
	cfg := config.Load().JWTSecret
	jwtSecret := cfg
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
}

func (s *AuthService) GenerateSecureRandom6DigitToken() (string, error) {
	min := big.NewInt(100000)
	max := big.NewInt(900000)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	result := n.Add(n, min)
	fmt.Println(result)
	return strconv.FormatInt(result.Int64(), 10), nil
}

func (s *AuthService) CompareOldNewPasswordsAndResetToken(payload *types.ResetPassword) error {
	fmt.Println(payload, "reset payload")
	// if payload.Password != payload.ConfirmPassword {
	// 	return fmt.Errorf("Password must be provide")
	// }

	user, err := s.repo.GetUserByEmail(payload.Email)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(payload.Password))
	if err == nil {
		return errors.New("new password must be different from the old password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(payload.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	errr := s.repo.ConfirmResetToken(payload.Token, user.ID)

	if errr != nil {
		fmt.Println(errr, "token error")
		return errors.New("Invalid token provided")
	}

	s.repo.DeleteResetEntity(payload.Token, user.ID)

	return s.repo.ResetPassword(payload, string(hashedPassword))
}
