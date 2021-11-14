package usecase

import (
	"github.com/orensimple/trade-order-app/internal/app/adapter/service"
	"github.com/orensimple/trade-order-app/internal/app/domain"
)

// BillingPay is the UseCase of pay money in billing
func BillingPay(e service.Billing, o *domain.Order) error {
	return e.Pay(o)
}

// BillingBlocked is the UseCase of blocked money in billing
func BillingBlocked(e service.Billing, o *domain.Order) error {
	return e.Blocked(o)
}
