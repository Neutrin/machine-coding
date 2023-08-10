package repositiories

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/todo_list/internal/core/domain"
	"github.com/todo_list/internal/core/domain/enums"
	"github.com/todo_list/internal/core/ports"
)

type MapRepo struct {
	mp       map[int64][]byte
	index    Index
	stopFlag chan bool
}

func NewMapRepo(index Index) ports.Repositiory {
	mRepo := MapRepo{
		mp:       make(map[int64][]byte),
		index:    index,
		stopFlag: make(chan bool, 1),
	}
	go markSpilled(&mRepo)
	return &mRepo
}

/*
check all the task with status active
move them to spilled if their due date is already done
add them to spilled event and count will increase automatically
*/

func markSpilled(repo *MapRepo) {
	fmt.Println("**********************Started go routine************************")
	for len(repo.stopFlag) == 0 {
		tasks, _ := repo.GetAllTask()
		for _, curTask := range tasks {
			if curTask.Status == enums.Active {
				dueDate, _ := time.ParseInLocation(
					domain.TimeFormat, curTask.DueDate, domain.Location)
				curTime := time.Now().In(domain.Location)
				//fmt.Println(" due date becomes =", dueDate)
				//fmt.Println(" cur time becomes =", curTime)
				if dueDate.Before(curTime) {
					curTask.Status = enums.Spilled
					curTask.UpdatedAt = time.Now().In(domain.Location).Format(domain.TimeFormat)
					repo.SaveTask(curTask, enums.Spilled.String())
					//fmt.Println(" task moved to spilled ", curTask.Id)
				}
			}
		}
		time.Sleep(30 * time.Second)
	}
	fmt.Println(" ******************Exiting go routine**************")
}

func (repo *MapRepo) StopRoutine() {
	repo.stopFlag <- false

}

// To save task in database
func (repo *MapRepo) SaveTask(task domain.Task, eventType string) (int64, error) {
	var (
		err     error
		byteRep []byte
	)
	if byteRep, err = json.Marshal(task); err != nil {
		return 0, err
	}
	repo.mp[task.Id] = byteRep
	if strings.Compare(eventType, enums.Created.String()) == 0 {

		repo.index.AddEntry(task.CreatedAt, task.Id, eventType)
	} else {

		repo.index.AddEntry(task.UpdatedAt, task.Id, eventType)
	}
	return task.Id, nil
}

func (repo *MapRepo) GetTask(id int64) (domain.Task, error) {
	return repo.taskById(id)
}

func (repo *MapRepo) taskById(id int64) (domain.Task, error) {
	var (
		task domain.Task
		err  error
	)
	if byteRep, found := repo.mp[id]; found {
		//TODO support errors in future
		_ = json.Unmarshal(byteRep, &task)
	} else {
		err = fmt.Errorf(" task with id = %d not found", id)
	}
	return task, err
}

func (repo *MapRepo) GetAllTask() ([]domain.Task, error) {
	var (
		err   error
		tasks = make([]domain.Task, 0)
	)
	for taskId, _ := range repo.mp {
		curTask, err := repo.taskById(taskId)
		if err != nil {
			break
		}
		tasks = append(tasks, curTask)
	}
	return tasks, err
}

func (repo *MapRepo) EventsInRange(rangeOne string, rangeTwo string) domain.History {
	var history = domain.History{
		Tasks: make([]domain.TaskHistory, 0),
	}
	for _, curEntry := range repo.index.FindEventInRanges(rangeOne, rangeTwo) {
		history.Tasks = append(history.Tasks, domain.TaskHistory{
			TaskId:   curEntry.Id,
			EventTyp: enums.GetEventIndex(curEntry.Event),
			Time:     curEntry.key,
		})
	}
	return history
}
