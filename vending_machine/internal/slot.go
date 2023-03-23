package internal

var slotId int

type Slot struct {
	id       int
	item     *Item
	quantity int
}

func NewSlot(item *Item, quantity int) *Slot {
	slotId++
	return &Slot{
		id:       slotId,
		item:     item,
		quantity: quantity,
	}
}

//check if slot is having sufficient quantity or not
func (s *Slot) IsOOS(quantity int) bool {
	return s.quantity < quantity
}

//BuySlot : will buy the slot as quantity  is available
func (s *Slot) BuySlot(quantity int) {
	s.quantity -= quantity
}

//IsAmountValid : will check amount is valid to buy slot or not
func (s *Slot) IsAmountValid(buyQuantity, buyCash int) bool {
	return buyCash >= s.item.GetTotalPrice(buyQuantity)
}

func (s *Slot) ItemPrice(quantity int) int {
	return s.item.GetTotalPrice(quantity)
}
