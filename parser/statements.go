package parser

import (
	"ast"
	"token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.TokenKind {
	case token.TYPE:
		return p.parseTypeDeclaration()
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
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
	if !p.expectPeek(token.ASSIGN) {
		// type declaration should be of the form:
		// type ident = expr
		// if `ident` is not followed by an `assign`
		// then shit is fucked, and we need to fail;
		return nil
	}

	if !p.isCurrentToken(token.SEMI) {
		// the expression continues until the next semi
		stmt.Value = p.parseInfixExpression(stmt.Name)
	}

	// consume semi
	p.nextToken()
	return stmt
}

func (p *Parser) parseLetStatement() *ast.LetDeclarationStatement {
	stmt := &ast.LetDeclarationStatement{Token: p.currentToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.isPeekToken(token.SEMI) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currentToken}

	p.nextToken()

	stmt.ReturnValue = p.parseExpression(LOWEST)

	for !p.isCurrentToken(token.SEMI) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.currentToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.isPeekToken(token.SEMI) {
		p.nextToken()
	}
	return stmt
}
