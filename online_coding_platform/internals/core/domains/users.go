package domains

import (
	"sync/atomic"
)

var userID int64

// TODO : move to util packages
const defaultScore = int64(1500)

type User struct {
	id    int64
	name  string
	score int64
	//TODO : need to memorize code for location
	//createdAt time.Time
}

// TODO : add validation of struct
func NewUserWithDefScore(name string) User {
	userID = atomic.AddInt64(&userID, 1)
	return User{
		id:    userID,
		name:  name,
		score: defaultScore,
	}
}

func (user User) ID() int64 {
	return user.id
}

func (user User) Name() string {
	return user.name
}

func (user User) Score() int64 {
	return user.score
}

func (user User) SetID(id int64) User {
	user.id = id
	return user
}

func (user User) SetName(name string) User {
	user.name = name
	return user
}

func (user User) SetScore(score int64) User {
	user.score = score
	return user
}
