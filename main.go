package main

import (
	"fmt"
	"os"

	"github.com/Rajeevnita1993/calculator/calc"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: calc 'expression'")
		return
	}

	expression := os.Args[1]
	calc.Calculate(expression)
}
