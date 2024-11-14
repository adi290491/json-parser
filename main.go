package main

import (
	"fmt"
	"json-parser/lexer"
	"json-parser/parser"
	"json-parser/token"
	"log"
)

func main() {
	fmt.Println("-----JSON PARSER-----")

	inputStr := `{{}}`

	lexer := lexer.NewLexer(inputStr)
	lexer.Run()

	lexerTokens := token.Tokens
	token.PrintTokens()
	parser := parser.NewParser(lexerTokens)
	err := parser.Parse()

	if err != nil {
		log.Fatalf("Parsing error: %v\n", err)
	} else {
		log.Println("JSON is Valid!!")
	}

}
