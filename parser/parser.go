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

	return err
}

func (p *Parser) parseValue() (interface{}, error) {

	switch p.currToken.TokenType {
	case token.LEFT_CURLY_BRACKET:
		return p.parseObject()
	case token.LEFT_BRACKET:
		return p.parseArray()
	default:
		return nil, fmt.Errorf("unexpected token: %v", p.currToken)
	}
}

func (p *Parser) parseObject() (interface{}, error) {

	jsonObject := make(map[string]interface{})
	p.next()

	for p.currToken.TokenType != token.RIGHT_CURLY_BRACKET {

		if p.currToken.TokenType == token.EOF {
			return nil, fmt.Errorf("unexpected EOF while parsing object")
		}

		value, err := p.parseValue()

		if err != nil {
			return nil, err
		}

		jsonObject["value"] = value

		p.next()
	}

	return jsonObject, nil
}

func (p *Parser) parseArray() (interface{}, error) {

	JsonArray := []interface{}{}

	p.next()

	for p.currToken.TokenType != token.RIGHT_BRACKET {

		if p.currToken.TokenType == token.EOF {
			return nil, fmt.Errorf("unexpected EOF while parsing array")
		}

		value, err := p.parseValue()

		if err != nil {
			return nil, err
		}

		JsonArray = append(JsonArray, value)

		p.next()
	}

	return JsonArray, nil
}
