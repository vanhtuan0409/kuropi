package kuropi

import (
	"net/http"
)

const (
	AppScope        = "AppScope"
	RequestScope    = "RequestScope"
	SubRequestScope = "SubRequestScope"
)

type Context interface {
	Scope() string

	SubContext(scope string) Context
	Parent() Context

	Request() *http.Request
	SetRequest(r *http.Request)

	ResponseWriter() http.ResponseWriter
	SetResponseWriter(rw http.ResponseWriter)

	Responser(name string) Responser
	SetResponser(name string, rs Responser)
	SetResponsers(rss map[string]Responser)
	FastResponse(responserName string, result interface{}, err error)

	AddDefinition(def Definition) error
	GetDefinitions() []Definition
	GetInstance(name string) (interface{}, error)
	FastGetInstance(name string) interface{}
	Destroy()
}

type context struct {
	scope          string
	parent         Context
	children       Context
	request        *http.Request
	responseWriter http.ResponseWriter
	responsers     ResponserMap
	container      Container
}

func NewContext(scope string, parent Context) Context {
	return &context{
		scope:      scope,
		parent:     parent,
		responsers: ResponserMap{},
		container:  NewContainer(),
	}
}

func (c *context) Scope() string {
	return c.scope
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

func (c *context) SetResponsers(rss map[string]Responser) {
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

func (c *context) SubContext(scope string) Context {
	subContext := NewContext(scope, c).(*context)
	subContext.request = c.request
	subContext.responseWriter = c.responseWriter
	subContext.responsers = c.responsers.Clone()
	subContext.container = NewContainer()
	defs := c.GetDefinitions()
	applyDefinitionToContext(subContext, defs)

	c.children = subContext
	return subContext
}

func (c *context) Parent() Context {
	return c.parent
}

func (c *context) AddDefinition(def Definition) error {
	return c.container.AddDefinition(def)
}

func (c *context) GetDefinitions() []Definition {
	return c.container.GetDefinitions()
}

func (c *context) GetInstance(name string) (interface{}, error) {
	return c.container.Get(c, name)
}

func (c *context) FastGetInstance(name string) interface{} {
	obj, _ := c.GetInstance(name)
	return obj
}

func (c *context) Destroy() {
	if c.children != nil {
		c.children.Destroy()
	}
	c.container.Destroy()
}
