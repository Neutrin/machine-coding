package domain

import (
	"fmt"
)

type Board struct {
	Size   int
	peices [][]Peice
}

func NewDefBoard(size int) *Board {
	board := Board{
		Size: size,
	}
	board.peices = make([][]Peice, size+1)
	for count := 1; count <= size; count++ {
		board.peices[count] = make([]Peice, size+1)
	}
	for colIndex := 1; colIndex <= size; colIndex++ {
		board.peices[2][colIndex] = NewPawn(White)
		board.peices[size-1][colIndex] = NewPawn(Black)
	}
	for colIndex := 1; colIndex <= size; colIndex++ {
		if colIndex == 1 || colIndex == size {
			board.peices[1][colIndex] = NewRook(White)
			board.peices[size][colIndex] = NewRook(Black)
		}
		if colIndex == 2 || colIndex == (size-1) {
			board.peices[1][colIndex] = NewKnight(White)
			board.peices[size][colIndex] = NewKnight(Black)
		}
		if colIndex == 3 || colIndex == (size-2) {
			board.peices[1][colIndex] = NewBishop(White)
			board.peices[size][colIndex] = NewBishop(Black)
		}
		if colIndex == 4 {
			board.peices[1][colIndex] = NewQueen(White)
			board.peices[size][colIndex] = NewQueen(Black)
		}
		if colIndex == 5 {
			board.peices[1][colIndex] = NewKing(White)
			board.peices[size][colIndex] = NewKing(Black)
		}

	}
	return &board
}

func (b *Board) Print() {
	fmt.Println("************************************************")
	fmt.Println(" ")
	for rowIndex := b.Size; rowIndex >= 1; rowIndex-- {
		for colIndex := 1; colIndex <= b.Size; colIndex++ {
			if b.peices[rowIndex][colIndex] == nil {
				fmt.Print("-- ")
			} else {
				fmt.Print(b.peices[rowIndex][colIndex].Layout() + " ")
			}
		}
		fmt.Println(" ")
	}
	fmt.Println("*************************************************")
}

func (b *Board) IsValidPos(pos *Position) bool {
	if pos.RowNo >= 1 && pos.RowNo <= b.Size && pos.ColNo >= 1 && pos.ColNo <= b.Size {
		return true
	}
	return false
}

func (b *Board) MovePeice(peice Peice, from *Position, to *Position) {
	b.peices[to.RowNo][to.ColNo] = peice
	b.peices[from.RowNo][from.ColNo] = nil
}

func (b *Board) KillPeice(peice Peice, pos *Position) {
	b.peices[pos.RowNo][pos.ColNo] = nil
}

func (b *Board) Peice(pos *Position) Peice {
	var peice Peice
	peice = b.peices[pos.RowNo][pos.ColNo]
	return peice

}
