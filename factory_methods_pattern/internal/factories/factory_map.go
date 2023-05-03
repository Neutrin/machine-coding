package factories

import (
	"fmt"
	"sync"

	"github.com/neutrin/factory_methods/internal/domains"
)

var paymentFactoryMap map[string]domains.PaymentService

type PaymentFactoryMap struct {
	serviceName string
	service     domains.PaymentService
}

func intiFactoryMap() {
	var once sync.Once
	once.Do(func() {
		paymentFactoryMap = make(map[string]domains.PaymentService)
		paymentFactoryMap["paypal"] = domains.NewPayPalService(domains.PayPalConfig{Auth: "paypal_password"})
		paymentFactoryMap["stripe"] = domains.NewStripeService(domains.StripeConfig{Auth: "stripe_password"})
	})

}

func NewPaymentFactoryMap() PaymentFactory {
	intiFactoryMap()
	return &PaymentFactoryMap{}
}

func (factory *PaymentFactoryMap) MakePayment(det interface{}) (domains.PaymentResp, error) {
	return factory.service.MakePayment(det)
}

func (factory *PaymentFactoryMap) SetPaymentService(serviceName string) error {
	factory.serviceName = serviceName
	if service, exist := paymentFactoryMap[serviceName]; exist {
		factory.service = service
		return nil
	}
	return fmt.Errorf(" factory not registered = %s \n", serviceName)
}
