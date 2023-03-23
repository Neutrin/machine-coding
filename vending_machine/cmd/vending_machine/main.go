package main

import (
	"fmt"

	"github.com/vending_machine/internal"
	"github.com/vending_machine/internal/wallets"
)

func main() {
	var (
		slots  []*internal.Slot
		err    error
		change int
	)
	item := internal.NewItem("coca cola", 10)
	slots = append(slots, internal.NewSlot(item, 1))
	machine := internal.NewVendingMachine("tokopedia machine",
		slots, wallets.NewCoinWallet(0))

	if err = machine.AccpetCash(15); err == nil {
		if err = machine.ChooseSlot(1, 1); err == nil {
			change, err = machine.DispenseItem()
		}
	}

	if err == nil {
		fmt.Println(" purchase succesfull your change is = ", change)
	} else {
		fmt.Printf(" purchase failed  err = %+v\n", err)
	}
}
