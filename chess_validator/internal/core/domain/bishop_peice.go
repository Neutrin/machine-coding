package domain

import (
	"fmt"
	"math"
)

type Bishop struct {
	Color         ColorEnum
	IsJumpAllowed bool
}

func NewBishop(color ColorEnum) Peice {
	return &Bishop{
		Color:         color,
		IsJumpAllowed: false,
	}
}

func (bishop *Bishop) IsValidMove(to, from *Position, board *Board, move ...int) error {
	var (
		rowDiff = int(math.Abs(float64(to.RowNo - from.RowNo)))
		colDiff = int(math.Abs(float64(from.ColNo - to.ColNo)))
	)
	if rowDiff != colDiff {
		return fmt.Errorf(" invalid moive")
	}
	if board.peices[to.RowNo][to.ColNo] != nil && !board.peices[to.RowNo][to.ColNo].IsSameColor(bishop.Color) {
		return fmt.Errorf(" ius invalid position")
	}
	return nil
}

func (bishop *Bishop) Layout() string {
	rep := "BB"
	if bishop.Color == White {
		rep = "WB"
	}
	return rep
}

func (bishop *Bishop) IsSameColor(color ColorEnum) bool {
	return color == bishop.Color
}
