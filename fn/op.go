package fn

import "github.com/cyucelen/golis/types"

var Sum = MakeReduceFn(sumReducer)
var Subtract = MakeReduceFn(subtractReducer)
var Multiply = MakeReduceFn(multiplyReducer)
var Divide = MakeReduceFn(divideReducer)

func sumReducer(prev, cur types.Object) types.Object {
	return types.NewNumber(prev.(types.Number).Value + cur.(types.Number).Value)
}

func subtractReducer(prev, cur types.Object) types.Object {
	return types.NewNumber(prev.(types.Number).Value - cur.(types.Number).Value)
}

func multiplyReducer(prev, cur types.Object) types.Object {
	return types.NewNumber(prev.(types.Number).Value * cur.(types.Number).Value)
}

func divideReducer(prev, cur types.Object) types.Object {
	return types.NewNumber(prev.(types.Number).Value / cur.(types.Number).Value)
}
