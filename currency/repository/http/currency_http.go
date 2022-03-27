package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	
	"coinconv/currency"
	"coinconv/model"
)

type httpCurrencyRepository struct {
	client *http.Client
	apiKey string
}

func New(timeout time.Duration, apiKey string) currency.ConvertRepository {
	return &httpCurrencyRepository{
		client: &http.Client{Timeout: timeout},
		apiKey: apiKey,
	}
}

func (repo *httpCurrencyRepository) Convert(ctx context.Context, req model.ConvertReq) (model.ConvertResp, error) {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://sandbox-api.coinmarketcap.com/v2/tools/price-conversion", nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Accepts", "application/json")
	httpReq.Header.Add("X-CMC_PRO_API_KEY", repo.apiKey)
	httpReq.URL.RawQuery = toQueryParams(req)
	
	resp, err := repo.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()
	if resp.StatusCode != http.StatusOK {
		return nil, ErrNotOK
	}
	respObj := new(httpResp)
	err = json.NewDecoder(resp.Body).Decode(respObj)
	if err != nil {
		return nil, err
	}
	
	quote := respObj.Data[req.Symbol].Quote
	result := make(map[string]float64, len(req.Convert))
	for _, part := range req.Convert {
		result[part] = quote[part].Price
	}
	return result, nil
}
