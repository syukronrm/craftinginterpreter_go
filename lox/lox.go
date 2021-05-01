package lox

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"syukronrm/craftinginterpreter/scanner"
)

type Lox struct {
	HadError bool
}

func (lox *Lox) RunFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return errors.New("File not found: " + filename)
	}

	if lox.HadError {
		os.Exit(65)
	}

	lox.run(string(data))
	return nil
}

func (lox *Lox) RunPrompt() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(" > ")
		text, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		if text == "exit" {
			break
		}

		lox.run(text)

		lox.HadError = false
	}

	return nil
}

func (lox *Lox) run(source string) {
	fmt.Println("run", source)
	scanner := new(scanner.Scanner)
	scanner.Init(source)
	scanner.ScanTokens()
}

func (lox *Lox) error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	str := fmt.Sprintf("[line %d] Error %s: %s", line, where, message)
	fmt.Println(str)
}
