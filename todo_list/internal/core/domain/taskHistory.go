package domain

import (
	"fmt"

	"github.com/todo_list/internal/core/domain/enums"
)

type TaskHistory struct {
	TaskId   int64
	EventTyp enums.Event
	Time     string
}

type History struct {
	Tasks []TaskHistory
}

func (history TaskHistory) String() string {
	return fmt.Sprintf(" Task with id = %d was %s \n at time = %s\n",
		history.TaskId, history.EventTyp.String(), history.Time)
}
