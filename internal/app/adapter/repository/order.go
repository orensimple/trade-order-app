package repository

import (
	"errors"

	"github.com/orensimple/trade-order-app/internal/app/domain"
	"gorm.io/gorm"
)

// Order is the repository of domain.Order
type Order struct {
	repo *gorm.DB
}

func NewOrderRepo(db *gorm.DB) Order {
	return Order{repo: db.Debug()}
}

// Create new order
func (u Order) Create(a *domain.Order) error {
	return u.repo.Create(a).Error
}

// Get order by filter
func (u Order) Get(f *domain.Order) (*domain.Order, error) {
	out := new(domain.Order)

	err := u.repo.Where(f).Preload("Instrument").Take(out).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("order not found")
		}

		return nil, errors.New("failed get order")
	}

	return out, nil
}

// Find orders by filter
func (u Order) Find(f *domain.Order) ([]*domain.Order, error) {
	out := make([]*domain.Order, 0)

	err := u.repo.Where(f).Preload("Instrument").Find(&out).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("failed search orders")
	}

	return out, nil
}

// Update order info by id
func (u Order) Update(a *domain.Order) error {
	return u.repo.Save(&a).Error
}

// Delete order by id
func (u Order) Delete(f *domain.Order) error {
	return u.repo.Delete(&f).Error
}
