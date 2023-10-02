package ports

import "github.com/online_coding_platform/internals/core/domains"

type UserRepo interface {
	SaveUser(user domains.User) error
	GetUserById(id int64) (domains.User, error)
	GetAll() ([]domains.User, error)
}
