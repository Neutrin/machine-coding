package domain

type ExpenseType int64

const (
	Equal = iota + 1
	Percentage
	Exact
)

var stringRep map[ExpenseType]string = map[ExpenseType]string{
	Equal:      "equal",
	Percentage: "percentage",
	Exact:      "exact",
}

func (typ ExpenseType) String() string {
	rep, exists := stringRep[typ]
	if exists {
		return rep
	}
	return ""
}
