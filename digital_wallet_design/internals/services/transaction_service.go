package services

import (
	"github.com/digital_wallet_design/internals/core/domains"
	"github.com/digital_wallet_design/internals/core/ports"
	customerrors "github.com/digital_wallet_design/internals/custom_errors"
)

type TransactionsService struct {
	repo ports.AccountRepo
}

func NewTransactionsService(repo ports.AccountRepo) *TransactionsService {
	return &TransactionsService{
		repo: repo,
	}
}
func (service *TransactionsService) DebitTransaction(
	from *domains.Account, to string, amount *domains.Amount) (*domains.Transaction, error) {
	if !validateAccountBalance(from.Wallet.Amount.Value, amount.Value) {
		return nil, customerrors.BalanceLowForTransaction
	}
	newAmount, err := domains.NewAmount(from.Wallet.Amount.Value - amount.Value)
	if err != nil {
		return nil, err
	}
	from.Wallet.Amount = newAmount
	return domains.NewDebitTransaction(to, amount), nil
}

func (service *TransactionsService) CreditTransaction(to *domains.Account, from string, amount *domains.Amount) (
	*domains.Transaction, error) {
	newAmount, err := domains.NewAmount(to.Wallet.Amount.Value + amount.Value)
	if err != nil {
		return nil, err
	}
	to.Wallet.Amount = newAmount
	return domains.NewCreditTransaction(from, amount), nil
}

func validateAccountBalance(curBalance float64, deductBalance float64) bool {
	return curBalance > deductBalance
}
