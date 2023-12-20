package domain

type Expense interface {
	Transactions(model ExpenseModel) ([]*Transaction, error)
}
