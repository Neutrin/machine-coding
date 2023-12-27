package domain

import (
	"fmt"
	"math"
)

type Queen struct {
	Color         ColorEnum
	IsJumpAllowed bool
}

func NewQueen(color ColorEnum) Peice {
	return &Queen{
		Color:         color,
		IsJumpAllowed: false,
	}
}

func (queen *Queen) IsValidMove(to, from *Position, board *Board, move ...int) error {
	var (
		rowDiff = int(math.Abs(float64(to.RowNo - from.RowNo)))
		colDiff = int(math.Abs(float64(from.ColNo - to.ColNo)))
	)
	if rowDiff == colDiff || from.RowNo == to.RowNo || from.ColNo == to.ColNo {
		if board.peices[to.RowNo][to.ColNo] == nil || !board.peices[to.RowNo][to.ColNo].IsSameColor(queen.Color) {
			return nil
		}
	}

	return fmt.Errorf(" in valid position")
}

func (queen *Queen) Layout() string {
	rep := "BQ"
	if queen.Color == White {
		rep = "WQ"
	}
	return rep
}

func (queen *Queen) IsSameColor(color ColorEnum) bool {
	return color == queen.Color
}
