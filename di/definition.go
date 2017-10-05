package di

import (
	"github.com/vanhtuan0409/kuropi"
)

type Definition struct {
	Name  string
	Scope string
	Build func(ctx kuropi.Context) (interface{}, error)
	Close func(obj interface{})
}

type DefinitionMap map[string]Definition

func (dm DefinitionMap) Clone() DefinitionMap {
	defs := map[string]Definition{}
	for name, def := range dm {
		defs[name] = def
	}
	return defs
}
