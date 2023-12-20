package domain

import "fmt"

type PercentageExpense struct {
}

func NewPercentageExpense() Expense {
	return &PercentageExpense{}
}

func (expense *PercentageExpense) Transactions(model ExpenseModel) ([]*Transaction, error) {
	var (
		txn []*Transaction
		err error
	)
	err = valPercentageTransaction(model)
	if err != nil {
		return txn, err
	}
	userCount := len(model.To) - 1
	for index := 0; index < userCount; index++ {
		txn = append(txn, NewTransaction(model.To[index], model.Giver, perAmount(model.Expenses[index], model.TotalAmount)))

	}
	return txn, err
}

func perAmount(per float64, amount float64) float64 {
	var perVal float64
	perVal = per * amount
	perVal = perVal / float64(100)

	return perVal
}
func valPercentageTransaction(model ExpenseModel) error {
	if model.TotalAmount < 0.0 {
		return fmt.Errorf(" total amount less than zero")
	}
	var per float64
	for _, curPer := range model.Expenses {
		per += curPer
	}
	if per != 100.0 {
		return fmt.Errorf(" percentage should be 100")
	}
	return nil
}
