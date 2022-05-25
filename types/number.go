package types

type Number struct {
	Value int
}

func NewNumber(value int) Number {
	return Number{Value: value}
}

func IsNumber(o Object) bool {
	_, ok := o.(Number)
	return ok
}
