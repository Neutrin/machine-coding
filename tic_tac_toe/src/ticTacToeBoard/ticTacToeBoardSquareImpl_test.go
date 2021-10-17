package ticTacToeBoard

import (
	"reflect"
	"testing"

	"github.com/ticTacToe/src/user"
)

func Test_Intialize(t *testing.T) {
	tests := []struct {
		name     string
		expErr   bool
		expected TicTacToeSquareImpl
		args     string
	}{
		{
			name:   "Intialize with value positive value greater than zero",
			expErr: false,
			expected: TicTacToeSquareImpl{
				size:  10,
				board: getRuneBoard(10),
			},
			args: "10",
		},
		{
			name:     "input value is not decimal",
			expErr:   true,
			expected: TicTacToeSquareImpl{},
			args:     "abcd",
		},
		{
			name:     "input value is negetive",
			expErr:   true,
			expected: TicTacToeSquareImpl{},
			args:     "-10",
		},
		{
			name:     "Input size is zero",
			expErr:   true,
			expected: TicTacToeSquareImpl{},
			args:     "0",
		},
	}
	for _, curTest := range tests {
		t.Run(curTest.name, func(t *testing.T) {
			board := &TicTacToeSquareImpl{}
			err := board.Intialize(curTest.args)
			if (curTest.expErr == true && err == nil) || (curTest.expErr == false && err != nil) {
				t.Errorf("Wanted error = %v and got = %v", curTest.expErr, err)
			}
			if err == nil && !reflect.DeepEqual(*board, curTest.expected) {
				t.Errorf("wanted size = %d anf got = %d", curTest.expected, board.size)
			}
		})
	}
}

func getRuneBoard(size int) [][]rune {
	runeBoard := make([][]rune, size)
	for index := 0; index < size; index++ {
		runeBoard[index] = make([]rune, size)
	}
	return runeBoard
}

func Test_Move(t *testing.T) {
	type args struct {
		input string
		user  user.User
		board *TicTacToeSquareImpl
	}

	tests := []struct {
		name     string
		args     args
		expected TicTacToeSquareImpl
		wantErr  bool
	}{
		{
			name: "Positive move",
			args: args{
				input: "3 2",
				user:  user.GetUser("nitin", "nitin", rune('x')),
				board: &TicTacToeSquareImpl{
					size: 3,
					board: [][]rune{
						[]rune{rune(0), rune(0), rune(0)},
						[]rune{rune(0), rune(0), rune(0)},
						[]rune{rune(0), rune(0), rune(0)},
					},
				},
			},
			expected: TicTacToeSquareImpl{
				size: 3,
				board: [][]rune{
					[]rune{rune(0), rune(0), rune(0)},
					[]rune{rune(0), rune(0), rune(0)},
					[]rune{rune(0), rune('x'), rune(0)},
				},
			},
			wantErr: false,
		},
		{
			name: "Overlapping move",
			args: args{
				input: "3 2",
				user:  user.GetUser("nitin", "nitin", rune('x')),
				board: &TicTacToeSquareImpl{
					size: 3,
					board: [][]rune{
						[]rune{rune(0), rune(0), rune(0)},
						[]rune{rune(0), rune(0), rune(0)},
						[]rune{rune(0), rune('x'), rune(0)},
					},
				},
			},
			expected: TicTacToeSquareImpl{
				size: 3,
				board: [][]rune{
					[]rune{rune(0), rune(0), rune(0)},
					[]rune{rune(0), rune(0), rune(0)},
					[]rune{rune(0), rune('x'), rune(0)},
				},
			},
			wantErr: true,
		},
		{
			name: "out of bound move indexes",
			args: args{
				input: "4 4",
				user:  user.GetUser("nitin", "nitin", rune('x')),
				board: &TicTacToeSquareImpl{
					size: 3,
					board: [][]rune{
						[]rune{rune(0), rune(0), rune(0)},
						[]rune{rune(0), rune(0), rune(0)},
						[]rune{rune(0), rune(0), rune(0)},
					},
				},
			},
			expected: TicTacToeSquareImpl{
				size: 3,
				board: [][]rune{
					[]rune{rune(0), rune(0), rune(0)},
					[]rune{rune(0), rune(0), rune(0)},
					[]rune{rune(0), rune(0), rune(0)},
				},
			},
			wantErr: true,
		},
	}
	for _, curTest := range tests {
		t.Run(curTest.name, func(t *testing.T) {
			board := curTest.args.board
			err := board.Move(curTest.args.input, curTest.args.user)

			if (curTest.wantErr == true && err == nil) || (err != nil && curTest.wantErr == false) {
				t.Errorf("Wanted err = %v and got = %v", curTest.wantErr, err)
				return
			}
			if err == nil && !reflect.DeepEqual(*board, curTest.expected) {
				t.Errorf("Expcted = %+v and got = %+v", curTest.expected, *board)
			}
		})
	}
}

func Test_PrintBoard(t *testing.T) {
	tests := []struct {
		name     string
		args     *TicTacToeSquareImpl
		expected string
	}{
		{
			name: "Board with all filled with x",
			args: &TicTacToeSquareImpl{
				size: 3,
				board: [][]rune{
					[]rune{rune('x'), rune('x'), rune('x')},
					[]rune{rune('x'), rune('x'), rune('x')},
					[]rune{rune('x'), rune('x'), rune('x')},
				},
			},
			expected: "x x x\nx x x\nx x x\n",
		},
	}
	for _, curTests := range tests {
		t.Run(curTests.expected, func(t *testing.T) {
			layout := curTests.args.PrintBoard()
			if layout != curTests.expected {
				t.Errorf("wanted = %v \n  got = %v and the len of expected is = %d and got is %d", curTests.expected, layout, len(curTests.expected), len(layout))
			}

		})
	}
}
