package services

import (
	"fmt"

	"github.com/digital_wallet_design/internals/core/domains"
	"github.com/digital_wallet_design/internals/core/ports"
)

type AccountService struct {
	repo      ports.AccountRepo
	walletSvc *WalletService
	txnSrvc   *TransactionsService
}

type AccountCreatedResp struct {
	account *domains.Account
}

type AccountStatementResp struct {
	account *domains.Account
}

func (resp AccountStatementResp) ToString() string {
	res := ""
	res += fmt.Sprintf("Summary for account number: %s \n", resp.account.Name)
	res += fmt.Sprintf("Current Balance: %f \n", resp.account.Wallet.Amount.Value)
	res += fmt.Sprintf("Your Transaction History\n")
	for _, curTxn := range resp.account.Txns {
		res += fmt.Sprintf(" from = %s amount = %f, type = %s, date = %s \n",
			curTxn.Account, curTxn.Amount.Value, curTxn.Type, curTxn.Date)
	}
	return res
}

func (ac AccountCreatedResp) ToString() string {
	return fmt.Sprintf(" account of user %s with balance created %f", ac.account.Name,
		ac.account.Wallet.Amount.Value)
}

func NewAccountService(repo ports.AccountRepo) *AccountService {
	return &AccountService{
		repo:      repo,
		walletSvc: NewWalletService(),
		txnSrvc:   NewTransactionsService(repo),
	}
}

func (service *AccountService) NewAccount(name string, balance float64) (AccountCreatedResp, error) {
	var resp AccountCreatedResp
	amt, err := domains.NewAmount(balance)
	if err != nil {
		return resp, err
	}
	act := domains.NewAccount(name, domains.NewWallet(amt))
	service.repo.SaveAcount(act)
	return AccountCreatedResp{
		account: act,
	}, nil
}

func (service *AccountService) TransferMoney(from, to string, amountVal float64) (string, error) {
	fromAccount, err := service.repo.Account(from)
	if err != nil {
		return "", err
	}
	toAccount, err := service.repo.Account(to)
	if err != nil {
		return "", err
	}
	amount, err := domains.NewAmount(amountVal)
	if err != nil {
		return "", err
	}
	dbTxn, err := service.txnSrvc.DebitTransaction(fromAccount, toAccount.Name, amount)
	if err != nil {
		return "", err
	}
	crTxn, err := service.txnSrvc.CreditTransaction(toAccount, fromAccount.Name, amount)
	if err != nil {
		return "", err
	}
	fromAccount.Txns = append(fromAccount.Txns, dbTxn)
	toAccount.Txns = append(toAccount.Txns, crTxn)
	if checkOffer(fromAccount, toAccount) {
		offAmount, err := domains.NewAmount(10)
		if err != nil {
			return "", err
		}
		offerTxn, err := service.txnSrvc.CreditTransaction(toAccount, "offer", offAmount)
		if err != nil {
			return "", err
		}
		toAccount.Txns = append(toAccount.Txns, offerTxn)
		offerTxn, err = service.txnSrvc.CreditTransaction(fromAccount, "offer", offAmount)
		if err != nil {
			return "", err
		}
		fromAccount.Txns = append(fromAccount.Txns, offerTxn)
	}
	service.repo.SaveAcount(fromAccount)
	service.repo.SaveAcount(toAccount)
	return "Transfer succesfull", nil
}

func (service *AccountService) AccountStatement(name string) (AccountStatementResp, error) {
	curAccount, err := service.repo.Account(name)
	if err != nil {
		return AccountStatementResp{}, err
	}
	return AccountStatementResp{curAccount}, nil

}

func (service *AccountService) Overview() ([]string, error) {
	accounts, err := service.repo.Accounts()
	if err != nil {
		return []string{}, err
	}
	accountResps := make([]string, 0)
	for _, curAccount := range accounts {
		accountResps = append(accountResps, fmt.Sprintf("Balance for account = %s = %f", curAccount.Name, curAccount.Wallet.Amount.Value))
	}
	return accountResps, nil
}

func checkOffer(accountOne *domains.Account, accountTwo *domains.Account) bool {
	return accountOne.Wallet.Amount.Value == accountTwo.Wallet.Amount.Value
}
