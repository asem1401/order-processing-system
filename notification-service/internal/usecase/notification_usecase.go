package usecase

import (
	"log"

	"notification-service/internal/domain"
)

type NotificationUsecase struct {
	processed map[string]bool
}

func NewNotificationUsecase() *NotificationUsecase {
	return &NotificationUsecase{
		processed: make(map[string]bool),
	}
}

func (u *NotificationUsecase) Handle(event domain.PaymentEvent) {

	
	if u.processed[event.EventID] {
		return
	}

	
	log.Printf("[Notification] Sent email to %s for Order %s. Amount: %d",
		event.Email, event.OrderID, event.Amount)

	u.processed[event.EventID] = true
}