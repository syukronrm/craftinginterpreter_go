package main

import (
	"fmt"
	"os"
	"syukronrm/craftinginterpreter/lox"
)

func main() {
	lox := lox.Lox{}

	if len(os.Args) > 2 {
		fmt.Println("Usage main [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		err := lox.RunFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
		}
	} else {
		lox.RunPrompt()
	}
}
