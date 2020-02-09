package parser

import (
	"ast"
	"fmt"
	"token"
)

type prefixParseFn func() ast.Expression

func (p *Parser) registerPrefix(TokenKind token.TokenKind, fn prefixParseFn) {
	p.prefixParsers[TokenKind] = fn
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
}

func (p *Parser) noPrefixParser(t token.TokenKind) {
	msg := fmt.Sprintf("no prefix parse function defined for TokenKind: %s", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expr := &ast.PrefixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
	}

	p.nextToken()
	expr.Right = p.parseExpression(PREFIX)
	return expr
}

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
