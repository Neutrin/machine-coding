package domain

import "fmt"

type Rook struct {
	Color         ColorEnum
	IsJumpAllowed bool
}

func NewRook(color ColorEnum) Peice {
	return &Rook{
		Color:         color,
		IsJumpAllowed: false,
	}
}
func (rook *Rook) IsValidMove(to, from *Position, board *Board, moves ...int) error {
	if to.RowNo != from.RowNo && to.ColNo != from.ColNo {
		return fmt.Errorf(" invalid move")
	}
	if board.peices[to.RowNo][to.ColNo] != nil && board.peices[to.RowNo][to.ColNo].IsSameColor(rook.Color) {
		return fmt.Errorf(" invalid position")
	}
	return nil
}

func (rook *Rook) Layout() string {
	rep := "BR"
	if rook.Color == White {
		rep = "WR"
	}
	return rep
}

func (rook *Rook) IsSameColor(color ColorEnum) bool {
	return color == rook.Color
}
