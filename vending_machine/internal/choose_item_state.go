package internal

import "fmt"

type ChooseItemState struct {
	machine *VendingMachine
}

func NewChooseItemState(vending *VendingMachine) State {
	return &ChooseItemState{vending}
}

/*
AddAmount(amount int) error
	ChooseSlot(slotId int, quantity int, amount int ) error
	DispenseItem() error

*/

func (itemState *ChooseItemState) AddAmount(amount int) error {
	return fmt.Errorf(" invalid state")
}

func (itemState *ChooseItemState) ChooseSlot(slotId,
	quantity int) error {
	var err error
	//quantity := 1
	if err = itemState.machine.SlotManager().ValSlotPurchase(slotId, quantity,
		itemState.machine.txnDet.Amount); err != nil {
		return err
	}
	amountOne := itemState.machine.txnDet.Amount
	amountTwo := itemState.machine.SlotManager().Slot(slotId).ItemPrice(quantity)
	fmt.Printf(" amount one = %d and amount two = %d", amountOne, amountTwo)
	changeAmount := itemState.machine.txnDet.Amount -
		itemState.machine.SlotManager().Slot(slotId).ItemPrice(quantity)
	if err = itemState.machine.Wallet().DispenseAmount(changeAmount); err != nil {
		return err
	}
	itemState.machine.txnDet.SlotId = slotId
	itemState.machine.txnDet.Quantity = quantity
	itemState.machine.txnDet.Change = changeAmount
	//TODO : set next state for the transaction
	fmt.Println(" moving to dispense state")
	itemState.machine.SetState(NewDispenseItemState(itemState.machine))
	return nil

}

func (itemState *ChooseItemState) DispenseItem() (int, error) {
	return 0, fmt.Errorf(" not a valid state")
}
