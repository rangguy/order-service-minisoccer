package clients

import "github.com/google/uuid"

type PaymentResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    PaymentData `json:"data"`
}

type PaymentData struct {
	UUID          uuid.UUID `json:"uuid"`
	OrderID       string    `json:"order_id"`
	Amount        float64   `json:"amount"`
	Status        string    `json:"status"`
	PaymentLink   string    `json:"paymentLink"`
	InvoiceLink   *string   `json:"invoiceLink,omitempty"`
	Description   *string   `json:"description"`
	VANumber      *string   `json:"vaNumber,omitempty"`
	Bank          *string   `json:"bank,omitempty"`
	TransactionID *string   `json:"transactionId,omitempty"`
	Acquirer      *string   `json:"acquirer,omitempty"`
	PaidAt        *string   `json:"paidAt,omitempty"`
	ExpiredAt     *string   `json:"expiredAt"`
	CreatedAt     string    `json:"createdAt"`
	UpdatedAt     string    `json:"updatedAt"`
}
