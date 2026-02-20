package constants

type LoanStatus string

type OrderStatus string

const (
	OrderPending   OrderStatus = "pending"
	OrderPaid      OrderStatus = "paid"
	OrderCancelled OrderStatus = "cancelled"
)

const (
	LoanActive   LoanStatus = "active"
	LoanReturned LoanStatus = "returned"
	LoanOverdue  LoanStatus = "overdue"
)
