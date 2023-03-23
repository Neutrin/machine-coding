package internal

type State interface {
	AddAmount(amount int) error
	ChooseSlot(slotId int, quantity int) error
	DispenseItem() (int, error)
}
