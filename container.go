package kuropi

import (
	"errors"
	"sync"
)

type Container interface {
	AddDefinition(def Definition) error
	GetDefinitions() []Definition
	Get(ctx Context, name string) (interface{}, error)
	Destroy()
}

type container struct {
	definitions DefinitionMap
	lock        sync.Mutex
	objects     map[string]interface{}
	built       []string
}

func NewContainer() Container {
	return &container{
		definitions: DefinitionMap{},
		objects:     map[string]interface{}{},
	}
}

func (c *container) AddDefinition(def Definition) error {
	if err := validateDefinition(def); err != nil {
		return err
	}
	if def.Scope == "" {
		def.Scope = AppScope
	}
	c.definitions[def.Name] = def
	return nil
}

func (c *container) GetDefinitions() []Definition {
	defs := []Definition{}
	for _, def := range c.definitions {
		defs = append(defs, def)
	}
	return defs
}

func (c *container) Get(ctx Context, name string) (interface{}, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	def, ok := c.definitions[name]
	if !ok {
		return nil, errors.New("Definition is not defined")
	}
	if def.Scope != ctx.Scope() {
		if ctx.Parent() != nil {
			return ctx.Parent().GetInstance(name)
		}
		return nil, errors.New("Definition is not defined")
	}

	if stringSliceContains(c.built, name) {
		return nil, errors.New("Definitions must not contain cycle")
	}

	obj, ok := c.objects[name]
	if ok {
		return obj, nil
	}

	c.built = append(c.built, name)
	obj, err := def.Build(ctx)
	if err != nil {
		return nil, err
	}

	c.objects[name] = obj
	return obj, nil
}

func (c *container) IsDefined(name string) bool {
	_, ok := c.definitions[name]
	return ok
}

func (c *container) Destroy() {
	c.lock.Lock()
	defer c.lock.Unlock()

	for name, def := range c.definitions {
		if obj, ok := c.objects[name]; ok && def.Close != nil {
			def.Close(obj)
			delete(c.objects, name)
		}
	}
}
