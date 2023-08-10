package domain

import "time"

var (
	Location *time.Location
	err      error
)

const (
	TimeFormat = "02-01-2006 15:04:05"
)

func init() {

	Location, err = time.LoadLocation("Asia/Kolkata")
	if err != nil {
		panic(err)
	}
}
