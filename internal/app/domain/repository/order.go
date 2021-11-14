package repository

import (
	"github.com/orensimple/trade-order-app/internal/app/domain"
)

// Order is interface of order repository
type Order interface {
	Create(u *domain.Order) error
	Get(f *domain.Order) (*domain.Order, error)
	Find(f *domain.Order) ([]*domain.Order, error)
	Update(f *domain.Order) error
	Delete(f *domain.Order) error
}
