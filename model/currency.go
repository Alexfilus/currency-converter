package model

type ConvertReq struct {
	Amount  float64
	Symbol  string
	Convert []string
}

type ConvertResp map[string]float64
