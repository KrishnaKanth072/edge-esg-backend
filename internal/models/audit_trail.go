package models

import (
	"time"

	"github.com/google/uuid"
)

type AuditTrail struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	BankID       uuid.UUID `gorm:"type:uuid;not null;index"`
	UserID       string    `gorm:"type:text;not null"`
	Action       string    `gorm:"type:text;not null"`
	Resource     string    `gorm:"type:text"`
	Details      string    `gorm:"type:jsonb"`
	IPAddress    string    `gorm:"type:text"`
	BlockchainTx string    `gorm:"type:text"`
	CreatedAt    time.Time `gorm:"default:now()"`
}

func (AuditTrail) TableName() string {
	return "audit_trails"
}
