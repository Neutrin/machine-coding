package main

import (
	"bufio"
	"fmt"
	"os"

	inputPort "github.com/tic_tac_toe/internals/ports"

	"github.com/tic_tac_toe/internals/domains"

	"github.com/tic_tac_toe/internals/domains/enums"
)

func main() {
	strategy := domains.NewSquareWinningStrategy(3, []enums.Token{enums.TokenCircle, enums.TokenCross})
	board, err := domains.NewSquareBoard(3, strategy)
	if err != nil {
		fmt.Println(err.Error())
	}
	gameService := domains.NewGameService(board)
	scanner := bufio.NewScanner(os.Stdin)
	console := inputPort.NewConsole(gameService, scanner, 2)
	console.StartGame()

}
