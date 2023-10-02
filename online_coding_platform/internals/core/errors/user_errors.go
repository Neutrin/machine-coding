package errors

import "errors"

var (
	UserNotFoundError = errors.New("user not found")
	UserSaveFailed    = errors.New(" not able to save user")
)
