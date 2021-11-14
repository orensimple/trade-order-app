package usecase

import (
	"github.com/orensimple/trade-order-app/internal/app/domain"
	"github.com/orensimple/trade-order-app/internal/app/domain/repository"
)

// CreateOrder create new order
func CreateOrder(r repository.Order, a *domain.Order) (*domain.Order, error) {
	err := r.Create(a)

	return a, err
}

// GetOrder find order by filter
func GetOrder(r repository.Order, f *domain.Order) (*domain.Order, error) {
	res, err := r.Get(f)

	return res, err
}

// FindOrders find orders by filter
func FindOrders(r repository.Order, f *domain.Order) ([]*domain.Order, error) {
	res, err := r.Find(f)

	return res, err
}

// UpdateOrder update order
func UpdateOrder(r repository.Order, f *domain.Order) error {
	return r.Update(f)
}

// DeleteOrder delete order by id
func DeleteOrder(r repository.Order, f *domain.Order) error {
	return r.Delete(f)
}
