package dto

import "payment-service/constants"

type PaymentHistoryRequest struct {
	PaymentID uint                          `json:"paymentID"`
	Status    constants.PaymentStatusString `json:"status"`
}
