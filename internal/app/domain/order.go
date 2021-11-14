package domain

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Order is the model of Order
type Order struct {
	ID           uuid.UUID       `gorm:"type:uuid;pk"`
	AccountID    uuid.UUID       `gorm:"type:uuid;not null"`
	InstrumentID uuid.UUID       `gorm:"type:uuid;not null"`
	Type         string          `gorm:"type:text"`
	Price        decimal.Decimal `gorm:"type:decimal(10,2);not null"`
	Volume       int64           `gorm:"type:uint;not null"`
	Status       string          `gorm:"type:text"`

	Instrument Instrument `gorm:"foreignKey:InstrumentID"`
}

// TableName gets table name of Order
func (Order) TableName() string {
	return "orders"
}

// CreateOrderRequest struct for create new order.
type CreateOrderRequest struct {
	AccountID    uuid.UUID       `form:"account_id" json:"account_id"`
	InstrumentID uuid.UUID       `form:"instrument_id" json:"instrument_id"`
	Type         string          `form:"type" json:"type"`
	Price        decimal.Decimal `form:"price" json:"price"`
	Volume       int64           `form:"volume" json:"volume"`
}

func (e *CreateOrderRequest) Bind(*http.Request) error {
	return nil
}
