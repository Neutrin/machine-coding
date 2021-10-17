package winningStrategySquare

type WinningStrategySquareHashMapImpl struct {
	size             int
	rowUserMap       map[int]map[rune]int
	colUserMap       map[int]map[rune]int
	diagnalCount     map[rune]int
	antiDiagnalCount map[rune]int
}

func (strategy *WinningStrategySquareHashMapImpl) Intialize(size int, userCount int) error {
	strategy.size = size
	strategy.rowUserMap = make(map[int]map[rune]int)
	strategy.colUserMap = make(map[int]map[rune]int)
	strategy.diagnalCount = make(map[rune]int)
	strategy.antiDiagnalCount = make(map[rune]int)
	return nil

}

func (strategy *WinningStrategySquareHashMapImpl) MakeMove(peice rune, row int, col int) bool {
	count := 0
	if rowRuneMap, rowExists := strategy.rowUserMap[row]; rowExists {
		if runeCount, exists := rowRuneMap[peice]; exists {
			runeCount = runeCount + 1
			count = runeCount
			rowRuneMap[peice] = runeCount
		} else {
			rowRuneMap[peice] = 1
			count = 1
			strategy.rowUserMap[row] = rowRuneMap
		}
	} else {
		strategy.rowUserMap[row] = map[rune]int{
			peice: 1,
		}
		count = 1
	}
	if count == strategy.size {
		return true
	}
	count = 0

	if colRuneMap, colExist := strategy.colUserMap[col]; colExist {
		if runeCount, runeExists := colRuneMap[peice]; runeExists {
			runeCount++
			colRuneMap[peice] = runeCount
			count = runeCount
		} else {
			colRuneMap[peice] = 1
			count = 1
		}
	} else {
		strategy.colUserMap[col] = map[rune]int{
			peice: 1,
		}
	}
	if count == strategy.size {
		return true
	}

	if row == col {
		if diganlRune, exist := strategy.diagnalCount[peice]; exist {
			diganlRune++
			count = diganlRune
			strategy.diagnalCount[peice] = diganlRune
		} else {
			strategy.diagnalCount[peice] = 1
			count = 1
		}
	}
	if count == strategy.size {
		return true
	}

	curRow := 1
	curCol := strategy.size
	isAnti := false
	for curRow <= strategy.size && curCol >= 1 {
		if curRow == row && curCol == col {
			isAnti = true
			break
		}
		curRow++
		curCol--
	}
	if isAnti {
		count = 0
		if antiDiagnalCount, exists := strategy.antiDiagnalCount[peice]; exists {
			antiDiagnalCount++
			strategy.antiDiagnalCount[peice] = antiDiagnalCount
			count = antiDiagnalCount
		} else {
			strategy.antiDiagnalCount[peice] = 1
			count = 1
		}
		if count == strategy.size {
			return true
		}
	}
	return false
}
