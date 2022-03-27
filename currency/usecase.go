package currency

import (
	"context"
	
	"coinconv/model"
)

type UseCase interface {
	Convert(ctx context.Context, amount float64, symbol string, convert []string) (model.ConvertResp, error)
}
