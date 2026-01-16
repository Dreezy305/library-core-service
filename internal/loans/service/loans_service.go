package service

import (
	"errors"
	"fmt"
	"time"

	BookRepository "github.com/dreezy305/library-core-service/internal/books/repository"
	"github.com/dreezy305/library-core-service/internal/constants"
	"github.com/dreezy305/library-core-service/internal/loans/repository"
	"github.com/dreezy305/library-core-service/internal/model"
	"github.com/dreezy305/library-core-service/internal/types"
	UserRepository "github.com/dreezy305/library-core-service/internal/users/repository"
)

const (
	MinLoanDays = 1
	MaxLoanDays = 14
)

type LoansService struct {
	loansRepo repository.LoansRepository
	userRepo  UserRepository.UserRepository
	bookRepo  BookRepository.BookRepository
}

func NewLoansService(loansRepo repository.LoansRepository, userRepo UserRepository.UserRepository, bookRepo BookRepository.BookRepository) *LoansService {
	return &LoansService{loansRepo: loansRepo, userRepo: userRepo, bookRepo: bookRepo}
}

func (s *LoansService) CreateLoan(memberId string, bookId string, payload types.LoanPayload) error {
	// check if member exists
	_, err := s.userRepo.GetUser(memberId)
	if err != nil {
		return errors.New("member not found")
	}

	// check if book exists
	book, err := s.bookRepo.GetBook(bookId)
	if err != nil {
		return errors.New("book not found")
	}

	if book.CopiesAvailable <= 0 {
		return errors.New("no copies available for this book")
	}

	LoanDate := time.Now()
	DueDate := LoanDate.AddDate(0, 0, payload.DurationInDays)

	// validate loan days
	if payload.DurationInDays < MinLoanDays || payload.DurationInDays > MaxLoanDays {
		return errors.New("loan days must be between 1 and 14")
	}

	fmt.Println(payload.DurationInDays, "duration")

	existingLoan, _ := s.loansRepo.GetLoanByMemberAndBook(memberId, bookId)

	if existingLoan != nil && existingLoan.Status == string(constants.LoanActive) {
		return errors.New("member already has an active loan for this book")
	}

	loan := &model.LoanEntity{
		MemberID: memberId,
		BookID:   bookId,
		LoanDate: LoanDate,
		DueDate:  DueDate,
		Status:   string(constants.LoanActive),
		Duration: payload.DurationInDays,
	}

	er := s.loansRepo.CreateLoan(*loan)
	if er != nil {
		fmt.Println(er)
		return errors.New("failed to create loan")
	}

	// decrement available copies
	err = s.bookRepo.DecrementAvailable(bookId)
	if err != nil {
		fmt.Println(err)
		return errors.New("failed to create loan")
	}

	return nil
}

func (s *LoansService) GetLoans() ([]*types.LoanResponse, error) {
	return s.loansRepo.GetLoans()
}

func (s *LoansService) ReturnBook(loanId string, bookId string, memberId string) error {
	return s.loansRepo.ReturnBook(loanId, bookId, memberId)
}

func (s *LoansService) GetMemberLoans(memberId string) error {
	return s.loansRepo.GetMemberLoans(memberId)
}
