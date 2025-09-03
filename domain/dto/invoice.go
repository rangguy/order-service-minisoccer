package dto

type InvoiceRequest struct {
	InvoiceNumber string      `json:"invoiceNumber"`
	Data          InvoiceData `json:"data"`
}

type InvoiceData struct {
	PaymentDetail InvoicePaymentDetail `json:"paymentDetail"`
	Items         []InvoiceItem        `json:"items"`
	Total         string               `json:"total"`
}

type InvoicePaymentDetail struct {
	BankName      string `json:"bank_name"`
	PaymentMethod string `json:"paymentMethod"`
	VANumber      string `json:"vaNumber"`
	Date          string `json:"date"`
	IsPaid        bool   `json:"isPaid"`
}

type InvoiceItem struct {
	Description string `json:"description"`
	Price       string `json:"price"`
}
