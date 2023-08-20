package domains

import (
	"github.com/tic_tac_toe/internals/domains/customerrors"
	"github.com/tic_tac_toe/internals/domains/enums"
)

type GameService struct {
	players    []*Player
	playerTurn int
	status     enums.StatusEnum
	board      Board
}

func NewGameService(board Board) *GameService {

	return &GameService{
		status: enums.Running,
		board:  board,
	}
}

func (game *GameService) AddPlayers(players []*Player, playerTurn int) error {
	if len(players) == 0 {
		return customerrors.NewGameError("no player error", customerrors.ErrorCodeBadRequest)
	}
	if playerTurn < 0 || playerTurn >= len(players) {
		return customerrors.NewGameError("invalid playe turn", customerrors.ErrorCodeBadRequest)
	}
	game.players = players
	game.playerTurn = playerTurn
	return nil
}

func (game *GameService) PlayMove(move Move) GameResp {
	isWinning, err := game.board.AddMove(move, game.players[game.playerTurn].Token())
	if err != nil {
		return InvalidMoveResp(err)
	}
	if isWinning {
		game.status = enums.Win
		return WonGameResp(*game.players[game.playerTurn])
	}
	if game.board.IsBoardFilled() {
		game.status = enums.Completed
		return DrawGameResp()
	}
	game.playerTurn = (game.playerTurn + 1) % len(game.players)
	return ContinueGameResp()

}

func (game *GameService) PrintLayout() string {
	return game.board.Layout()
}
