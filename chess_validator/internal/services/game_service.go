package services

import (
	"fmt"

	"github.com/neutrin/chess_validator/internal/core/domain"
)

type GameService struct {
	game      *domain.Game
	moveCount int
}

func NewGameService(game *domain.Game) *GameService {
	return &GameService{
		game: game,
	}
}

func (service *GameService) Move(from, to *domain.Position) error {
	service.moveCount++
	if !service.game.Board.IsValidPos(from) || !service.game.Board.IsValidPos(to) {
		return fmt.Errorf(" invalid move 23")
	}
	curPeice := service.game.Board.Peice(from)
	if curPeice == nil {
		return fmt.Errorf(" invalid move 27")
	}
	player := service.game.Player()
	if player == nil {
		return fmt.Errorf(" invalid move 31")
	}
	if !curPeice.IsSameColor(player.Color) {
		return fmt.Errorf("invalid move 34")
	}
	err := curPeice.IsValidMove(to, from, service.game.Board, directionByColor(curPeice, service.moveCount))
	if err != nil {

		return err
	}
	if toPeice := service.game.Board.Peice(to); toPeice != nil {
		service.game.Board.KillPeice(toPeice, to)
	}
	service.game.Board.MovePeice(curPeice, from, to)
	service.game.NextIndex()
	return nil
}

func (service *GameService) Print() {
	service.game.Board.Print()
}
func directionByColor(peice domain.Peice, moveCount int) int {
	dir := 1
	if moveCount <= 2 {
		dir = 2
	}
	if peice.IsSameColor(domain.Black) {
		dir = dir * -1
	}
	return dir
}
