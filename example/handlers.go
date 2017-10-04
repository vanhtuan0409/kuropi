package main

import (
	"errors"

	"github.com/vanhtuan0409/kuropi"
)

func SuccessHandler(ctx kuropi.Context) {
	ctx.FastResponse("json", "sample data", nil)
}

func ErrorHandler(ctx kuropi.Context) {
	ctx.FastResponse("json", nil, errors.New("my custom error"))
}
