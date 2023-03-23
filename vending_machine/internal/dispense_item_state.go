package internal

import "fmt"

/*

AddAmount(amount int) error
	ChooseSlot(slotId int, quantity int, amount int) error
	DispenseItem() error

*/
type DispenseItemState struct {
	machine *VendingMachine
}

func NewDispenseItemState(machine *VendingMachine) State {
	return &DispenseItemState{machine}
}

func (state *DispenseItemState) AddAmount(amount int) error {
	return fmt.Errorf(" invalid state")
}

func (state *DispenseItemState) ChooseSlot(slotId, quantity int) error {
	return fmt.Errorf(" invalid state")
}

func (state *DispenseItemState) DispenseItem() (int, error) {
	var (
		err          error
		changeAmount int
	)
	//Update slot details
	state.machine.SlotManager().Slot(state.machine.txnDet.SlotId).
		BuySlot(state.machine.txnDet.Quantity)
	//This will return if there is any change or not
	if err = state.machine.Wallet().DispenseAmount(state.machine.txnDet.Change); err == nil {

		changeAmount = state.machine.txnDet.Change
		state.machine.Reset()
		state.machine.SetState(NewCashAcceptState(state.machine))
	}
	return changeAmount, err
}
