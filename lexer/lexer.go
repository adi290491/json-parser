package lexer

import (
	"json-parser/token"
	"log"
	"unicode"
)

type Lexer struct {
	input       []rune
	position    int
	currentChar rune
}

func NewLexer(inputStr string) *Lexer {
	l := &Lexer{
		input: []rune(inputStr),
	}

	return l
}

func (l *Lexer) next() {
	l.position++

	if l.position < len(l.input) {
		l.currentChar = l.input[l.position]
	} else {
		l.currentChar = 0
	}
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.currentChar) {
		l.next()
	}
}

func (l *Lexer) Run() {

	log.Println("Lexer running...")

	for l.position < len(l.input) {
		l.currentChar = l.input[l.position]

		switch l.currentChar {
		case '[':
			token.Tokens = append(token.Tokens, token.Token{TokenType: token.LEFT_BRACKET, Literal: string(l.currentChar)})
		case ']':
			token.Tokens = append(token.Tokens, token.Token{TokenType: token.RIGHT_BRACKET, Literal: string(l.currentChar)})
		case '{':
			token.Tokens = append(token.Tokens, token.Token{TokenType: token.LEFT_CURLY_BRACKET, Literal: string(l.currentChar)})
		case '}':
			token.Tokens = append(token.Tokens, token.Token{TokenType: token.RIGHT_CURLY_BRACKET, Literal: string(l.currentChar)})
		case ':':
			token.Tokens = append(token.Tokens, token.Token{TokenType: token.COLON, Literal: string(l.currentChar)})
		case ',':
			token.Tokens = append(token.Tokens, token.Token{TokenType: token.COMMA, Literal: string(l.currentChar)})
		case 0:
			token.Tokens = append(token.Tokens, token.Token{TokenType: token.EOF, Literal: ""})
		case '"':
			str := l.lexString()
			token.Tokens = append(token.Tokens, token.Token{TokenType: token.STRING, Literal: str})
		default:
			log.Fatalf("Unexpected token: %c\n", l.currentChar)
		}
		l.next()
	}

}

func (l *Lexer) lexString() string {
	l.next()
	start := l.position

	for l.currentChar != '"' && l.currentChar != 0 {
		l.next()
	}

	if l.currentChar == 0 {
		log.Fatal("unterminated string")
	}

	str := string(l.input[start:l.position])
	l.next()
	return str
}
