package handlers

import (
	"strconv"

	"github.com/todo_list/internal/core/domain"
)

type TaskCreatedResp struct {
	Resp     string
	ErrorStr string
}

func (res TaskCreatedResp) SetResponse(resp string, err error) {
	if err == nil {
		res.Resp = resp
	} else {
		res.ErrorStr = err.Error()
	}
}

type TaskRespBody struct {
	Id          string
	Desc        string
	DueDate     string
	CompletedAt string
	CreatedBy   string
	Tag         string
	Status      string
}

type TaskResponse struct {
	TaskRespBody []TaskRespBody
	Error        string
}

func copyTaskResp(task domain.Task) TaskRespBody {
	respBody := TaskRespBody{
		Id:          strconv.Itoa(int(task.Id)),
		Desc:        task.Desc,
		DueDate:     task.DueDate,
		CompletedAt: task.CompletedAt,
		CreatedBy:   strconv.Itoa(int(task.CreatedBy)),
		Status:      task.Status.String(),
	}
	if task.Tags > 0 {
		respBody.Tag = task.Tags.String()
	}
	return respBody
}

func GenerateTaskResponse(tasks []domain.Task, err error) TaskResponse {
	resp := TaskResponse{}
	if err != nil {
		resp.Error = err.Error()
		return resp
	}
	resp.TaskRespBody = make([]TaskRespBody, 0)
	for _, curTask := range tasks {
		resp.TaskRespBody = append(resp.TaskRespBody, copyTaskResp(curTask))
	}
	return resp
}

type ActivityLog struct {
	Log string
}

func GenerateActivityLog(history domain.History) []ActivityLog {
	logs := make([]ActivityLog, 0)
	for _, curTaskHistory := range history.Tasks {
		logs = append(logs, ActivityLog{Log: curTaskHistory.String()})
	}
	return logs
}

type TaskStatsResponse struct {
	CompletedCount int64
	SpilledCount   int64
	ActiveCount    int64
}

func GenerateTaskStatsResponse(input domain.TaskStats) *TaskStatsResponse {
	return &TaskStatsResponse{
		CompletedCount: input.CompletedCount,
		SpilledCount:   input.SpilledCount,
		ActiveCount:    input.ActiveCount,
	}
}
