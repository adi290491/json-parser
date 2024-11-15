package main

import (
	"encoding/json"
	"fmt"
	"json-parser/lexer"
	"json-parser/parser"
	"json-parser/token"
	"log"
)

func main() {
	fmt.Println("-----JSON PARSER-----")

	inputStr := `{
  "id": 123,
  "name": "John Doe",
  "isActive": true,
  "contact": {
    "email": "john.doe@example.com",
    "phone": "+1234567890"
  },
  "address": {
    "street": "123 Elm St",
    "city": "Metropolis",
    "zipcode": "54321",
    "coordinates": {
      "latitude": 40.7128,
      "longitude": -74.0060
    }
  },
  "skills": ["Go", "Python", "JavaScript"],
  "projects": [
    {
      "title": "Project A",
      "status": "completed",
      "technologies": ["Go", "Docker", "Kubernetes"]
    },
    {
      "title": "Project B",
      "status": "in progress",
      "technologies": ["Python", "Flask", "PostgreSQL"]
    }
  ],
  "experienceYears": 5,
  "isAvailableForHire": null
}
`

	lexer := lexer.NewLexer(inputStr)
	lexer.Run()

	lexerTokens := token.Tokens
	token.PrintTokens()
	parser := parser.NewParser(lexerTokens)
	result, err := parser.Parse()

	if err != nil {
		log.Fatalf("Parsing error: %v\n", err)
	}
	log.Println("JSON is Valid!!")

	formattedoutput, err := json.MarshalIndent(result, "", "\t")

	if err != nil {
		log.Fatalf("Error formatting output: %v\n", err)
	}

	fmt.Println("Parsed JSON")
	fmt.Println(string(formattedoutput))

}
