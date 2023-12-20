package domain

type ExpenseModel struct {
	Giver       string
	To          []string
	TotalAmount float64
	//todo
	Type     ExpenseType //enum
	Expenses []float64
}

func NewExpense(giver string, to []string, totalAmount float64, expenseType ExpenseType, expenses []float64) ExpenseModel {
	return ExpenseModel{
		Giver:       giver,
		To:          to,
		TotalAmount: totalAmount,
		Type:        expenseType,
		Expenses:    expenses,
	}

}
