package repositiories

import (
	"time"

	"github.com/todo_list/internal/core/domain"
)

var CompareFunction = func(valueOne string, valueTwo string) bool {
	dateOne, _ := time.ParseInLocation(domain.TimeFormat, valueOne, domain.Location)
	dateTwo, _ := time.ParseInLocation(domain.TimeFormat, valueTwo, domain.Location)
	if dateOne.Equal(dateTwo) || dateOne.Before(dateTwo) {
		return true
	}
	return false
}
