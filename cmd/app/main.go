package main

import (
	"context"
	"log"
	"os"
	"time"
	
	"coinconv/currency/delivery/cli"
	httpCurrencyRepo "coinconv/currency/repository/http"
	"coinconv/currency/usecase"
)

const apiKey = "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c" // should be in config
const timeout = 3 * time.Second                       // should be in config

func main() {
	currencyRepo := httpCurrencyRepo.New(timeout, apiKey)
	currencyUseCase := usecase.New(currencyRepo)
	handler := cli.New(currencyUseCase, os.Stdout)
	err := handler.Convert(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
}
