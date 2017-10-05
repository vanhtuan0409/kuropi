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

func RandomHandler(ctx kuropi.Context) {
	appNumber := ctx.FastGetInstance("appRandomNumber").(int)
	requestNumbder := ctx.FastGetInstance("requestRandomNumber").(int)
	ctx.FastResponse("json", map[string]interface{}{
		"app":     appNumber,
		"request": requestNumbder,
	}, nil)
}
