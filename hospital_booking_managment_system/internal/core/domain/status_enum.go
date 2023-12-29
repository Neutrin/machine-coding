package domain

type StatusEnum int

const (
	Created = iota + 1
	Cancelled
	Waiting
	Booked
	Available
)

var statusString = map[StatusEnum]string{
	Created:   "Created",
	Cancelled: "Cancelled",
	Waiting:   "Waiting",
	Booked:    "Booked",
	Available: "Available",
}

func (status StatusEnum) String() string {
	return statusString[status]
}
