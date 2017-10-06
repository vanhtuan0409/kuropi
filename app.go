package kuropi

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	EmptyMdwChain = []Middleware{}
)

type App interface {
	Use(mdws ...Middleware)
	Context() Context
	Get(path string, mdws []Middleware, handler HandlerFunc)
	Post(path string, mdws []Middleware, handler HandlerFunc)
	Put(path string, mdws []Middleware, handler HandlerFunc)
	Delete(path string, mdws []Middleware, handler HandlerFunc)
	Serve(port int) error
	Responser(name string, rs Responser)
	SetInstance(name string, obj interface{}) error
	AddDefinition(def Definition) error
}

type app struct {
	port        int
	appContext  Context
	globalMdws  []Middleware
	router      *mux.Router
	definitions []Definition
}

func NewApp() App {
	return &app{
		globalMdws: []Middleware{},
		appContext: NewContext(AppScope, nil),
		router:     mux.NewRouter(),
	}
}

func (a *app) Use(mdws ...Middleware) {
	a.globalMdws = append(a.globalMdws, mdws...)
}

func (a *app) Context() Context {
	return a.appContext
}

func (a *app) Get(path string, mdws []Middleware, f HandlerFunc) {
	a.addRoute(Route{
		Method:      GET,
		Middlewares: mdws,
		Path:        path,
		HandlerFunc: f,
	})
}

func (a *app) Post(path string, mdws []Middleware, f HandlerFunc) {
	a.addRoute(Route{
		Method:      POST,
		Middlewares: mdws,
		Path:        path,
		HandlerFunc: f,
	})
}

func (a *app) Put(path string, mdws []Middleware, f HandlerFunc) {
	a.addRoute(Route{
		Method:      PUT,
		Middlewares: mdws,
		Path:        path,
		HandlerFunc: f,
	})
}

func (a *app) Delete(path string, mdws []Middleware, f HandlerFunc) {
	a.addRoute(Route{
		Method:      DELETE,
		Middlewares: mdws,
		Path:        path,
		HandlerFunc: f,
	})
}

func (a *app) Serve(port int) error {
	if err := applyDefinitionToContext(a.appContext, a.definitions); err != nil {
		return err
	}

	a.port = port
	addr := fmt.Sprintf(":%d", port)
	return http.ListenAndServe(addr, a.router)
}

func (a *app) Responser(name string, rs Responser) {
	a.appContext.SetResponser(name, rs)
}

func (a *app) SetInstance(name string, obj interface{}) error {
	return a.appContext.SetInstance(name, obj)
}

func (a *app) AddDefinition(def Definition) error {
	if err := validateDefinition(def); err != nil {
		return err
	}
	a.definitions = append(a.definitions, def)
	return nil
}

func (a *app) addRoute(route Route) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		requestContext := a.appContext.SubContext(RequestScope)
		requestContext.SetRequest(r)
		requestContext.SetResponseWriter(w)
		if err := applyDefinitionToContext(requestContext, a.definitions); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer func() {
			requestContext.Destroy()
			if err := recover(); err != nil {
				a.Serve(a.port)
			}
		}()

		appliedMdws := a.getAppliedMiddleware(route)
		wrappedHandler := a.getWrappedHandler(appliedMdws, route.HandlerFunc)
		wrappedHandler(requestContext)

		requestContext.Destroy()
	}
	a.router.HandleFunc(route.Path, handler).Methods(string(route.Method))
}

func (a *app) getWrappedHandler(mdws []Middleware, handler HandlerFunc) HandlerFunc {
	h := handler
	for i := len(mdws) - 1; i >= 0; i-- {
		h = mdws[i](h)
	}
	return h
}

func (a *app) getAppliedMiddleware(route Route) []Middleware {
	return append(a.globalMdws, route.Middlewares...)
}
