package domain

import (
	"github.com/google/uuid"
)

// Instrument is the model of Instrument
type Instrument struct {
	ID           uuid.UUID `gorm:"type:uuid;pk"`
	CurrencyCode string    `gorm:"type:text;not null"`
	Code         string    `gorm:"type:text"`
	Name         string    `gorm:"type:text"`
	Description  string    `gorm:"type:text"`
	Lot          int       `gorm:"type:uint;not null"`
	Type         string    `gorm:"type:text"`
}

// TableName gets table name of Order
func (Instrument) TableName() string {
	return "instruments"
}
