package main

import (
	"math/rand"

	"github.com/vanhtuan0409/kuropi"
)

func main() {
	app := kuropi.NewApp()
	responser := &JsonResponser{}
	app.Responser("json", responser)

	app.AddDefinition(kuropi.Definition{
		Name: "appRandomNumber",
		Build: func(ctx kuropi.Context) (interface{}, error) {
			return rand.Int(), nil
		},
	})
	app.AddDefinition(kuropi.Definition{
		Name:  "requestRandomNumber",
		Scope: kuropi.RequestScope,
		Build: func(ctx kuropi.Context) (interface{}, error) {
			return rand.Int(), nil
		},
	})

	app.Use(Logger1)
	app.Get("/success", []kuropi.Middleware{Logger2, Logger3}, SuccessHandler)
	app.Get("/error", kuropi.EmptyMdwChain, ErrorHandler)
	app.Get("/random", kuropi.EmptyMdwChain, RandomHandler)

	app.Serve(3000)
}
