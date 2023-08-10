package enums

import "strings"

type Tags int64

const (
	Work Tags = iota + 1
	Casual
	HighPriority
	EmptyTag
)

var (
	tagString = map[Tags]string{
		Work:         "work",
		Casual:       "casual",
		HighPriority: "high priority",
		EmptyTag:     "empty",
	}
	tagIndex = []string{"", "work", "casual", "high priority"}
)

func (t Tags) String() string {
	if str, exists := tagString[t]; exists {
		return str
	}
	return "not present"
}

func GetTagIndex(input string) int64 {
	for index, curValue := range tagIndex {
		if strings.Compare(input, curValue) == 0 {
			return int64(index)
		}
	}
	return int64(0)
}
