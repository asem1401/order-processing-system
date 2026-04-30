package usecase

import (
	"log"

	"order-service/internal/domain"
	"github.com/google/uuid"
)

type PaymentClient interface {
	ProcessPayment(orderID string, amount int64, email string) (string, error)
}

type OrderUsecase struct {
	payment PaymentClient
}

func NewOrderUsecase(p PaymentClient) *OrderUsecase {
	return &OrderUsecase{payment: p}
}

func (u *OrderUsecase) CreateOrder(amount int64, email string) domain.CreateOrderResponse {
	orderID := uuid.New().String()

	status, err := u.payment.ProcessPayment(orderID, amount, email)

	if err != nil {
		log.Println(err)
		return domain.CreateOrderResponse{
			OrderID: orderID,
			Status:  "failed",
			Amount:  amount,
		}
	}

return domain.CreateOrderResponse{
	OrderID: orderID,
	Status:  status, 
	Amount:  amount,
}
}