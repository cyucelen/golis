package types

import (
	"errors"
)

type HashMapKey interface {
	String() string
}

type HashMap struct {
	values map[HashMapKey]Object
}

func NewHashMap(kvList ...Object) (*HashMap, error) {
	if len(kvList)%2 == 1 {
		return nil, errors.New("number of arguments must be even")
	}

	values := make(map[HashMapKey]Object)
	for i := 0; i < len(kvList)-1; i += 2 {
		key, ok := kvList[i].(HashMapKey)
		if !ok {
			return nil, errors.New("hash map key should be string or keyword")
		}

		values[key] = kvList[i+1]
	}

	return &HashMap{values: values}, nil
}

func (h *HashMap) Add(key HashMapKey, value Object) {
	h.values[key] = value
}

func (h *HashMap) Values() map[HashMapKey]Object {
	return h.values
}

func (h *HashMap) KVList() []Object {
	objects := []Object{}
	for k, v := range h.values {
		objects = append(objects, k, v)
	}
	return objects
}
