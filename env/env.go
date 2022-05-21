package env

import (
	"fmt"

	"github.com/cyucelen/golis/types"
)

type Env struct {
	outer *Env
	data  map[types.Symbol]types.Object
}

func NewEnv(outer *Env) *Env {
	return &Env{
		outer: outer,
		data:  make(map[types.Symbol]types.Object),
	}
}

func (e *Env) Set(k types.Symbol, v types.Object) types.Object {
	e.data[k] = v
	return v
}

func (e *Env) Find(k types.Symbol) (*Env, error) {
	_, ok := e.data[k]
	if ok {
		return e, nil
	}

	if e.outer == nil {
		return nil, fmt.Errorf("%s not found", k.Name)
	}

	return e.outer.Find(k)
}

func (e *Env) Get(k types.Symbol) (types.Object, error) {
	foundEnv, err := e.Find(k)
	if err != nil {
		return nil, err
	}

	return foundEnv.data[k], nil
}
