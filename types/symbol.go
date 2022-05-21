package types

type Symbol struct {
	Name string
}

func NewSymbol(name string) Symbol {
	return Symbol{Name: name}
}

var AdditionSymbol = Symbol{Name: "+"}
var SubtractionSymbol = Symbol{Name: "-"}
var MultiplicationSymbol = Symbol{Name: "*"}
var DivisionSymbol = Symbol{Name: "/"}

var DefineSymbol = Symbol{Name: "def!"}
var LetSymbol = Symbol{Name: "let*"}
