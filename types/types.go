package types

type Object interface{}

type Sequence interface {
	Values() []Object
	Get(idx int) Object
	Length() int
}
