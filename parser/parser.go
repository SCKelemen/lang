package parser

import (
	"ast"
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

func (p *Parser) Errors() []string {
	return p.errors
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

// if the language is completely expression based, we won't need this, but that's currently not in the cards

type Precendence int

const (
	_ Precendence = iota
	LOWEST
	EQUALITY
	COMPARE
	SUMMATION
	PRODUCT
	PREFIX
	INVOCATION
)

func (p *Parser) parseTypeDeclaration() *ast.TypeDeclarationStatement {
	stmt := &ast.TypeDeclarationStatement{Token: p.currentToken}

	if !p.expectPeek(token.IDENT) {
		// type declaration should be of the form:
		// type ident = expr
		// if `type` is not followed by an `ident`
		// then shit is fucked, and we need to fail;
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
	if !p.expectPeek(token.EQL) {
		// type declaration should be of the form:
		// type ident = expr
		// if `ident` is not followed by an `eql`
		// then shit is fucked, and we need to fail;
		return nil
	}

	for !p.isCurrentToken(token.SEMI) {
		// the expression continues until the next semi
		p.nextToken()
	}

	return stmt
}
