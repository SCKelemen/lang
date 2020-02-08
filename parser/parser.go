package parser

import (
	"scanner"
	"token"
)

// SanityCheck checks your sanity
func SanityCheck(message string) string {
	return "checked " + message
}

type Parser struct {
	lxr          *scanner.Scanner
	currentToken token.Token
	peekToken    token.Token

	errors []string
}

func New(lxr *scanner.Scanner) *Parser {

	p := &Parser{
		lxr:    lxr,
		errors: []string{},
	}

	// p.succ
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lxr.NextToken()
}
