package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ESGScore struct {
	ID                       uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	BankID                   uuid.UUID `gorm:"type:uuid;not null;index"`
	CompanyName              string    `gorm:"type:text"`
	CompanyEncrypted         []byte    `gorm:"type:bytea"`
	RevenueEncrypted         []byte    `gorm:"type:bytea"`
	CarbonEmissionsEncrypted []byte    `gorm:"type:bytea"`
	ESGScore                 float64   `gorm:"type:numeric(3,2);not null"`
	TradingSignal            string    `gorm:"type:jsonb"`
	RiskAction               string    `gorm:"type:text"`
	AuditHash                string    `gorm:"type:text"`
	CreatedAt                time.Time `gorm:"default:now()"`
	UpdatedAt                time.Time
}

func (ESGScore) TableName() string {
	return "esg_scores"
}

func (e *ESGScore) BeforeCreate(tx *gorm.DB) error {
	if e.ID == uuid.Nil {
		e.ID = uuid.New()
	}
	return nil
}
