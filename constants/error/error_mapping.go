package error

import (
	errPayment "payment-service/constants/error/payment"
)

func ErrMapping(err error) bool {
	var (
		GeneralErrors = GeneralErrors
		TimeErrors    = errPayment.PaymentErrors
	)

	allErrors := make([]error, 0)
	allErrors = append(allErrors, GeneralErrors...)
	allErrors = append(allErrors, TimeErrors...)

	for _, item := range allErrors {
		if err.Error() == item.Error() {
			return true
		}
	}

	return false
}
