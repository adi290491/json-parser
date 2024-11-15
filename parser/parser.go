package parser

import (
	"fmt"
	"json-parser/token"
	"log"
)

type Parser struct {
	tokens    []token.Token
	position  int
	currToken token.Token
}

func NewParser(lexerTokens []token.Token) *Parser {
	p := &Parser{
		tokens: lexerTokens,
	}

	return p
}

func (p *Parser) next() {
	p.position++

	if p.position < len(p.tokens) {
		p.currToken = p.tokens[p.position]
	} else {
		p.currToken = token.Token{TokenType: token.EOF}
	}
}

func (p *Parser) Parse() error {
	log.Println("Parser running...")
	if len(p.tokens) == 0 {
		return fmt.Errorf("no tokens to parse")
	}

	p.currToken = p.tokens[p.position]

	_, err := p.parseValue()

	log.Println("Finished parsing...")
	return err
}

func (p *Parser) parseValue() (interface{}, error) {

	switch p.currToken.TokenType {
	case token.LEFT_CURLY_BRACKET:
		return p.parseObject()
	case token.LEFT_BRACKET:
		return p.parseArray()
	case token.STRING:
		return p.parseString()
	case token.NUMBER:
		return p.parseNumber()
	case token.BOOLEAN:
		return p.parseBoolean()
	case token.NULL:
		return p.parseNull()
	default:
		return nil, fmt.Errorf("unexpected token: %v", p.currToken)
	}
}

func (p *Parser) parseObject() (interface{}, error) {

	jsonObject := make(map[string]interface{})
	p.next()

	for p.currToken.TokenType != token.RIGHT_CURLY_BRACKET {

		//1 Parse key must be string
		if p.currToken.TokenType != token.STRING {
			return nil, fmt.Errorf("expected string, got: %v", p.currToken.Literal)
		}

		key := p.currToken.Literal
		p.next()
		//2. Parse colon
		if p.currToken.TokenType != token.COLON {
			return nil, fmt.Errorf("expected colon, got %v", p.currToken.Literal)
		}
		p.next()

		value, err := p.parseValue()

		if err != nil {
			return nil, err
		}
		jsonObject[key] = value
		p.next()
		//4. Parse comma
		if p.currToken.TokenType == token.COMMA {
			p.next()
		} else if p.currToken.TokenType != token.RIGHT_CURLY_BRACKET {
			return nil, fmt.Errorf("expected '}' or ',', got: %v", p.currToken)
		}
	}

	return jsonObject, nil
}

func (p *Parser) parseArray() (interface{}, error) {

	jsonArray := []interface{}{}

	p.next()

	for p.currToken.TokenType != token.RIGHT_BRACKET {

		value, err := p.parseValue()

		if err != nil {
			return nil, err
		}

		jsonArray = append(jsonArray, value)

		p.next()

		if p.currToken.TokenType == token.COMMA {
			p.next()
		} else if p.currToken.TokenType != token.RIGHT_BRACKET {
			return nil, fmt.Errorf("expected ']' or ',', got: %v", p.currToken)
		}
	}

	return jsonArray, nil
}

func (p *Parser) parseString() (interface{}, error) {
	value := p.currToken.Literal

	return value, nil
}

func (p *Parser) parseNumber() (interface{}, error) {
	var number float64

	_, err := fmt.Sscanf(p.currToken.Literal, "%f", &number)
	if err != nil {
		return nil, fmt.Errorf("invalid number format: %v", p.currToken.Literal)
	}

	return number, nil
}

func (p *Parser) parseBoolean() (interface{}, error) {
	value := p.currToken.Literal == "true"
	return value, nil
}

func (p *Parser) parseNull() (interface{}, error) {
	return nil, nil
}
