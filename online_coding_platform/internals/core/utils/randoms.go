package utils

import (
	"math/rand"
	"time"
)

var random *rand.Rand

func InitRandom() {
	random = rand.New(rand.NewSource(time.Now().Unix()))
}

func Number() int64 {
	return (1 + random.Int63n(2))
}
