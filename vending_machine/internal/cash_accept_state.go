package internal

import "fmt"

type CashAcceptState struct {
	machine *VendingMachine
}

func NewCashAcceptState(machine *VendingMachine) State {
	return &CashAcceptState{machine}
}

func (c *CashAcceptState) AddAmount(amount int) error {
	c.machine.Wallet().AddAmount(amount)

	c.machine.SetState(NewChooseItemState(c.machine))

	c.machine.txnDet.Amount = amount
	fmt.Println(" moving to choose item state")
	return nil
}

func (c *CashAcceptState) ChooseSlot(slotId int,
	quantity int) error {
	return fmt.Errorf(" invalid state")
}

func (c *CashAcceptState) DispenseItem() (int, error) {
	return 0, fmt.Errorf(" invalid state ")
}
