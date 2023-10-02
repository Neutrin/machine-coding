package domains

import (
	"sync/atomic"

	custom_enums "github.com/online_coding_platform/internals/core/domains/enums"
	custom_errors "github.com/online_coding_platform/internals/core/errors"
)

var id int64

type Question struct {
	id     int64
	points int64
	level  int64
}

func NewQuestion(points, level int64) (Question, error) {
	if !custom_enums.ValidLevel(level) {
		return Question{}, custom_errors.QuestionInvalidLevel
	}
	if points <= 0 {
		return Question{}, custom_errors.QuestionPointLessThanZero
	}
	id = atomic.AddInt64(&id, 1)
	return Question{
		id:     id,
		points: points,
		level:  level,
	}, nil

}

func NewQuestionWithId(id, points, level int64) (Question, error) {
	if !custom_enums.ValidLevel(level) {
		return Question{}, custom_errors.QuestionInvalidLevel
	}
	if points <= 0 {
		return Question{}, custom_errors.QuestionPointLessThanZero
	}
	return Question{
		id:     id,
		points: points,
		level:  level,
	}, nil
}

func (question Question) Id() int64 {
	return question.id
}

func (question Question) Points() int64 {
	return question.points
}

func (question Question) Level() int64 {
	return question.level
}
