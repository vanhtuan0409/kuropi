package main

import (
	"net/http"

	"github.com/vanhtuan0409/kuropi"
)

func EchoHandler(ctx kuropi.Context) {
	rw := ctx.ResponseWriter()
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Hello world!!!"))
}

func main() {
	app := kuropi.NewApp()
	app.Get("/echo", kuropi.EmptyMdwChain, EchoHandler)
	app.Serve(3000)
}
