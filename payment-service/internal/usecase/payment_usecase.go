package usecase

import (
	"payment-service/internal/domain"

	"github.com/google/uuid"
)

type Publisher interface {
	Publish(event domain.PaymentEvent)
}

type PaymentUsecase struct {
	pub Publisher
}

func NewPaymentUsecase(p Publisher) *PaymentUsecase {
	return &PaymentUsecase{pub: p}
}

func (u *PaymentUsecase) Process(orderID string, amount int64, email string) string {
	status := "Authorized"

	event := domain.PaymentEvent{
	EventID: uuid.New().String(),
	OrderID: orderID,
	Amount:  amount,
	Email:   email, 
	Status:  "completed",
}

	u.pub.Publish(event)

	return status
}