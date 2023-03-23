package internal

import "fmt"

type SlotManager struct {
	slotMap map[int]*Slot
	slots   []*Slot
}

func NewSlotManager(slots []*Slot) *SlotManager {
	slotMap := make(map[int]*Slot)
	for _, curSlot := range slots {
		slotMap[curSlot.id] = curSlot
	}
	return &SlotManager{
		slotMap: slotMap,
		slots:   slots,
	}

}

func (s *SlotManager) ValSlotPurchase(slotId int, quantity int, amount int) error {
	var (
		curSlot *Slot
		present bool
	)
	if curSlot, present = s.slotMap[slotId]; !present {
		return fmt.Errorf(" not valid slot")
	}
	if curSlot.IsOOS(quantity) {
		return fmt.Errorf(" curSlot is oos")
	}
	if !curSlot.IsAmountValid(quantity, amount) {
		return fmt.Errorf(" amount not valid pls start again ")
	}

	return nil
}

func (s *SlotManager) BuySlot(slotId int, quantity int) {
	var curSlot *Slot
	curSlot = s.slotMap[slotId]
	curSlot.BuySlot(quantity)

}

func (s *SlotManager) Slot(slotId int) *Slot {
	var slot *Slot
	slot = s.slotMap[slotId]
	return slot
}
