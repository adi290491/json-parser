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

	inputStr := `{
		"name": "Alice",
		"age": 30,
		"isStudent": false,
		"grades": [95.5, 87, 92],
		"address": {"city": "NYC", "zip": "10001"},
		"spouse": null
	}`

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
