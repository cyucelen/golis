package fn

import (
	"github.com/cyucelen/golis/types"
)

func Map(objects []types.Object, fn func(types.Object) (types.Object, error)) ([]types.Object, error) {
	res := []types.Object{}
	for _, object := range objects {
		r, err := fn(object)
		if err != nil {
			return nil, err
		}

		res = append(res, r)
	}
	return res, nil
}

func MapObjectToString(objects []types.Object, fn func(s types.Object) string) []string {
	ss := []string{}
	for _, object := range objects {
		ss = append(ss, fn(object))
	}
	return ss
}

func MapString(ss []string, f func(s string) string) []string {
	res := []string{}
	for _, s := range ss {
		res = append(res, f(s))
	}
	return res
}
