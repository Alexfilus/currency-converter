package usecase

import (
	"context"

	"coinconv/currency"
	"coinconv/model"
)

type currencyUseCaseWithCommission struct {
	useCase    currency.UseCase
	commission float64
}

var _ currency.UseCase = (*currencyUseCaseWithCommission)(nil)

func NewWithCommission(useCase currency.UseCase, value float64) *currencyUseCaseWithCommission {
	return &currencyUseCaseWithCommission{useCase: useCase, commission: value}
}

func (c *currencyUseCaseWithCommission) Convert(ctx context.Context, amount float64, symbol string, convert []string) (model.ConvertResp, error) {
	resp, err := c.useCase.Convert(ctx, amount, symbol, convert)
	if err != nil {
		return nil, err
	}
	result := make(map[string]float64, len(resp))
	for k, v := range resp {
		result[k] = v * (1 + c.commission)
	}
	return result, nil
}
