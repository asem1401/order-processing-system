package repository

import "payment-service/internal/domain"

type PaymentRepository interface {
	Save(p domain.Payment)
	GetByOrderID(orderID string) *domain.Payment
}

type InMemoryRepo struct {
	data map[string]domain.Payment
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		data: make(map[string]domain.Payment),
	}
}

func (r *InMemoryRepo) Save(p domain.Payment) {
	r.data[p.OrderID] = p
}

func (r *InMemoryRepo) GetByOrderID(orderID string) *domain.Payment {
	p, ok := r.data[orderID]
	if !ok {
		return nil
	}
	return &p
}
