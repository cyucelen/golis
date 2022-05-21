package types

type String struct {
	Value string
}

func NewString(value string) String {
	return String{Value: value}
}

func (s String) String() string {
	return s.Value
}
