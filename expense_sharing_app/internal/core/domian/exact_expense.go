package domain

import "fmt"

type ExactExpense struct {
}

func NewExactExpense() Expense {
	return &ExactExpense{}
}

func (expense *ExactExpense) Transactions(model ExpenseModel) ([]*Transaction, error) {
	var (
		txn []*Transaction
		err error
	)

	err = valExactExpense(model)
	if err != nil {
		return txn, err
	}

	totalUsers := len(model.To)
	for index := 0; index < totalUsers; index++ {
		txn = append(txn, NewTransaction(model.To[index], model.Giver, model.Expenses[index]))
	}
	return txn, err
}

func valExactExpense(model ExpenseModel) error {
	if model.TotalAmount < 0.0 {
		return fmt.Errorf(" total amount should be greater than zero")
	}
	var totalExpense float64
	for _, curExp := range model.Expenses {
		totalExpense += curExp
	}
	if totalExpense != model.TotalAmount {
		return fmt.Errorf(" total amount should be equal to added expenses")
	}
	return nil
}
