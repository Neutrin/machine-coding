package domain

type Transaction struct {
	OwedBy string
	OwedTo string
	Amount float64
}

func NewTransaction(owedBy, owedTo string, amount float64) *Transaction {
	return &Transaction{
		OwedBy: owedBy,
		OwedTo: owedTo,
		Amount: amount,
	}
}
