package ticTacToeBoard

import "github.com/ticTacToe/src/user"

type TicTacToeBoard interface {
	Intialize(input string) error
	Move(input string, user user.User) error
	PrintBoard() string
}
