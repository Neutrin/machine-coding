package domain

import (
	"fmt"
	"strings"
)

type EqualExpense struct {
}

func NewEqualExpense() Expense {
	return &EqualExpense{}
}

// Validate transaction
// Calculate split values
// Create transactions of same amount
func (expense *EqualExpense) Transactions(model ExpenseModel) ([]*Transaction, error) {
	var txn []*Transaction
	err := valEqualExpense(model)
	if err != nil {
		return txn, err
	}

	toUsersCount := len(model.To) - 1
	splitAmount := model.TotalAmount / float64(toUsersCount+1)
	txn = make([]*Transaction, 0)
	for _, OwedByUser := range model.To {
		if strings.Compare(OwedByUser, model.Giver) == 0 {
			continue
		}
		txn = append(txn, NewTransaction(OwedByUser, model.Giver, splitAmount))
	}
	return txn, nil
}

func valEqualExpense(model ExpenseModel) error {
	if model.TotalAmount < 0.0 {
		return fmt.Errorf(" amount should be greater than zero")
	}
	if len(model.To) <= 1 {
		return fmt.Errorf(" please add expense users")
	}
	return nil
}
