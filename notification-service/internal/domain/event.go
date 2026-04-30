package domain

type PaymentEvent struct {
	EventID string `json:"id"`
	OrderID string `json:"order_id"`
	Amount  int64  `json:"amount"`
	Email   string `json:"email"`
	Status  string `json:"status"`
}