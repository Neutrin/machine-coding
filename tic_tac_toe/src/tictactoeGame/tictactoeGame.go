package tictactoeGame

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ticTacToe/src/ticTacToeBoard"
	"github.com/ticTacToe/src/user"
	"github.com/ticTacToe/src/winningStrategySquare"
)

type TicTacToeGame struct {
	users     []user.User
	board     ticTacToeBoard.TicTacToeBoard
	strategy  winningStrategySquare.WinningStrategySquare
	userIndex int
}

func (game *TicTacToeGame) Intialize(size int, players []string, boardType string,
	strategyType string) error {
	var (
		runeSlice   []rune
		playerSplit []string
	)

	game.board = ticTacToeBoard.GetTicTacToeBoard(boardType)
	if game.board == nil {
		return fmt.Errorf("invalid board type")
	}
	if err := game.board.Intialize(strconv.Itoa(size)); err != nil {
		return err
	}
	for index := range players {
		playerSplit = strings.Split(players[index], " ")
		if len(playerSplit) != 2 {
			return fmt.Errorf("invalid players")
		}
		if runeSlice = []rune(playerSplit[0]); len(runeSlice) <= 0 {
			return fmt.Errorf("invalid rune type")
		}

		game.users = append(game.users, user.GetUser(playerSplit[1], "abcd", runeSlice[0]))
	}
	game.strategy = &winningStrategySquare.WinningStrategySquareHashMapImpl{}
	game.strategy.Intialize(size, len(players))
	return nil
}

func (game *TicTacToeGame) Play(moves []string) {
	for _, curMove := range moves {
		if err := game.board.Move(curMove, game.users[game.userIndex]); err != nil {
			fmt.Printf("%s \n", err.Error())
			continue
		}
		fmt.Println(game.board.PrintBoard())
		moveSplit := strings.Split(curMove, " ")
		row, _ := strconv.Atoi(moveSplit[0])
		col, _ := strconv.Atoi(moveSplit[1])
		if isWinning :=
			game.strategy.MakeMove(game.users[game.userIndex].GetMark(), row, col); isWinning {
			fmt.Printf("%s won the game \n ", game.users[game.userIndex].GetName())
			return
		}
		game.userIndex = (game.userIndex + 1) % len(game.users)
	}
}
