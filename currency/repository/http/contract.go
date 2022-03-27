package http

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"coinconv/model"
)

var ErrNotOK = errors.New("http response status is not 200")

func toQueryParams(req model.ConvertReq) string {
	q := url.Values{}
	q.Add("amount", fmt.Sprintf("%f", req.Amount))
	q.Add("symbol", req.Symbol)
	q.Add("convert", strings.Join(req.Convert, ","))

	return q.Encode()
}

type httpRespDataQuote struct {
	Price       float64   `json:"price"`
	LastUpdated time.Time `json:"last_updated"`
}

type httpRespData struct {
	Quote map[string]httpRespDataQuote `json:"quote"`
}

type httpResp struct {
	Data map[string]httpRespData `json:"data"`
}
