package domain

import (
	"fmt"
	"time"
)

var slotId int

type Slot struct {
	Id        int
	StartTime time.Time
	EndTime   time.Time
	Status    StatusEnum
}

func NewSlot(startTime, endTime string) (*Slot, error) {
	err := validateSlot(startTime, endTime)
	if err != nil {
		return nil, err
	}
	slotId = slotId + 1
	start, _ := time.Parse(TimeFromat, startTime)
	end, _ := time.Parse(TimeFromat, endTime)
	slot := &Slot{
		Id:        slotId,
		StartTime: start,
		EndTime:   end,
		Status:    Available,
	}
	return slot, nil
}

func validateSlot(startTime, endTime string) error {
	var (
		morningTime, _ = time.Parse(TimeFromat, MorningTime)
		EveningTime, _ = time.Parse(TimeFromat, EveningTime)
	)
	start, err := time.Parse(TimeFromat, startTime)
	if err != nil {
		return fmt.Errorf(" invalid time format = %s", startTime)
	}
	end, err := time.Parse(TimeFromat, endTime)
	if err != nil {
		return fmt.Errorf(" invalid time format = %s", startTime)
	}
	if start.Before(morningTime) || end.After(EveningTime) {
		return fmt.Errorf(" invalid slot")
	}
	if timeDiff := end.Sub(start).Minutes(); timeDiff != SlotRangeInMinutes {
		return fmt.Errorf(" all slots should be of = %d mintes", SlotRangeInMinutes)
	}

	return nil
}
