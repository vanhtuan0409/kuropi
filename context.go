package kuropi

import (
	"net/http"
)

type Context interface {
	Request() *http.Request
	ResponseWriter() http.ResponseWriter

	Responser(name string) Responser
	FastResponse(responserName string, result interface{}, err error)

	Query(name string) string
	Var(name string) string
	ParseForm(dst interface{}) error
	ParseJson(dst interface{}) error
}

type context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	responsers     ResponserMap
}

func (c *context) Request() *http.Request {
	return c.request
}

func (c *context) ResponseWriter() http.ResponseWriter {
	return c.responseWriter
}

func (c *context) Responser(name string) Responser {
	return c.responsers[name]
}

func (c *context) FastResponse(responserName string, result interface{}, err error) {
	rw := c.ResponseWriter()
	responser := c.Responser(responserName)
	if responser == nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	responser.Handle(rw, result, err)
}
