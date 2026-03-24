package service

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"errors"

	"gorm.io/gorm"

	BookRepository "github.com/dreezy305/library-core-service/internal/books/repository"
	OrderRepository "github.com/dreezy305/library-core-service/internal/orders/repository"
	UserRepository "github.com/dreezy305/library-core-service/internal/users/repository"
)

type PaystackWebhookService struct {
	DB             *gorm.DB
	bookRepo       *BookRepository.BookRepository
	userRepo       *UserRepository.UserRepository
	orderRepo      *OrderRepository.OrderRepository
	paystackSecret string
}

func NewPaystackWebhookService(orderRepo *OrderRepository.OrderRepository, bookRepo *BookRepository.BookRepository, userRepo *UserRepository.UserRepository, db *gorm.DB, paystackSecret string) *PaystackWebhookService {
	return &PaystackWebhookService{orderRepo: orderRepo, bookRepo: bookRepo, userRepo: userRepo, DB: db, paystackSecret: paystackSecret}
}

func (s *PaystackWebhookService) VerifySignature(body []byte, signature string) error {
	mac := hmac.New(sha512.New, []byte(s.paystackSecret))
	mac.Write(body)
	expected := hex.EncodeToString(mac.Sum(nil))
	if !hmac.Equal([]byte(expected), []byte(signature)) {
		return errors.New("invalid signature")
	}
	return nil
}

func (s *PaystackWebhookService) paystack() error {
	return nil
}
