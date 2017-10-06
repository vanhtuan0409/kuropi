package kuropi

import (
	"net/http"
)

func (a *app) Use(mdws ...Middleware) {
	a.globalMdws = append(a.globalMdws, mdws...)
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

func (a *app) addRoute(route Route) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				panic(err)
			}
		}()

		ctx := &context{
			request:        r,
			responseWriter: w,
			responsers:     a.responsers.Clone(),
		}
		appliedMdws := a.getAppliedMiddleware(route)
		wrappedHandler := a.getWrappedHandler(appliedMdws, route.HandlerFunc)
		wrappedHandler(ctx)
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
