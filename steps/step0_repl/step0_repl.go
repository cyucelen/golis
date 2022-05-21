package main

import (
	"bufio"
	"fmt"
	"os"
)

func Read(s string) string { return s }

func Eval(s string) string { return s }

func Print(s string) string { return s }

func Rep(s string) string {
	return Print(Eval(Read(s)))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("user> ")

		if !scanner.Scan() {
			break
		}

		fmt.Println(Rep(scanner.Text()))
	}
}
