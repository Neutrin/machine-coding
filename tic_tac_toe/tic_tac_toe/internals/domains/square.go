package domains

import (
	"github.com/tic_tac_toe/internals/domains/customerrors"
	"github.com/tic_tac_toe/internals/domains/enums"
)

type SquareBoard struct {
	size       int
	strategy   WinningStrategy
	cells      [][]enums.Token
	emptyCells int
}

func NewSquareBoard(size int, strategy WinningStrategy) (Board, *customerrors.GameError) {
	//validate if size is corrent or not
	if size < 0 {
		return &SquareBoard{}, customerrors.NewGameError("invalid dimension", customerrors.ErrorCodeBoardIntialisationFailed)
	}
	cells := make([][]enums.Token, size+1)
	for count := 1; count <= size; count++ {
		cells[count] = make([]enums.Token, size+1)
	}

	return &SquareBoard{size: size, cells: cells, emptyCells: size * size, strategy: strategy}, nil
}

/*
validate move request
Validate of move can be marked on board or not
mark the move on board by the token
add the move to the winning strategy
*/
func (board *SquareBoard) AddMove(move Move, Token enums.Token) (bool, *customerrors.GameError) {
	if move.RowNo <= 0 || move.ColNo <= 0 || move.RowNo > board.size || move.ColNo > board.size {
		return false, customerrors.NewGameError("move dimension invalid", customerrors.ErrorCodeBadRequest)
	}
	if board.cells[move.RowNo][move.ColNo] != enums.TokenEmpty {
		return false, customerrors.NewGameError("move invalid as filled", customerrors.ErrorCodeBadRequest)
	}
	board.cells[move.RowNo][move.ColNo] = Token
	isWinning := board.strategy.RecordMove(move, Token)
	board.emptyCells--
	return isWinning, nil
}

func (board *SquareBoard) IsBoardFilled() bool {
	return board.emptyCells == 0
}

func (board *SquareBoard) Layout() string {
	layourStr := ""
	for rowIndex := 1; rowIndex <= board.size; rowIndex++ {
		for colIndex := 1; colIndex <= board.size; colIndex++ {
			layourStr += enums.Layout(board.cells[rowIndex][colIndex])
			if colIndex != board.size {
				layourStr += " |"
			}
		}
		layourStr += "\n"
	}
	return layourStr
}
