package wallets

type Wallet interface {
	CheckBalance() int
	AddAmount(amount int)
	DispenseAmount(amount int) error
}
