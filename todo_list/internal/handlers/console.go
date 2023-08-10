package handlers

import (
	"github.com/todo_list/internal/core/domain"
	"github.com/todo_list/internal/core/domain/enums"
	"github.com/todo_list/internal/core/domain/validations"
	"github.com/todo_list/internal/core/ports"
)

type Console struct {
	service ports.Service
}

func NewConsole(service ports.Service) *Console {
	return &Console{service: service}
}

func (c *Console) CreateTask(desc string, dueDate string, createdBy int64, tags ...enums.Tags) (
	TaskCreatedResp, error) {
	var (
		resp TaskCreatedResp
		err  error
	)
	task := domain.NewTask(desc, dueDate, createdBy)
	if len(tags) > 0 {
		task.SetTag(tags[0])
	}
	if err = validations.ValidateStruct(*task); err != nil {
		return resp, err
	}
	if err = c.service.CreateTask(*task); err != nil {
		return resp, err
	}
	resp.SetResponse("Task created succesfully", err)
	return resp, err

}

func (c *Console) GetAllTask() TaskResponse {
	resp, err := c.service.AllTask()
	return GenerateTaskResponse(resp, err)
}

func (c *Console) Task(id int64) TaskResponse {
	resp, err := c.service.Task(id)
	return GenerateTaskResponse([]domain.Task{resp}, err)
}

func (c *Console) MarkCompleted(id int64) TaskCreatedResp {
	resp := TaskCreatedResp{}
	err := c.service.MarkCompleted(id)
	resp.SetResponse(" Task marked completed", err)
	return resp
}

func (c *Console) UpdateDesc(id int64, desc string) TaskCreatedResp {
	resp := TaskCreatedResp{}
	err := c.service.UpdateDesc(id, desc)
	resp.SetResponse("Task updated succesfully", err)
	return resp
}

func (c *Console) UpdateDueDate(id int64, dueDate string) TaskCreatedResp {
	resp := TaskCreatedResp{}
	err := c.service.ChangeDueDate(id, dueDate)
	resp.SetResponse("DueDate updated succesfully", err)
	return resp
}

func (c *Console) UpdateTag(id int64, tags enums.Tags) TaskCreatedResp {
	resp := TaskCreatedResp{}
	err := c.service.UpdateTag(id, tags)
	resp.SetResponse(" tags updated succesfully", err)
	return resp
}

func (c *Console) ActivityLogs(rangeOne string, rangeTwo string) []ActivityLog {
	return GenerateActivityLog(c.service.EventsInRange(rangeOne, rangeTwo))
}

func (c *Console) TaskStats(rangeOne string, rangeTwo string) *TaskStatsResponse {
	return GenerateTaskStatsResponse(c.service.TaskStatusInRange(rangeOne, rangeTwo))
}
