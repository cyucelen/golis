package fn

import "github.com/cyucelen/golis/types"

type ReduceFn func([]types.Object) types.Object

type ReducerFn func(prev, cur types.Object) types.Object

func Reduce(objects []types.Object, reducer ReducerFn) types.Object {
	acc := objects[0]
	for i := 1; i < len(objects); i++ {
		acc = reducer(acc, objects[i])
	}

	return acc
}

func MakeReduceFn(reducer ReducerFn) ReduceFn {
	return func(args []types.Object) types.Object {
		return Reduce(args, reducer)
	}
}
