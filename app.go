package kuropi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	EmptyMdwChain = []Middleware{}
)

type App interface {
	Use(mdws ...Middleware)
	Get(path string, mdws []Middleware, handler HandlerFunc)
	Post(path string, mdws []Middleware, handler HandlerFunc)
	Put(path string, mdws []Middleware, handler HandlerFunc)
	Delete(path string, mdws []Middleware, handler HandlerFunc)
	Serve(port int) error
	Server(port int) *http.Server
	Responser(name string, rs Responser)
}

type app struct {
	port       int
	globalMdws []Middleware
	router     *mux.Router
	responsers ResponserMap
}

func NewApp() App {
	return &app{
		globalMdws: []Middleware{},
		responsers: ResponserMap{},
		router:     mux.NewRouter(),
	}
}

func (a *app) Serve(port int) error {
	server := a.Server(port)
	return server.ListenAndServe()
}

func (a *app) Server(port int) *http.Server {
	a.port = port
	addr := fmt.Sprintf(":%d", port)
	return &http.Server{
		Addr:         addr,
		Handler:      a.router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}