package domain

import (
	"fmt"
	"math"
)

type King struct {
	Color         ColorEnum
	IsJumpAllowed bool
}

func NewKing(color ColorEnum) Peice {
	return &King{
		Color:         color,
		IsJumpAllowed: false,
	}
}

func (king *King) IsValidMove(to, from *Position, board *Board, move ...int) error {
	var (
		rowDiff = int(math.Abs(float64(to.RowNo - from.RowNo)))
		colDiff = int(math.Abs(float64(from.ColNo - to.ColNo)))
	)
	if rowDiff >= 2 || colDiff >= 2 {
		return fmt.Errorf(" invalid move")
	}
	if board.peices[to.RowNo][to.ColNo] != nil && board.peices[to.RowNo][to.ColNo].IsSameColor(king.Color) {
		return fmt.Errorf(" invalid move")
	}
	return nil
}

func (king *King) Layout() string {
	rep := "BK"
	if king.Color == White {
		rep = "WK"
	}
	return rep
}

func (king *King) IsSameColor(color ColorEnum) bool {
	return color == king.Color
}
