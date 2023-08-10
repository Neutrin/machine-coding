package enums

type Status int64

const (
	Active Status = iota + 1
	Completed
	Spilled
	EmptyStatus
)

var (
	statusString = map[Status]string{
		Active:      "active",
		Completed:   "completed",
		Spilled:     "spilled",
		EmptyStatus: "empty",
	}
)

func (s Status) String() string {
	if str, exists := statusString[s]; exists {
		return str
	}
	return "not present"
}
