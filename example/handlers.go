package main

import (
	"errors"
	"fmt"

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
	fmt.Println("app", appNumber)
	requestNumbder := ctx.FastGetInstance("requestRandomNumber").(int)
	fmt.Println("request", requestNumbder)

	ctx.FastResponse("json", map[string]interface{}{
		"app":     appNumber,
		"request": requestNumbder,
	}, nil)
}
