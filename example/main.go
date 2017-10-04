package main

import (
	"github.com/vanhtuan0409/kuropi"
)

func main() {
	app := kuropi.NewApp()
	responser := &JsonResponser{}
	app.Responser("json", responser)
	app.Use(Logger1)
	app.Get("/success", []kuropi.Middleware{Logger2, Logger3}, SuccessHandler)
	app.Get("/error", kuropi.EmptyMdwChain, ErrorHandler)
	app.Serve(3000)
}
