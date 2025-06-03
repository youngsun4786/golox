package main

import (
	"fmt"
	"os"
	"github.com/codecrafters-io/interpreter-starter-go/lexer"
	"github.com/codecrafters-io/interpreter-starter-go/token"
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

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	fileContents := string(rawFileContents)
	l := lexer.New(fileContents, filename)
	hasLexicalErrors := false
	for { 
		tok, err := l.NextToken()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err.Error())
			hasLexicalErrors = true
			continue
		}
		fmt.Println(tok.String())
		if tok.Type == token.EOF {
			break
		}
	}
	if hasLexicalErrors {
		os.Exit(65)
	}
	os.Exit(0)
}
