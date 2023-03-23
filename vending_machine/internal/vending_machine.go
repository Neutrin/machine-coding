package internal

import (
	"fmt"
	"sync"

	"github.com/vending_machine/internal/wallets"
)

type VendingMachine struct {
	name        string
	slotManager *SlotManager
	state       State
	txnDet      TxnDet
	wallet      wallets.Wallet
}

func NewVendingMachine(name string, slots []*Slot, wallet wallets.Wallet) *VendingMachine {
	var (
		once    = &sync.Once{}
		machine *VendingMachine
	)
	once.Do(func() {
		machine = &VendingMachine{
			name:        name,
			slotManager: NewSlotManager(slots),

			txnDet: TxnDet{},
			wallet: wallet,
		}
		fmt.Println(" intialised machine with state cash accept")
		machine.SetState(NewCashAcceptState(machine))
	})
	return machine
}

func (v *VendingMachine) Wallet() wallets.Wallet {
	return v.wallet
}

func (v *VendingMachine) SetState(state State) {
	v.state = state
}

func (v *VendingMachine) SlotManager() *SlotManager {
	return v.slotManager
}

func (v *VendingMachine) Reset() {
	v.txnDet = TxnDet{}

}

func (v *VendingMachine) AccpetCash(amount int) error {
	return v.state.AddAmount(amount)
}

func (v *VendingMachine) ChooseSlot(
	slotId int, quantity int) error {
	return v.state.ChooseSlot(slotId, 1)
}

func (v *VendingMachine) DispenseItem() (int, error) {
	return v.state.DispenseItem()
}
