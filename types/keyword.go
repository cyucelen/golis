package types

type Keyword struct {
	Value string
}

func NewKeyword(value string) Keyword {
	return Keyword{Value: value}
}

func (k Keyword) String() string {
	return k.Value
}
