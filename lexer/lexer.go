package lexer

import (
	"fmt"
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

func (l *Lexer) Run() error {

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
			str, err := l.lexString()
			if err != nil {
				return err
			}
			token.Tokens = append(token.Tokens, token.Token{TokenType: token.STRING, Literal: str})
		default:

			if unicode.IsNumber(l.currentChar) || l.currentChar == '-' {
				number, err := l.lexNumber()
				if err != nil {
					return err
				}
				token.Tokens = append(token.Tokens, token.Token{TokenType: token.NUMBER, Literal: number})
			} else if unicode.IsLetter(l.currentChar) {
				keyword := l.lexKeyword()

				switch keyword {
				case "true", "false":
					token.Tokens = append(token.Tokens, token.Token{TokenType: token.BOOLEAN, Literal: keyword})
				case "null":
					token.Tokens = append(token.Tokens, token.Token{TokenType: token.NULL, Literal: keyword})
				default:
					return fmt.Errorf("unexpected keyword: %s", keyword)
				}

			} else {
				return fmt.Errorf("unexpected token: %c", l.currentChar)
			}
		}
		l.skipWhitespace()
	}

	log.Println("Lexical analysis completed...")
	return nil
}

func (l *Lexer) lexString() (string, error) {
	l.next()
	start := l.position

	for l.currentChar != '"' && l.currentChar != 0 {
		l.next()
	}

	if l.currentChar == 0 {
		return "", fmt.Errorf("unterminated string at position %d", l.position)
	}

	str := string(l.input[start:l.position])
	l.next()
	return str, nil
}

func (l *Lexer) lexNumber() (string, error) {
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
			return "", fmt.Errorf("Invalid number format: no digits after decimal point")
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
			return "", fmt.Errorf("Invalid number format: no digits after exponent")
		}
		for unicode.IsDigit(l.currentChar) {
			l.next()
		}
	}

	return string(l.input[start:l.position]), nil
}

func (l *Lexer) lexKeyword() string {
	start := l.position

	for unicode.IsLetter(l.currentChar) {
		l.next()
	}

	return string(l.input[start:l.position])
}
