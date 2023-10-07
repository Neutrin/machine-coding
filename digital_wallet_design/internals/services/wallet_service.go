package services

import "github.com/digital_wallet_design/internals/core/domains"

type WalletService struct{}

func NewWalletService() *WalletService {
	return &WalletService{}
}

func (w *WalletService) CreateWallet(balance float64) (*domains.Wallet, error) {
	amount, err := domains.NewAmount(balance)
	if err != nil {
		return nil, err
	}
	return domains.NewWallet(amount), err
}
