package main

import (
	"fmt"
	"os"
	"github.com/youngsun4786/golox/lexer"
)


const (
	LEFT_PAREN rune = '(';
	RIGHT_PAREN rune = ')';
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	// Uncomment this block to pass the first stage
	
	filename := os.Args[2]
	rawFileContents, err := os.ReadFile(filename)

	for _, word := range rawFileContents {
		fmt.Printf("type of content %T\n", word)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	fileContents := string(rawFileContents)

	for _, current := range fileContents {
		switch current {
		case LEFT_PAREN:
			fmt.Println("LEFT_PAREN ( null")
		case RIGHT_PAREN:
			fmt.Println("RIGHT_PAREN ) null")
		}
	}
	fmt.Println("EOF  null")
}
