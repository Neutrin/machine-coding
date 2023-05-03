package domains

type PaymentService interface {
	MakePayment(paymentDet interface{}) (PaymentResp, error)
}
