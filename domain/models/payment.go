package models

import (
	"github.com/google/uuid"
	"payment-service/constants"
	"time"
)

type Payment struct {
	ID               uint                     `gom:"primaryKey;autoIncrement"`
	UUID             uuid.UUID                `gorm:"type:uuid;not null"`
	OrderID          uuid.UUID                `gorm:"type:uuid;not null"`
	Amount           float64                  `gorm:"not null"`
	Status           *constants.PaymentStatus `gorm:"not null"`
	PaymentLink      string                   `gorm:"type:varchar(255);not null"`
	InvoiceLink      *string                  `gorm:"type:varchar(255);default:null"`
	VANumber         *string                  `gorm:"type:varchar(50);default:null"`
	Bank             *string                  `gorm:"type:varchar(100);default:null"`
	Acquirer         *string                  `gorm:"type:varchar(100);default:null"`
	TransactionID    *string                  `gorm:"type:varchar(100);default:null"`
	Description      *string                  `gorm:"type:text;default:null"`
	PaidAt           *time.Time
	ExpiredAt        *time.Time
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
	PaymentHistories []PaymentHistory `gorm:"foreignKey:payment_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
