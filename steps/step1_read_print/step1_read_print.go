package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cyucelen/golis/printer"
	"github.com/cyucelen/golis/reader"
	"github.com/cyucelen/golis/types"
)

func Read(s string) (types.Object, error) {
	return reader.ReadString(s)
}

func Eval(object types.Object) types.Object { return object }

func Print(object types.Object) string {
	return printer.PrintString(object)
}

func Rep(s string) (string, error) {
	r, err := Read(s)
	if err != nil {
		return "", err
	}

	return Print(Eval(r)), nil
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
