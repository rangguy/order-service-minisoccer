package dto

import (
	"github.com/google/uuid"
	"time"
)

type PaymentRequest struct {
	OrderID        uuid.UUID       `json:"orderID"`
	ExpiredAt      time.Time       `json:"expiredAt"`
	Amount         float64         `json:"amount"`
	Description    *string         `json:"description"`
	CustomerDetail *CustomerDetail `json:"customerDetail"`
	ItemDetails    []ItemDetails   `json:"itemDetails"`
}

type CustomerDetail struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type ItemDetails struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Amount   float64   `json:"amount"`
	Quantity int       `json:"quantity"`
}
