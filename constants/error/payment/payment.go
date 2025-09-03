package error

import "errors"

var (
	ErrPaymentNotFound = errors.New("payment not found")
	ErrExpireAtInvalid = errors.New("expired time must be greater than current time")
)

var PaymentErrors = []error{
	ErrPaymentNotFound,
	ErrExpireAtInvalid,
}
