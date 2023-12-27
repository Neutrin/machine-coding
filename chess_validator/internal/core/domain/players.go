package domain

type Player struct {
	Name  string
	Color ColorEnum
}

func NewPlayer(name string, color ColorEnum) *Player {
	return &Player{
		Name:  name,
		Color: color,
	}
}
