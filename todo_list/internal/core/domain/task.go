package domain

import (
	"sync/atomic"
	"time"

	"github.com/todo_list/internal/core/domain/enums"
)

var (
	taskID int64 = 0
)

type Task struct {
	//Sequential id generator
	Id          int64
	Desc        string       `validate:"gt=0"`
	Status      enums.Status `validate:"omitempty,custom_status"`
	Tags        enums.Tags   `validate:"omitempty,custom_tags"`
	DueDate     string       `validate:"datetime=02-01-2006 15:04:05"`
	CompletedAt string       `validate:"omitempty,datetime=02-01-2006 15:04:05"`
	CreatedAt   string       `validate:"omitempty,datetime=02-01-2006 15:04:05"`
	UpdatedAt   string       `validate:"omitempty,datetime=02-01-2006 15:04:05"`
	CreatedBy   int64        `validate:"required"`
	//This needs to be created

}

func NewTask(desc string, dueDate string, createdBy int64) *Task {

	return &Task{
		Id:        atomic.AddInt64(&taskID, 1),
		Desc:      desc,
		DueDate:   dueDate,
		CreatedBy: createdBy,
		CreatedAt: time.Now().Format(TimeFormat),
	}

}

func (t *Task) MarkActive() {
	t.Status = enums.Active
}

func (t *Task) SetTag(tag enums.Tags) {
	t.Tags = tag
}
