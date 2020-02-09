package parser

import (
	"ast"
	"token"
)

type infixParseFn func(ast.Expression) ast.Expression

func (p *Parser) registerInfix(TokenKind token.TokenKind, fn infixParseFn) {
	p.infixParsers[TokenKind] = fn
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expr := &ast.InfixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
		Left:     left,
	}
	precedence := p.currentPrecedence()
	p.nextToken()
	expr.Right = p.parseExpression(precedence)

	return expr
}
