package domain

type Peice interface {
	IsValidMove(to, from *Position, board *Board, move ...int) error
	Layout() string
	IsSameColor(color ColorEnum) bool
}
