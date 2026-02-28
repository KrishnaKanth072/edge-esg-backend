package models

import (
	"time"

	"github.com/google/uuid"
)

type TradeSignal struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	BankID      uuid.UUID `gorm:"type:uuid;not null;index"`
	Symbol      string    `gorm:"type:text;not null"`
	Action      string    `gorm:"type:text;not null"`
	TargetPrice float64   `gorm:"type:numeric(10,2)"`
	Confidence  float64   `gorm:"type:numeric(3,2)"`
	ESGScore    float64   `gorm:"type:numeric(3,2)"`
	Status      string    `gorm:"type:text;default:'PENDING'"`
	ExecutedAt  *time.Time
	CreatedAt   time.Time `gorm:"default:now()"`
}

func (TradeSignal) TableName() string {
	return "trade_signals"
}
