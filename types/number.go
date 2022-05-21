package types

type Number struct {
	Value int
}

func NewNumber(value int) Number {
	return Number{Value: value}
}
