package factories

import (
	"github.com/neutrin/factory_methods/internal/domains"
)

type PaymentFactory interface {
	SetPaymentService(serviceName string) error
	MakePayment(det interface{}) (domains.PaymentResp, error)
}
