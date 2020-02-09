package parser

import (
	"ast"
	"token"
)

func (p *Parser) parseExpression(precedence Precedence) ast.Expression {
	prefix := p.prefixParsers[p.currentToken.TokenKind]
	if prefix == nil {
		// no matching prefix parser for this tokenkind
		p.noPrefixParser(p.currentToken.TokenKind)
		return nil
	}

	leftExpr := prefix()

	// read until next semi
	for !p.isPeekToken(token.SEMI) && precedence < p.peekPrecedence() {
		infix := p.infixParsers[p.peekToken.TokenKind]
		if infix == nil {
			// end of expr tree
			return leftExpr
		}
		p.nextToken()

		leftExpr = infix(leftExpr)

	}
	return leftExpr
}
