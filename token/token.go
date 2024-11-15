package token

import "fmt"

type Token struct {
	TokenType tokenType
	Literal   string
}

const (

	//EOF
	EOF tokenType = "EOF"

	// structural type
	LEFT_BRACKET        tokenType = "left-bracket"
	RIGHT_BRACKET       tokenType = "right-bracket"
	LEFT_CURLY_BRACKET  tokenType = "left-curly-bracket"
	RIGHT_CURLY_BRACKET tokenType = "right-curly-bracket"
	COLON               tokenType = "colon"
	COMMA               tokenType = "comma"

	// number type
	NUMBER tokenType = "number"

	// string type
	STRING tokenType = "string"

	// boolean type
	BOOLEAN tokenType = "boolean"

	// Null type
	NULL tokenType = "null"
)

type tokenType string

var Tokens = make([]Token, 0)

func PrintTokens() {
	for _, t := range Tokens {
		fmt.Println(t)
	}
}

func (t Token) String() string {
	return fmt.Sprintf("Token{Type: %s, Literal: '%s'}", t.TokenType, t.Literal)
}
