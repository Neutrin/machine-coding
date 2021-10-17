package ticTacToeBoard

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ticTacToe/src/user"
)

type TicTacToeSquareImpl struct {
	size  int
	board [][]rune
}

func (board *TicTacToeSquareImpl) Intialize(input string) error {
	if size, err := strconv.Atoi(input); err != nil {
		return fmt.Errorf("input size should be decimal")
	} else {
		if size <= 0 {
			return fmt.Errorf("Input size should be greater than zero")
		}
		board.size = size
	}
	board.board = make([][]rune, board.size)
	for index := 0; index < board.size; index++ {
		board.board[index] = make([]rune, board.size)
	}
	return nil
}

func (board *TicTacToeSquareImpl) Move(move string, user user.User) error {

	if movesSplit := strings.Split(move, " "); len(movesSplit) == 2 {

		if row, err := strconv.Atoi(movesSplit[0]); err == nil {

			if col, err := strconv.Atoi(movesSplit[1]); err == nil {

				if row <= board.size && row >= 1 {

					if col <= board.size && col >= 1 {
						if board.board[row-1][col-1] != rune(0) {
							return fmt.Errorf("This is already marked")
						}
						board.board[row-1][col-1] = user.GetMark()
					} else {
						return fmt.Errorf("Invalid column = %d", col)
					}
				} else {

					return fmt.Errorf("Invalid row = %d", row)

				}
			} else {
				return fmt.Errorf("Invalid column = %s", movesSplit[1])
			}
		} else {
			return fmt.Errorf("Invalid row = %s", movesSplit[0])
		}
	} else {
		return fmt.Errorf("Invalid move = %s", move)
	}

	return nil
}

func (board *TicTacToeSquareImpl) PrintBoard() string {
	var (
		size   = board.size
		layout string
		params []interface{}
	)
	fmt.Println("here the length is")
	for curRow := 0; curRow < size; curRow++ {
		for curCol := 0; curCol < size; curCol++ {
			layout += "%s "
			if board.board[curRow][curCol] == rune(0) {
				params = append(params, "-")
			} else {
				params = append(params, string(board.board[curRow][curCol]))
			}
		}
		layout = strings.Trim(layout, " ")
		layout += "\n"
	}
	return fmt.Sprintf(layout, params...)
}
