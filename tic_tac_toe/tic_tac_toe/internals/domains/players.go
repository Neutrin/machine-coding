package domains

import "github.com/tic_tac_toe/internals/domains/enums"

type Player struct {
	name  string
	token enums.Token
}

func NewPlayer(name string, token enums.Token) *Player {
	return &Player{
		name:  name,
		token: token,
	}
}

func (player *Player) Name() string {
	return player.name
}

func (p *Player) Token() enums.Token {
	return p.token
}
