package customerrors

const (
	ErrorCodeBadRequest               = 400
	ErrorCodeBoardIntialisationFailed = 401
)

type GameError struct {
	Code int
	Desc string
}

func (GameError *GameError) Error() string {
	return GameError.Desc
}

func NewGameError(desc string, Code int) *GameError {
	return &GameError{Code: Code, Desc: desc}
}
