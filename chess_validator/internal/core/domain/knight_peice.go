package domain

import "fmt"

type Knight struct {
	Color            ColorEnum
	IsJumpAllowed    bool
	rowJump, colJump []int
}

func NewKnight(color ColorEnum) Peice {
	return &Knight{
		Color:         color,
		IsJumpAllowed: false,
		rowJump:       []int{1, 1, 2, 2, 1, -2, -1, -2, -1},
		colJump:       []int{-2, 2, 1, -1, -2, -1, -2, 1, 2},
	}
}

func (knight *Knight) IsValidMove(to, from *Position, board *Board, move ...int) error {
	var (
		rowDiff  = to.RowNo - from.RowNo
		colDiff  = to.ColNo - from.ColNo
		validPos bool
	)
	for index := range knight.rowJump {
		if knight.rowJump[index] == rowDiff && knight.colJump[index] == colDiff {
			validPos = true
			break
		}
	}
	if validPos {
		if board.peices[to.RowNo][to.ColNo] == nil || !board.peices[to.RowNo][to.ColNo].IsSameColor(knight.Color) {
			return nil
		}
	}
	return fmt.Errorf(" not a valid move ")
}

func (knight *Knight) Layout() string {
	rep := "BN"
	if knight.Color == White {
		rep = "WN"
	}
	return rep
}

func (knight *Knight) IsSameColor(color ColorEnum) bool {
	return color == knight.Color
}
