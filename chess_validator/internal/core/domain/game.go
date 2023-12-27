package domain

type Game struct {
	Board   *Board
	Players []*Player
	index   int
}

func NewGame(board *Board, players []*Player) *Game {
	return &Game{
		Board:   board,
		Players: players,
		index:   0,
	}
}

func (game *Game) NextIndex() {
	game.index = (game.index + 1) % len(game.Players)
}

func (game *Game) Player() *Player {
	return game.Players[game.index]
}
