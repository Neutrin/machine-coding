package customerrors

import "errors"

var (
	AmountLessThanMin        = errors.New("amount less than minimum")
	AccountNotFound          = errors.New(" account not found")
	BalanceLowForTransaction = errors.New(" balance low for txn")
)
