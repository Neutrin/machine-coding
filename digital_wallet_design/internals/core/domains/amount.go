package domains

import (
	customerrors "github.com/digital_wallet_design/internals/custom_errors"
)

const minAmount = 0.0001

type Amount struct {
	Value float64
}

func NewAmount(value float64) (*Amount, error) {
	if value < minAmount {
		//TODO : error
		return &Amount{}, customerrors.AmountLessThanMin
	}
	return &Amount{
		Value: value,
	}, nil
}
