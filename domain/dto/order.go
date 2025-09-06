package dto

import (
	"github.com/google/uuid"
	"order-service/constants"
	"time"
)

type OrderRequest struct {
	FieldScheduleIDs []string `json:"fieldScheduleIDs" validate:"required"`
}

type OrderRequestParam struct {
	Page       int     `form:"page" validate:"required"`
	Limit      int     `form:"limit" validate:"required"`
	SortColumn *string `form:"sortColumn"`
	SortOrder  *string `form:"sortOrder"`
}

type OrderResponse struct {
	UUID        uuid.UUID                   `json:"uuid"`
	Code        string                      `json:"code"`
	Username    string                      `json:"username"`
	Amount      float64                     `json:"amount"`
	Status      constants.OrderStatusString `json:"status"`
	PaymentLink string                      `json:"paymentLink"`
	OrderDate   time.Time                   `json:"orderDate"`
	CreatedAt   time.Time                   `json:"createdAt"`
	UpdatedAt   time.Time                   `json:"updatedAt"`
}

type OrderByUserIDResponse struct {
	Code        string                      `json:"code"`
	Amount      float64                     `json:"amount"`
	Status      constants.OrderStatusString `json:"status"`
	OrderDate   string                      `json:"orderDate"`
	PaymentLink string                      `json:"paymentLink"`
	InvoiceLink *string                     `json:"invoiceLink"`
}
