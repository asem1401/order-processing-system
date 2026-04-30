package repository

import "order-service/internal/domain"

type OrderRepository interface {
	Save(order domain.Order)
	Update(order domain.Order)
	GetByID(id string) *domain.Order
}

type InMemoryOrderRepo struct {
	data map[string]domain.Order
}

func NewOrderRepo() *InMemoryOrderRepo {
	return &InMemoryOrderRepo{
		data: make(map[string]domain.Order),
	}
}

func (r *InMemoryOrderRepo) Save(order domain.Order) {
	r.data[order.ID] = order
}

func (r *InMemoryOrderRepo) Update(order domain.Order) {
	r.data[order.ID] = order
}

func (r *InMemoryOrderRepo) GetByID(id string) *domain.Order {
	order, ok := r.data[id]
	if !ok {
		return nil
	}
	return &order
}