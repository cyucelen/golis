package fn

import "github.com/cyucelen/golis/types"

type ReduceFn func([]types.Object) (types.Object, error)

type ReducerFn func(prev, cur types.Object) (types.Object, error)

func Reduce(objects []types.Object, reducer ReducerFn) (types.Object, error) {
	acc := objects[0]
	for i := 1; i < len(objects); i++ {
		var err error
		acc, err = reducer(acc, objects[i])
		if err != nil {
			return nil, err
		}
	}

	return acc, nil
}

func MakeReduceFn(reducer ReducerFn) types.Function {
	return func(args []types.Object) (types.Object, error) {
		return Reduce(args, reducer)
	}
}
