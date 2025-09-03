package constants

type PaymentStatusString string

const (
	PendingPaymentStatus    PaymentStatusString = "pending"
	SettlementPaymentStatus PaymentStatusString = "settlement"
	ExpirePaymentStatus     PaymentStatusString = "expire"
)
