package main

import (
	"fmt"

	"github.com/neutrin/chess_validator/internal/core/domain"
	"github.com/neutrin/chess_validator/internal/services"
)

func main() {
	playerOne := domain.NewPlayer("player one", domain.White)
	playerTwo := domain.NewPlayer(" player two", domain.Black)
	board := domain.NewDefBoard(8)
	gameService := services.NewGameService(domain.NewGame(board, []*domain.Player{playerOne, playerTwo}))
	gameService.Print()
	fromPos := []string{"e2", "e7", "f1", "b8", "d1", "g8", "h5", "f8", "g7", "h8", "d8", "c6", "c4", "h8", "f2", "e5", "f7"}
	toPos := []string{"e4", "e5", "c4", "c6", "h5", "f6", "f7", "f7", "f7", "f7", "f7", "f7", "f7", "g8", "f4", "f4", "e8"}
	for index := range fromPos {
		from, _ := domain.NewPosition(fromPos[index])
		to, _ := domain.NewPosition(toPos[index])
		err := gameService.Move(from, to)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		gameService.Print()

	}

}
