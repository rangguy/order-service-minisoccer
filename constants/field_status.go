package constants

type FieldStatusString string

const (
	AvailableStatus FieldStatusString = "available"
	BookedStatus    FieldStatusString = "booked"
)

func (p FieldStatusString) String() string {
	return string(p)
}
