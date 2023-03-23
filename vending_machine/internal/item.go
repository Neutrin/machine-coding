package internal

var itemId int

type Item struct {
	id    int
	name  string
	price int
}

//NewItem : this will return new item
func NewItem(name string, price int) *Item {
	itemId++
	return &Item{
		id:    itemId,
		name:  name,
		price: price,
	}
}

//Id() : this will return id of the item
func (i *Item) Id() int {
	return i.id
}

//GetTotalPrice : this will return total price for the product
func (i *Item) GetTotalPrice(quantity int) int {
	return quantity * i.price
}
