package enums

type Token int

const (
	Empty = iota
	Cross
	Circle
)

var tokenIndex = []Token{Empty, Cross, Circle}

// Layout : This will return layout of token
func Layout(token Token) string {
	var lay string
	switch token {
	case Cross:
		lay = "X"
	case Circle:
		lay = "O"
	default:
		lay = " "
	}
	return lay
}

func TokenByIndex(index int) Token {
	if index < len(tokenIndex) {
		return tokenIndex[index]
	}
	return Empty
}
