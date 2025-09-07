package constants

type OrderStatus int
type OrderStatusString string

const (
	Pending        OrderStatus = 100
	PendingPayment OrderStatus = 200
	PaymentSuccess OrderStatus = 300
	Expired        OrderStatus = 400

	PendingString        OrderStatusString = "pending"
	PendingPaymentString OrderStatusString = "pending-payment"
	PaymentSuccessString OrderStatusString = "payment-success"
	ExpiredString        OrderStatusString = "expired"
)

var mapStatusStringToInt = map[OrderStatusString]OrderStatus{
	PendingString:        Pending,
	PendingPaymentString: PendingPayment,
	PaymentSuccessString: PaymentSuccess,
	ExpiredString:        Expired,
}

var mapStatusIntToString = map[OrderStatus]OrderStatusString{
	Pending:        PendingString,
	PendingPayment: PendingPaymentString,
	PaymentSuccess: PaymentSuccessString,
	Expired:        ExpiredString,
}

func (o OrderStatusString) String() string {
	return string(o)
}

func (o OrderStatus) Int() int {
	return int(o)
}

func (o OrderStatus) GetStatusString() OrderStatusString {
	return mapStatusIntToString[o]
}

func (o OrderStatusString) GetStatusInt() OrderStatus {
	return mapStatusStringToInt[o]
}

func (o OrderStatusString) GetStatusString() OrderStatusString {
	return o
}
