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
		input:       []rune(inputStr),
		position:    0,
		currentChar: rune(inputStr[0]),
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
		l.skipWhitespace()
		l.currentChar = l.input[l.position]

		switch l.currentChar {
		case '[':
			token.Tokens = append(token.Tokens, token.Token{TokenType: token.LEFT_BRACKET, Literal: string(l.currentChar)})
			l.next()
		case ']':
			token.Tokens = append(token.Tokens, token.Token{TokenType: token.RIGHT_BRACKET, Literal: string(l.currentChar)})
			l.next()
		case '{':
			token.Tokens = append(token.Tokens, token.Token{TokenType: token.LEFT_CURLY_BRACKET, Literal: string(l.currentChar)})
			l.next()
		case '}':
			token.Tokens = append(token.Tokens, token.Token{TokenType: token.RIGHT_CURLY_BRACKET, Literal: string(l.currentChar)})
			l.next()
		case ':':
			token.Tokens = append(token.Tokens, token.Token{TokenType: token.COLON, Literal: string(l.currentChar)})
			l.next()
		case ',':
			token.Tokens = append(token.Tokens, token.Token{TokenType: token.COMMA, Literal: string(l.currentChar)})
			l.next()
		case '"':
			str := l.lexString()
			token.Tokens = append(token.Tokens, token.Token{TokenType: token.STRING, Literal: str})
		default:

			if unicode.IsNumber(l.currentChar) || l.currentChar == '-' {
				number := l.lexNumber()
				token.Tokens = append(token.Tokens, token.Token{TokenType: token.NUMBER, Literal: number})
			} else if unicode.IsLetter(l.currentChar) {
				keyword := l.lexKeyword()

				switch keyword {
				case "true", "false":
					token.Tokens = append(token.Tokens, token.Token{TokenType: token.BOOLEAN, Literal: keyword})
				case "null":
					token.Tokens = append(token.Tokens, token.Token{TokenType: token.NULL, Literal: keyword})
				default:
					log.Fatalf("Unexpected keywork: %s", keyword)
				}

			} else {
				log.Fatalf("Unexpected token: %c\n", l.currentChar)
			}
		}
		l.skipWhitespace()
	}

	log.Println("Lexical analysis completed...")
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

func (l *Lexer) lexNumber() string {
	start := l.position

	//check for negative
	if l.currentChar == '-' {
		l.next()
	}

	//check for integer part
	for unicode.IsDigit(l.currentChar) {
		l.next()
	}

	//check floating part
	if l.currentChar == '.' {
		l.next()

		if !unicode.IsDigit(l.currentChar) {
			log.Fatal("Invalid number format: no digits after decimal point")
		}

		for unicode.IsDigit(l.currentChar) {
			l.next()
		}
	}

	// Check for scientific notation (e or E)
	if l.currentChar == 'e' || l.currentChar == 'E' {
		l.next() // Consume the 'e' or 'E'

		// Check for an optional '+' or '-' sign
		if l.currentChar == '+' || l.currentChar == '-' {
			l.next()
		}

		// There must be digits after the exponent
		if !unicode.IsDigit(l.currentChar) {
			log.Fatal("Invalid number format: no digits after exponent")
		}
		for unicode.IsDigit(l.currentChar) {
			l.next()
		}
	}

	return string(l.input[start:l.position])
}

func (l *Lexer) lexKeyword() string {
	start := l.position

	for unicode.IsLetter(l.currentChar) {
		l.next()
	}

	return string(l.input[start:l.position])
}
