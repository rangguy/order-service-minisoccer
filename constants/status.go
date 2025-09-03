package constants

type PaymentStatus int
type PaymentStatusString string

const (
	Initial    PaymentStatus = 0
	Pending    PaymentStatus = 100
	Settlement PaymentStatus = 200
	Expire     PaymentStatus = 300

	InitialString    PaymentStatusString = "initial"
	PendingString    PaymentStatusString = "pending"
	SettlementString PaymentStatusString = "settlement"
	ExpireString     PaymentStatusString = "expire"
)

var mapStatusStringToInt = map[PaymentStatusString]PaymentStatus{
	InitialString:    Initial,
	PendingString:    Pending,
	SettlementString: Settlement,
	ExpireString:     Expire,
}

var mapStatusIntToString = map[PaymentStatus]PaymentStatusString{
	Initial:    InitialString,
	Pending:    PendingString,
	Settlement: SettlementString,
	Expire:     ExpireString,
}

func (p PaymentStatusString) String() string {
	return string(p)
}

func (p PaymentStatus) Int() int {
	return int(p)
}

func (p PaymentStatus) GetStatusString() PaymentStatusString {
	return mapStatusIntToString[p]
}

func (p PaymentStatusString) GetStatusInt() PaymentStatus {
	return mapStatusStringToInt[p]
}
