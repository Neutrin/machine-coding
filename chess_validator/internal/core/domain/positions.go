package domain

import "fmt"

type Position struct {
	RowNo int
	ColNo int
}

func NewPosition(pos string) (*Position, error) {
	var (
		position *Position
		err      error
	)
	if len(pos) != 2 {
		return position, fmt.Errorf(" invalid postion")
	}
	position = &Position{}
	for index, curRune := range pos {
		if index == 0 {

			position.ColNo = int(curRune) - 96

		} else {

			position.RowNo = int(curRune) - 48
		}
	}
	return position, err
}
