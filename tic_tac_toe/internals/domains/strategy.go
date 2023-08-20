package domains

import (
	"github.com/tic_tac_toe/internals/domains/enums"
)

type WinningStrategy interface {
	RecordMove(move Move, Token enums.Token) bool
}
