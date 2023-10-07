package domains

import (
	"time"

	"github.com/digital_wallet_design/internals/enums"
)

type Wallet struct {
	Amount *Amount
	//TODO : enums can be created later
	Status    int64
	createdAt time.Time
}

func NewWallet(amount *Amount) *Wallet {
	return &Wallet{
		Amount:    amount,
		Status:    enums.StatusActive,
		createdAt: time.Now(),
	}
}
