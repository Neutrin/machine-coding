package factories

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/neutrin/factory_methods/internal/domains"
	"github.com/neutrin/factory_methods/internal/mocks"
)

func Test_MakePayment(t *testing.T) {
	type args struct {
		serviceName string
	}

	tests := []struct {
		name string
		mock func(paymentMock *mocks.MockPaymentService)
	}{
		{
			name: "Paypal service",
			mock: func(paymentMock *mocks.MockPaymentService) {
				paymentMock.EXPECT().MakePayment("paypal").Return(domains.PaymentResp{}, nil)
			},
		},
	}
	for _, curTests := range tests {
		t.Run(curTests.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockPaymentService := mocks.NewMockPaymentService(ctrl)
			curTests.mock(mockPaymentService)
			factory := PaymentFactoryMap{service: mockPaymentService}
			factory.MakePayment("paypal")
		})
	}
}
