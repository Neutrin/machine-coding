package services

import (
	"strings"
	"time"

	"github.com/todo_list/internal/core/domain"
	"github.com/todo_list/internal/core/domain/enums"
	"github.com/todo_list/internal/core/ports"
)

type Service struct {
	repo ports.Repositiory
}

func NewService(repo ports.Repositiory) ports.Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateTask(task domain.Task) error {
	//We will mark the task as active
	//To do we can check if the timeline of task is less than currrent
	task.MarkActive()
	task.CreatedAt = time.Now().Format(domain.TimeFormat)
	_, err := s.repo.SaveTask(task, enums.Created.String())
	return err
}

func (s *Service) Task(id int64) (domain.Task, error) {
	return s.repo.GetTask(id)
}

func (s *Service) MarkCompleted(id int64) error {
	task, err := s.repo.GetTask(id)
	if err != nil {
		return err
	}
	if strings.Compare(task.Status.String(), enums.Completed.String()) != 0 {
		task.Status = enums.Completed
		task.CompletedAt = time.Now().Format(domain.TimeFormat)
		task.UpdatedAt = time.Now().Format(domain.TimeFormat)
		_, err = s.repo.SaveTask(task, enums.Compeleted.String())
	}
	return err
}

func (s *Service) UpdateDesc(id int64, desc string) error {
	task, err := s.repo.GetTask(id)
	if err != nil {
		return err
	}
	if strings.Compare(task.Desc, desc) != 0 {
		task.Desc = desc
		task.UpdatedAt = time.Now().Format(domain.TimeFormat)
		_, err = s.repo.SaveTask(task, enums.Updated.String())
	}
	return err
}

func (s *Service) ChangeDueDate(id int64, newDate string) error {
	task, err := s.repo.GetTask(id)
	if err != nil {
		return err
	}
	if strings.Compare(task.DueDate, newDate) != 0 {
		task.DueDate = newDate
		task.UpdatedAt = time.Now().Format(domain.TimeFormat)
		_, err = s.repo.SaveTask(task, enums.Updated.String())
	}
	return err
}

func (s *Service) UpdateTag(id int64, tag enums.Tags) error {
	task, err := s.repo.GetTask(id)
	if err != nil {
		return err
	}
	if strings.Compare(task.Tags.String(), tag.String()) != 0 {
		task.Tags = tag
		task.UpdatedAt = time.Now().Format(domain.TimeFormat)
		_, err = s.repo.SaveTask(task, enums.Updated.String())
	}
	return err
}

func (s *Service) AllTask() ([]domain.Task, error) {
	return s.repo.GetAllTask()
}

func (s *Service) EventsInRange(rangeOne string, rangeTwo string) domain.History {
	return s.repo.EventsInRange(rangeOne, rangeTwo)
}

func (s *Service) TaskStatusInRange(rangeOne string, rangeTwo string) domain.TaskStats {
	var stats domain.TaskStats
	for _, curTask := range s.repo.EventsInRange(rangeOne, rangeTwo).Tasks {

		switch curTask.EventTyp {
		case enums.Compeleted:
			stats.CompletedCount++
		case enums.Spill:
			stats.SpilledCount++
		case enums.Created:
			stats.ActiveCount++
		}
	}
	return stats
}
