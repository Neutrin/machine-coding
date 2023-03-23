package wallets

import "fmt"

type CoinWallet struct {
	//Here will be a driver to vending machine ches t wwhih has all the denominations
	amount int
}

func NewCoinWallet(amount int) Wallet {
	return &CoinWallet{amount: amount}
}
func (c *CoinWallet) CheckBalance() int {
	return c.amount
}

func (c *CoinWallet) AddAmount(amount int) {
	c.amount += amount

}

func (c *CoinWallet) DispenseAmount(amount int) error {
	if c.amount < amount {
		return fmt.Errorf(" amount not available")
	}
	c.amount -= amount
	return nil
}
