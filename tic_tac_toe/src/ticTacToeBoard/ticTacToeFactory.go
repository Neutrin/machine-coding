package ticTacToeBoard

func GetTicTacToeBoard(boardTyp string) (board TicTacToeBoard) {

	switch boardTyp {
	case "square":
		board = &TicTacToeSquareImpl{}
	default:
		board = nil
	}
	return

}
