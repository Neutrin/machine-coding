package errors

import "errors"

var (
	QuestionPointLessThanZero = errors.New(" points should be greater than 0")
	QuestionInvalidLevel      = errors.New(" invalid level")
)
