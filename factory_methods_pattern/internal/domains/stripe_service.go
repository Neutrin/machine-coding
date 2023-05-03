package domains

import "fmt"

type StripeService struct {
	auth StripeConfig
}

func NewStripeService(auth StripeConfig) PaymentService {
	return &StripeService{auth: auth}
}

func (service *StripeService) MakePayment(paymentDet interface{}) (PaymentResp, error) {
	//Unmashal all the details here
	fmt.Println(" made payment via stripe")
	return PaymentResp{}, nil
}
