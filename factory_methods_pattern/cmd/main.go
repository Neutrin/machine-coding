package main

import (
	"fmt"

	"github.com/neutrin/factory_methods/internal/factories"
)

func main() {
	var err error
	factory := factories.NewPaymentFactoryMap()
	if err = factory.SetPaymentService("paypal"); err != nil {
		fmt.Println(" some error came = %+v \n", err)
	}
	factory.MakePayment("input_det")
	if err = factory.SetPaymentService("stripe"); err != nil {
		fmt.Println(" some error came = %+v \n", err)
	}
	factory.MakePayment("input_det")

}
