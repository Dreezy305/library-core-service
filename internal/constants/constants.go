package constants

type LoanStatus string

const (
	LoanActive   LoanStatus = "active"
	LoanReturned LoanStatus = "returned"
	LoanOverdue  LoanStatus = "overdue"
)
