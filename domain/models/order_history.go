package models

import (
	"order-service/constants"
	"time"
)

type OrderHistory struct {
	ID        uint                        `gorm:"primary_key;autoIncrement"`
	OrderID   uint                        `gorm:"type:bigint;not null"`
	Status    constants.OrderStatusString `gorm:"type:varchar(30);not null"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
