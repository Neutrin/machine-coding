package errors

import "errors"

var (
	InvalidContestLevel           = errors.New("invalid level")
	ContestNotFound               = errors.New(" contest not found ")
	ScoreCalNotDefined            = errors.New(" calculator not defined")
	ContestCreationAuthentication = errors.New(" user cannot start contest ")
	ContestInvalidStatus          = errors.New(" status invalid")
)
