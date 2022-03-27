package usecase

import (
	"context"
	
	"coinconv/currency"
	"coinconv/model"
)

type currencyUseCase struct {
	repo currency.ConvertRepository
}

func New(repo currency.ConvertRepository) *currencyUseCase {
	return &currencyUseCase{repo: repo}
}

func (c *currencyUseCase) Convert(ctx context.Context, amount float64, symbol string, convert []string) (model.ConvertResp, error) {
	return c.repo.Convert(ctx, model.ConvertReq{
		Amount:  amount,
		Symbol:  symbol,
		Convert: convert,
	})
}
