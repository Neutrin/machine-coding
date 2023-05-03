package domains

type PaymentResp struct {
	Id      string
	Message string
}

type PayPalConfig struct {
	Auth string
}

type StripeConfig struct {
	Auth string
}
