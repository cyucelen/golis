package types

type Symbol struct {
	Name string
}

func NewSymbol(name string) Symbol {
	return Symbol{Name: name}
}

func MakeSymbol(o Object) Symbol {
	return o.(Symbol)
}

func MakeSymbols(os []Object) []Symbol {
	symbols := []Symbol{}
	for _, o := range os {
		symbols = append(symbols, MakeSymbol(o))
	}
	return symbols
}

var AdditionSymbol = Symbol{Name: "+"}
var SubtractionSymbol = Symbol{Name: "-"}
var MultiplicationSymbol = Symbol{Name: "*"}
var DivisionSymbol = Symbol{Name: "/"}

var DefineSymbol = Symbol{Name: "def!"}
var LetSymbol = Symbol{Name: "let*"}
var FunctionSymbol = Symbol{Name: "fn*"}
