package cli

import (
	"context"
	"flag"
	"fmt"
	"io"
	"strconv"
	"strings"
	
	"coinconv/currency"
)

type Handler struct {
	useCase currency.UseCase
	output  io.Writer
}

func New(useCase currency.UseCase, output io.Writer) *Handler {
	return &Handler{
		useCase: useCase,
		output:  output,
	}
}

func (u *Handler) Convert(ctx context.Context) error {
	flag.Parse()
	args := flag.Args()
	if len(args) < 3 {
		return ErrNotEnough
	}
	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return err
	}
	symbol := args[1]
	convert := args[2]
	convertParts := strings.Split(convert, ",")
	
	resp, err := u.useCase.Convert(ctx, amount, symbol, convertParts)
	if err != nil {
		return err
	}
	for _, part := range convertParts {
		_, err = fmt.Fprintln(u.output, part, resp[part])
		if err != nil {
			return err
		}
	}
	return nil
}
