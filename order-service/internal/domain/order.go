package domain

import "time"

type Order struct {
	ID         string
	CustomerID string
	ItemName   string
	Amount     int64
	Status     string 
	CreatedAt  time.Time
}
type CreateOrderResponse struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
	Amount  int64  `json:"amount"`
}
const (
	StatusPending   = "pending"
	StatusPaid      = "paid"
	StatusFailed    = "failed"
	StatusCancelled = "cancelled"
)
