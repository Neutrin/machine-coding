package domains

import (
	"sync/atomic"
	"time"

	custom_enums "github.com/online_coding_platform/internals/core/domains/enums"
	custom_errors "github.com/online_coding_platform/internals/core/errors"
)

var contestId int64

type Contest struct {
	id     int64
	level  int64
	name   string
	status int64

	createdBy   User
	registredBy []User

	createdAt time.Time
}

func NewContest(level int64, name string, createdBy User) (Contest, error) {
	if !custom_enums.ValidLevel(level) {
		return Contest{}, custom_errors.InvalidContestLevel
	}
	contestId = atomic.AddInt64(&contestId, 1)
	return Contest{
		id:          contestId,
		level:       level,
		name:        name,
		createdBy:   createdBy,
		registredBy: make([]User, 0),
		createdAt:   time.Now(),
		status:      custom_enums.Created,
	}, nil
}

func NewContestWithId(id int64, createdBy User, name string, level int64,
	regiteredBy []User, status int64, createdAt time.Time) Contest {
	return Contest{
		id:          id,
		level:       level,
		createdBy:   createdBy,
		registredBy: regiteredBy,
		createdAt:   createdAt,
		name:        name,
		status:      status,
	}
}
func (contest Contest) Id() int64 {
	return contest.id
}

func (contest Contest) Level() int64 {
	return contest.level
}

func (contest Contest) CreatedBy() User {
	return contest.createdBy
}

func (contest Contest) Register(user User) Contest {
	contest.registredBy = append(contest.registredBy, user)
	return contest
}

func (contest Contest) CreatedAt() time.Time {
	return contest.createdAt
}

func (contest Contest) RegisterBy() []User {
	return contest.registredBy
}

func (contest Contest) Name() string {
	return contest.name
}

func (contest Contest) Status() int64 {
	return contest.status
}
