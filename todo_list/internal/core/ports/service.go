package ports

import (
	"github.com/todo_list/internal/core/domain"
	"github.com/todo_list/internal/core/domain/enums"
)

type Service interface {
	CreateTask(task domain.Task) error
	Task(id int64) (domain.Task, error)
	MarkCompleted(id int64) error
	AllTask() ([]domain.Task, error)
	UpdateDesc(id int64, desc string) error
	ChangeDueDate(id int64, newDate string) error
	UpdateTag(id int64, tag enums.Tags) error
	EventsInRange(rangeOne string, rangeTwo string) domain.History
	TaskStatusInRange(rangeOne string, rangeTwo string) domain.TaskStats
}
