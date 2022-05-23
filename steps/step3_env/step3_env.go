package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/cyucelen/golis/env"
	"github.com/cyucelen/golis/fn"
	"github.com/cyucelen/golis/printer"
	"github.com/cyucelen/golis/reader"
	"github.com/cyucelen/golis/types"
)

func Read(s string) (types.Object, error) {
	return reader.ReadString(s)
}

func Eval(ast types.Object, env *env.Env) (types.Object, error) {
	if !types.IsList(ast) {
		return EvalAST(ast, env)
	}

	list := types.MustMakeList(ast)

	if list.IsEmpty() {
		return ast, nil
	}

	op := list.Values()[0].(types.Symbol)
	switch op {
	case types.DefineSymbol:
		return EvalDefine(list, env)
	case types.LetSymbol:
		return EvalLet(list, env)
	}

	evaluatedAst, err := EvalAST(list, env)
	if err != nil {
		return nil, err
	}

	evaluatedList := types.MustMakeList(evaluatedAst)

	fn := evaluatedList.Values()[0].(fn.ReduceFn)
	args := evaluatedList.Values()[1:]

	return fn(args), nil
}

func EvalAST(ast types.Object, env *env.Env) (types.Object, error) {
	switch astVal := ast.(type) {
	case types.Symbol:
		symbolFn, err := env.Get(astVal)
		if err != nil {
			return nil, err
		}

		return symbolFn, nil
	case *types.List:
		evaluatedList, err := fn.Map(astVal.Values(), evalWitEnv(env))
		if err != nil {
			return nil, err
		}

		return types.NewList(evaluatedList...), nil
	default:
		return astVal, nil
	}
}

func EvalDefine(ast *types.List, env *env.Env) (types.Object, error) {
	if ast.Length() != 3 {
		return nil, errors.New("def! must have 2 args")
	}

	k := ast.Values()[1].(types.Symbol)
	v := ast.Values()[2]

	evaluatedV, err := Eval(v, env)
	if err != nil {
		return nil, err
	}

	return env.Set(k, evaluatedV), nil
}

func EvalLet(list *types.List, e *env.Env) (types.Object, error) {
	if list.Length() != 3 {
		return nil, errors.New("let* must have 2 args")
	}

	newEnv := env.NewEnv(e)

	bindings := list.Values()[1].(types.Sequence)
	body := list.Values()[2]

	bindingChunks := fn.Chunk(bindings, 2)
	for _, binding := range bindingChunks {
		symbol := binding[0].(types.Symbol)
		value := binding[1]

		evaluated, err := Eval(value, newEnv)
		if err != nil {
			return nil, err
		}

		newEnv.Set(symbol, evaluated)
	}

	return Eval(body, newEnv)
}

func evalWitEnv(env *env.Env) func(ast types.Object) (types.Object, error) {
	return func(ast types.Object) (types.Object, error) {
		return Eval(ast, env)
	}
}

func Print(s types.Object) string {
	return printer.PrintString(s)
}

func Rep(sexp string, env *env.Env) (string, error) {
	ast, err := Read(sexp)
	if err != nil {
		return "", err
	}

	res, err := Eval(ast, env)
	if err != nil {
		return "", err
	}

	return Print(res), nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	replEnv := env.NewEnv(nil)
	replEnv.Set(types.AdditionSymbol, fn.Sum)
	replEnv.Set(types.SubtractionSymbol, fn.Subtract)
	replEnv.Set(types.MultiplicationSymbol, fn.Multiply)
	replEnv.Set(types.DivisionSymbol, fn.Divide)

	for {
		fmt.Print("user> ")

		if !scanner.Scan() {
			break
		}

		p, err := Rep(scanner.Text(), replEnv)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(p)
		}
	}
}
