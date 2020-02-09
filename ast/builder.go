package ast

import "token"

type AstBuilder struct {
	Node Node
}

func (b *AstBuilder) Identifier(identity string) *Identifier {

	return &Identifier{
		Token: token.Token{TokenKind: token.IDENT, Literal: identity},
		Value: identity,
	}

}

func (b *AstBuilder) LetStatement(identifier string, value Expression) *LetDeclarationStatement {
	return &LetDeclarationStatement{
		Token: token.Token{TokenKind: token.LET, Literal: "let"},
		Name:  b.Identifier(identifier),
		Value: value,
	}
}

func (b *AstBuilder) InfixExpression(operator string, left, right Expression) *InfixExpression {
	return &InfixExpression{
		Token:    token.Token{},
		Left:     left,
		Operator: operator,
		Right:    right,
	}

}

func (b *AstBuilder) PrefixExpression(operator string, right Expression) *PrefixExpression {
	return &PrefixExpression{
		Token:    token.Token{},
		Operator: operator,
		Right:    right,
	}
}
