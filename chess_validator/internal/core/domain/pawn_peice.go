package domain

import (
	"fmt"
	"math"
)

type Pawn struct {
	Color         ColorEnum
	IsJumpAllowed bool
}

func NewPawn(color ColorEnum) Peice {
	return &Pawn{
		Color:         color,
		IsJumpAllowed: false,
	}
}

func (pawn *Pawn) IsValidMove(to, from *Position, board *Board, move ...int) error {
	var (
		rowDiff = to.RowNo - from.RowNo
		colDiff = int(math.Abs(float64(from.ColNo - to.ColNo)))
	)
	if board.peices[to.RowNo][to.ColNo] != nil {
		if !(rowDiff == move[0] && colDiff == 1) {
			return fmt.Errorf(" invalid move 27")
		}
	} else {
		if !(rowDiff == move[0] && colDiff == 0) {
			return fmt.Errorf(" invalid move 31")
		}
	}
	return nil
}

func (pawn *Pawn) Layout() string {
	rep := "BP"
	if pawn.Color == White {
		rep = "WP"
	}
	return rep
}

func (pwan *Pawn) IsSameColor(color ColorEnum) bool {
	return color == pwan.Color
}
