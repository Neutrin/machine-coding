package domains

import (
	"github.com/tic_tac_toe/internals/domains/customerrors"
	"github.com/tic_tac_toe/internals/domains/enums"
)

type Board interface {
	AddMove(move Move, Token enums.Token) (bool, *customerrors.GameError)
	IsBoardFilled() bool
	Layout() string
}
