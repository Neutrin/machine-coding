package models

import (
	"fmt"
)

type DoctorDetBySpeciality struct {
	Name      string
	StartTime string
	EndTime   string
}

func (resp DoctorDetBySpeciality) String() string {
	return fmt.Sprintf("Dr.%s: (%s-%s)", resp.Name, resp.StartTime, resp.EndTime)
}

type BookingPateintResp struct {
	BookingId int
	DocName   string
	Time      string
}

func (resp BookingPateintResp) String() string {
	return fmt.Sprintf("Booking id : %d Dr. %s %s", resp.BookingId, resp.DocName, resp.Time)
}
