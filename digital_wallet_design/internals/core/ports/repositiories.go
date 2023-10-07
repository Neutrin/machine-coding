package ports

import "github.com/digital_wallet_design/internals/core/domains"

type AccountRepo interface {
	SaveAcount(account *domains.Account)
	Account(name string) (*domains.Account, error)
	Accounts() ([]*domains.Account, error)
}
