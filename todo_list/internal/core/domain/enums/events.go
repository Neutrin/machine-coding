package enums

import "strings"

type Event int64

const (
	Created Event = iota + 1
	Updated
	Compeleted
	Spill
	Deleted
)

var eventString = []string{
	"Created",
	"Updated",
	"Completed",
	"Spill",
	"Deleted",
}

func (e Event) String() string {
	if int64(e) >= 1 && int64(e) <= int64(len(eventString)) {
		return eventString[e-1]
	}
	return ""
}

func GetEventIndex(input string) Event {
	for index, curValue := range eventString {
		if strings.Compare(input, curValue) == 0 {
			return Event(index + 1)
		}
	}
	return Event(0)
}
