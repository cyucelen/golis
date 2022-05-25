package fn

import (
	"errors"

	"github.com/cyucelen/golis/types"
)

var Sum = MakeReduceFn(sumReducer)
var Subtract = MakeReduceFn(subtractReducer)
var Multiply = MakeReduceFn(multiplyReducer)
var Divide = MakeReduceFn(divideReducer)

func sumReducer(prev, cur types.Object) (types.Object, error) {
	if !types.IsNumber(prev) || !types.IsNumber(cur) {
		return nil, errors.New("args must be number")
	}
	return types.NewNumber(prev.(types.Number).Value + cur.(types.Number).Value), nil
}

func subtractReducer(prev, cur types.Object) (types.Object, error) {
	if !types.IsNumber(prev) || !types.IsNumber(cur) {
		return nil, errors.New("args must be number")
	}
	return types.NewNumber(prev.(types.Number).Value - cur.(types.Number).Value), nil
}

func multiplyReducer(prev, cur types.Object) (types.Object, error) {
	if !types.IsNumber(prev) || !types.IsNumber(cur) {
		return nil, errors.New("args must be number")
	}
	return types.NewNumber(prev.(types.Number).Value * cur.(types.Number).Value), nil
}

func divideReducer(prev, cur types.Object) (types.Object, error) {
	if !types.IsNumber(prev) || !types.IsNumber(cur) {
		return nil, errors.New("args must be number")
	}
	return types.NewNumber(prev.(types.Number).Value / cur.(types.Number).Value), nil
}
