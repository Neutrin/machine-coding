package ports

import "github.com/online_coding_platform/internals/core/domains"

type QuestionReposteries interface {
	Save(question domains.Question) error
	GetAll() ([]domains.Question, error)
	GetByLevel(level int64) ([]domains.Question, error)
}
