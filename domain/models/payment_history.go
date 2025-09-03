package models

import (
	"payment-service/constants"
	"time"
)

type PaymentHistory struct {
	ID        uint                          `gorm:"primaryKey;autoIncrement"`
	PaymentID uint                          `gorm:"bigint;autoIncrement"`
	Status    constants.PaymentStatusString `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
