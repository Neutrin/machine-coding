package repositiories

import (
	"fmt"

	"github.com/digital_wallet_design/internals/core/domains"
)

type AccountRepoImpl struct {
	mp map[string]domains.Account
}

func NewAccountRepoImpl() *AccountRepoImpl {
	return &AccountRepoImpl{
		mp: make(map[string]domains.Account),
	}
}

/*
type AccountRepo interface {
	SaveAcount(account *domains.Account)
	Account(name string) (*domains.Account, error)
	Accounts() ([]*domains.Account, error)
}


*/

func (repo *AccountRepoImpl) SaveAcount(account *domains.Account) {
	repo.mp[account.Name] = domains.Account{
		Name:      account.Name,
		Wallet:    account.Wallet,
		Txns:      account.Txns,
		CreatedAt: account.CreatedAt,
		Status:    account.Status,
	}
}

func (repo *AccountRepoImpl) Account(name string) (*domains.Account, error) {
	accountDB, contains := repo.mp[name]
	if !contains {
		return nil, fmt.Errorf(" account with name not found")
	}
	return &domains.Account{
		Name:      accountDB.Name,
		Wallet:    accountDB.Wallet,
		Txns:      accountDB.Txns,
		CreatedAt: accountDB.CreatedAt,
		Status:    accountDB.Status,
	}, nil
}

func (repo *AccountRepoImpl) Accounts() ([]*domains.Account, error) {
	accounts := make([]*domains.Account, 0)
	for curName := range repo.mp {
		curAccountNew, _ := repo.Account(curName)
		accounts = append(accounts, curAccountNew)
	}
	if len(accounts) == 0 {
		fmt.Errorf(" no account found ")
	}
	return accounts, nil
}
