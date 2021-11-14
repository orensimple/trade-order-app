package service

import (
	"github.com/orensimple/trade-order-app/internal/app/domain"
)

// Billing is interface of billing app http
type Billing interface {
	Pay(o domain.Order) error
	Blocked(o domain.Order) error
}
