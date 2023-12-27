package domain

type ColorEnum int64

const (
	Black = iota + 1
	White
)

var colorStringRep = map[ColorEnum]string{
	Black: "Black",
	White: "White",
}

func (enum ColorEnum) String() string {
	return colorStringRep[enum]

}
