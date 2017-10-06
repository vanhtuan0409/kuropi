package kuropi

import (
	"net/http"
)

type Context interface {
	Request() *http.Request
	SetRequest(r *http.Request)

	ResponseWriter() http.ResponseWriter
	SetResponseWriter(rw http.ResponseWriter)

	Responser(name string) Responser
	SetResponser(name string, rs Responser)
	SetResponsers(rss ResponserMap)
	FastResponse(responserName string, result interface{}, err error)
}

type context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	responsers     ResponserMap
}

func NewContext() Context {
	return &context{
		responsers: ResponserMap{},
	}
}

func (c *context) Request() *http.Request {
	return c.request
}

func (c *context) SetRequest(r *http.Request) {
	c.request = r
}

func (c *context) ResponseWriter() http.ResponseWriter {
	return c.responseWriter
}

func (c *context) SetResponseWriter(rw http.ResponseWriter) {
	c.responseWriter = rw
}

func (c *context) Responser(name string) Responser {
	return c.responsers[name]
}

func (c *context) SetResponser(name string, rs Responser) {
	c.responsers[name] = rs
}

func (c *context) SetResponsers(rss ResponserMap) {
	c.responsers = rss
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
