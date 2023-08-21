package domains

import (
	"github.com/tic_tac_toe/internals/domains/enums"
)

const (
	row         = "row"
	col         = "col"
	diagnal     = "diag"
	antiDiagnal = "antidiagnal"
)

type SquareWinningStrategy struct {
	size                 int
	patternCount         map[enums.Token]map[string][]int
	antiDiganalCordinate map[int]int
}

func NewSquareWinningStrategy(size int, tokens []enums.Token) WinningStrategy {
	strategy := &SquareWinningStrategy{size: size}
	strategy.patternCount = make(map[enums.Token]map[string][]int)
	for _, curToken := range tokens {
		strategy.patternCount[curToken] = map[string][]int{
			row:         make([]int, size+1),
			col:         make([]int, size+1),
			diagnal:     make([]int, 1),
			antiDiagnal: make([]int, 1),
		}

	}
	strategy.antiDiganalCordinate = antiDiagnalCordinate(size)
	return strategy
}

func (strategy *SquareWinningStrategy) RecordMove(move Move, Token enums.Token) bool {
	tokenMap, exists := strategy.patternCount[Token]
	if !exists {
		return false
	}

	tokenMap[row][move.RowNo] += 1
	tokenMap[col][move.ColNo] += 1
	if strategy.isDiagnal(move.RowNo, move.ColNo) {
		tokenMap[diagnal][0] += 1
	}
	if strategy.isAntiDiagnal(move.RowNo, move.ColNo) {
		tokenMap[antiDiagnal][0] += 1
	}
	if tokenMap[row][move.RowNo] == strategy.size || tokenMap[col][move.ColNo] == strategy.size ||
		tokenMap[diagnal][0] == strategy.size || tokenMap[antiDiagnal][0] == strategy.size {
		return true
	}
	return false

}

func antiDiagnalCordinate(size int) map[int]int {
	antiDiganalCordinate := make(map[int]int)
	for rowIndex, colIndex := 1, size; rowIndex <= size && colIndex >= 1; rowIndex, colIndex = rowIndex+1, colIndex-1 {
		antiDiganalCordinate[rowIndex] = colIndex
	}
	return antiDiganalCordinate
}

func (strategy *SquareWinningStrategy) isDiagnal(row int, col int) bool {
	return row == col
}

func (strategy *SquareWinningStrategy) isAntiDiagnal(row int, col int) bool {
	if colNext, exist := strategy.antiDiganalCordinate[row]; exist {
		if col == colNext {
			return true
		}
	}
	return false
}
