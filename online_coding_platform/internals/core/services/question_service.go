package services

import (
	"github.com/online_coding_platform/internals/core/domains"
	"github.com/online_coding_platform/internals/core/ports"
)

type QuestionService struct {
	repo ports.QuestionReposteries
}

func NewQuestionService(repo ports.QuestionReposteries) *QuestionService {
	return &QuestionService{
		repo: repo,
	}
}

func (service *QuestionService) CreateQuestion(level, points int64) error {
	question, err := domains.NewQuestion(points, level)
	if err != nil {
		return err
	}
	return service.repo.Save(question)
}

func (service *QuestionService) ListQuestions(level ...int64) ([]domains.Question, error) {
	if len(level) == 1 {
		return service.repo.GetByLevel(level[0])
	}
	return service.repo.GetAll()
}
