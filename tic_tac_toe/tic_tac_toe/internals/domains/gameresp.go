package domains

import (
	"fmt"
)

type GameResp struct {
	Continue bool
	Error    error
	Msg      string
}

func WonGameResp(player Player) GameResp {
	return GameResp{
		Msg: fmt.Sprintf("%s won the game", player.Name()),
	}
}

func DrawGameResp() GameResp {
	return GameResp{
		Msg: "GAME OVER",
	}
}

func InvalidMoveResp(err error) GameResp {
	return GameResp{
		Continue: true,
		Error:    err,
		Msg:      "INVALID MOVE",
	}
}

func ContinueGameResp() GameResp {
	return GameResp{
		Continue: true,
	}
}
