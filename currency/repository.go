package currency

import (
	"context"
	
	"coinconv/model"
)

type ConvertRepository interface {
	Convert(ctx context.Context, req model.ConvertReq) (model.ConvertResp, error)
}
