package parser

import (
	"fmt"
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

	prefixParsers map[token.TokenKind]prefixParseFn
	infixParsers  map[token.TokenKind]infixParseFn
}

func (p *Parser) Errors() []string {
	return p.errors
}

func New(lxr *scanner.Scanner) *Parser {

	p := &Parser{
		lxr:    lxr,
		errors: []string{},
	}

	p.prefixParsers = make(map[token.TokenKind]prefixParseFn)
	p.infixParsers = make(map[token.TokenKind]infixParseFn)

	// p.succ
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lxr.NextToken()
}

func (p *Parser) isCurrentToken(t token.TokenKind) bool {
	return p.currentToken.TokenKind == t
}

func (p *Parser) isPeekToken(t token.TokenKind) bool {
	return p.peekToken.TokenKind == t
}

func (p *Parser) expectPeek(t token.TokenKind) bool {
	if p.isPeekToken(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekError(t token.TokenKind) {
	msg := fmt.Sprintf("expected next token to be '%s', received %s", t, p.peekToken.TokenKind)
	p.errors = append(p.errors, msg)
}
