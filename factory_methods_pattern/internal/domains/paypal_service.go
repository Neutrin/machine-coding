package domains

import "fmt"

type PayPalService struct {
	config PayPalConfig
}

func NewPayPalService(config PayPalConfig) *PayPalService {
	return &PayPalService{config}
}

func (service *PayPalService) MakePayment(paymentDet interface{}) (PaymentResp, error) {
	//Fetch details and unmarsh for the make payment
	fmt.Println("made payment via Paypal")
	return PaymentResp{}, nil
}
