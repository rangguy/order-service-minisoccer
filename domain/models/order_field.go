package models

import (
	"github.com/google/uuid"
	"time"
)

type OrderField struct {
	ID              uint      `gorm:"primary_key;autoIncrement"`
	OrderID         uint      `gorm:"type:bigint;not null"`
	FieldScheduleID uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
}
