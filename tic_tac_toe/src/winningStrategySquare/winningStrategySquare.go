package winningStrategySquare

type WinningStrategySquare interface {
	Intialize(size int, userCount int) error
	MakeMove(peice rune, row int, col int) bool
}
