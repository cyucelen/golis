package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/cyucelen/golis/fn"
	"github.com/cyucelen/golis/printer"
	"github.com/cyucelen/golis/reader"
	"github.com/cyucelen/golis/types"
)

type ReplEnv map[string]types.Function

var replEnv = ReplEnv{
	"+": fn.Sum,
	"-": fn.Subtract,
	"*": fn.Multiply,
	"/": fn.Divide,
}

func Read(s string) (types.Object, error) {
	return reader.ReadString(s)
}

func Eval(ast types.Object, env ReplEnv) (types.Object, error) {
	switch astVal := ast.(type) {
	case *types.List:
		if astVal.IsEmpty() {
			return ast, nil
		}

		evaluatedList, err := EvalAST(astVal, env)
		if err != nil {
			return nil, err
		}

		list, ok := types.MakeList(evaluatedList)
		if !ok {
			return nil, errors.New("unexpected")
		}

		fn := list.Values()[0].(types.Function)
		args := list.Values()[1:]
		return fn(args)
	default:
		return EvalAST(ast, env)
	}
}

func EvalAST(ast types.Object, env ReplEnv) (types.Object, error) {
	switch astVal := ast.(type) {
	case types.Symbol:
		opFn, ok := env[astVal.Name]
		if !ok {
			return nil, errors.New("not defined")
		}

		return opFn, nil
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

func evalWitEnv(env ReplEnv) func(ast types.Object) (types.Object, error) {
	return func(ast types.Object) (types.Object, error) {
		return Eval(ast, env)
	}
}

func Print(s types.Object) string {
	return printer.PrintString(s)
}

func Rep(sexp string) (string, error) {
	ast, err := Read(sexp)
	if err != nil {
		return "", err
	}

	res, err := Eval(ast, replEnv)
	if err != nil {
		return "", err
	}

	return Print(res), nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("user> ")

		if !scanner.Scan() {
			break
		}

		p, err := Rep(scanner.Text())
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(p)
		}
	}
}
