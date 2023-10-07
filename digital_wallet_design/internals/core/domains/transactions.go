package domains

import (
	"time"

	"github.com/digital_wallet_design/internals/enums"
)

type Transaction struct {
	Account string
	Amount  *Amount
	Status  int64
	Type    string
	Date    string
}

func NewDebitTransaction(Account string, Amount *Amount) *Transaction {
	return &Transaction{
		Account: Account,
		Amount:  Amount,
		Status:  enums.StatusActive,
		Type:    enums.Debit,
		Date:    time.Now().Format("02-01-2006"),
	}
}

func NewCreditTransaction(Account string, Amount *Amount) *Transaction {
	return &Transaction{
		Account: Account,
		Amount:  Amount,
		Status:  enums.StatusActive,
		Type:    enums.Credit,
		Date:    time.Now().Format("02-01-2006"),
	}
}
