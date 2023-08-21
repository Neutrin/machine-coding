package enums

type Token int

const (
	TokenEmpty = iota
	TokenCross
	TokenCircle
)

var tokenIndex = []Token{TokenEmpty, TokenCross, TokenCircle}

// Layout : This will return layout of token
func Layout(token Token) string {
	var lay string
	switch token {
	case TokenCross:
		lay = "X"
	case TokenCircle:
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
	return TokenEmpty
}
