package main

import (
	"fmt"

	"github.com/vanhtuan0409/kuropi"
)

func Logger1(next kuropi.HandlerFunc) kuropi.HandlerFunc {
	return func(ctx kuropi.Context) {
		fmt.Println("logger 1")
		next(ctx)
	}
}

func Logger2(next kuropi.HandlerFunc) kuropi.HandlerFunc {
	return func(ctx kuropi.Context) {
		fmt.Println("logger 2")
		next(ctx)
	}
}

func Logger3(next kuropi.HandlerFunc) kuropi.HandlerFunc {
	return func(ctx kuropi.Context) {
		fmt.Println("logger 3")
		next(ctx)
	}
}
