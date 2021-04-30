package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage main [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		err := runFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
		}
	} else {
		runPrompt()
	}
}

func runFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return errors.New("File not found: " + filename)
	}
	run(string(data))
	return nil
}

func runPrompt() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(" > ")
		text, _ := reader.ReadString('\n')
		if text == "exit" {
			return nil
		}
		run(text)
	}
}

func run(source string) {
	fmt.Println("run", source)
}
