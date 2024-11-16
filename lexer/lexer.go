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
	Tokens      []token.Token
}

func NewLexer(inputStr string) *Lexer {
	l := &Lexer{
		input:       []rune(inputStr),
		position:    0,
		currentChar: rune(inputStr[0]),
		Tokens:      []token.Token{},
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
	for isWhitespace(l.currentChar) {
		l.next()
	}
}

func isWhitespace(c rune) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}

func (l *Lexer) Run() ([]token.Token, error) {

	log.Println("Lexer running...")

	for l.position < len(l.input) {
		l.skipWhitespace()
		l.currentChar = l.input[l.position]

		switch l.currentChar {
		case '[':
			l.Tokens = append(l.Tokens, token.Token{TokenType: token.LEFT_BRACKET, Literal: string(l.currentChar)})
			l.next()
		case ']':
			l.Tokens = append(l.Tokens, token.Token{TokenType: token.RIGHT_BRACKET, Literal: string(l.currentChar)})
			l.next()
		case '{':
			l.Tokens = append(l.Tokens, token.Token{TokenType: token.LEFT_CURLY_BRACKET, Literal: string(l.currentChar)})
			l.next()
		case '}':
			l.Tokens = append(l.Tokens, token.Token{TokenType: token.RIGHT_CURLY_BRACKET, Literal: string(l.currentChar)})
			l.next()
		case ':':
			l.Tokens = append(l.Tokens, token.Token{TokenType: token.COLON, Literal: string(l.currentChar)})
			l.next()
		case ',':
			l.Tokens = append(l.Tokens, token.Token{TokenType: token.COMMA, Literal: string(l.currentChar)})
			l.next()
		case '"':
			str, err := l.lexString()
			if err != nil {
				return nil, err
			}
			l.Tokens = append(l.Tokens, token.Token{TokenType: token.STRING, Literal: str})
		default:

			if unicode.IsNumber(l.currentChar) || l.currentChar == '-' {
				number, err := l.lexNumber()
				if err != nil {
					return nil, err
				}
				l.Tokens = append(l.Tokens, token.Token{TokenType: token.NUMBER, Literal: number})
			} else if unicode.IsLetter(l.currentChar) {
				keyword := l.lexKeyword()

				switch keyword {
				case "true", "false":
					l.Tokens = append(l.Tokens, token.Token{TokenType: token.BOOLEAN, Literal: keyword})
				case "null":
					l.Tokens = append(l.Tokens, token.Token{TokenType: token.NULL, Literal: keyword})
				default:
					return nil, fmt.Errorf("unexpected keyword: %s", keyword)
				}

			} else {
				return nil, fmt.Errorf("unexpected token: %c", l.currentChar)
			}
		}
		l.skipWhitespace()
	}

	log.Println("Lexical analysis completed...")
	return l.Tokens, nil
}

func isHexDigit(c rune) bool {
	return (c >= '0' && c <= '9') || (c >= 'A' && c <= 'F') || (c >= 'a' && c <= 'f')
}

func (l *Lexer) lexString() (string, error) {
	l.next()

	var result []rune
	for l.currentChar != '"' && l.currentChar != 0 {

		if l.currentChar == '\\' {
			l.next()

			switch l.currentChar {
			case '"', '\\', '/', 'b', 'f', 'n', 'r', 't':
				result = append(result, '\\', l.currentChar)

			case 'u':
				unicode := l.position + 1
				for i := 0; i < 4; i++ {
					l.next()
					if !isHexDigit(l.currentChar) {
						return "", fmt.Errorf("invalid unicode escape sequence at position %d", l.position)
					}
				}
				result = append(result, l.input[unicode-2:l.position+1]...)
			default:
				return "", fmt.Errorf("invalid escape character '\\%c' at position %d", l.currentChar, l.position)
			}

		} else {
			result = append(result, l.currentChar)
		}
		l.next()
	}

	if l.currentChar == 0 {
		return "", fmt.Errorf("unterminated string at position %d", l.position)
	}

	l.next()
	return string(result), nil
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
