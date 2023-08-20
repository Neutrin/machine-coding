package domains

// TODO : read about singleton pattern here how to use
type Move struct {
	RowNo int
	ColNo int
}

func NewMove(rowNo int, colNo int) *Move {
	return &Move{RowNo: rowNo, ColNo: colNo}
}
