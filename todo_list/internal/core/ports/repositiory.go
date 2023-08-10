package ports

import (
	"github.com/todo_list/internal/core/domain"
)

type Repositiory interface {
	//This will save task in db
	SaveTask(domain.Task, string) (int64, error)
	//To get a task
	GetTask(id int64) (domain.Task, error)
	//Returns list of all task
	GetAllTask() ([]domain.Task, error)
	EventsInRange(rangeOne string, rangeTwo string) domain.History
	StopRoutine()
}
