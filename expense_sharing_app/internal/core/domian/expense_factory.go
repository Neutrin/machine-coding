package domain

func ExpenseByType(typ ExpenseType) Expense {

	switch typ {
	case Equal:

		return NewEqualExpense()
	case Percentage:

		return NewPercentageExpense()
	case Exact:

		return NewExactExpense()
	}

	return nil
}
