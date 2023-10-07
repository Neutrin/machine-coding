package domains

import (
	"time"

	"github.com/digital_wallet_design/internals/enums"
)

type Account struct {
	Name      string
	Wallet    *Wallet
	Txns      []*Transaction
	CreatedAt time.Time
	Status    int64
}

func NewAccount(name string, wallet *Wallet) *Account {
	return &Account{
		Name:      name,
		Wallet:    wallet,
		Txns:      make([]*Transaction, 0),
		CreatedAt: time.Now(),
		Status:    enums.StatusActive,
	}
}
