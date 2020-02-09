package ast

import (
	"token"
)

type ExpressionBuilder interface {
	InfixExpression() *InfixExpressionBuilder
	PrefixExpression() *PrefixExpressionBuilder
	Build() *Expression
}

type InfixExpressionBuilder struct {
	token    token.Token
	left     Expression
	operator string
	right    Expression
}

func (b *InfixExpressionBuilder) Operator(op string) *InfixExpressionBuilder {
	b.operator = op
	return b
}

func (b *InfixExpressionBuilder) Left(expr Expression) *InfixExpressionBuilder {
	b.left = expr
	return b
}
func (b *InfixExpressionBuilder) Right(expr Expression) *InfixExpressionBuilder {
	b.right = expr
	return b
}
func (b *InfixExpressionBuilder) Token(token token.Token) *InfixExpressionBuilder {
	b.token = token
	return b
}

func (b *InfixExpressionBuilder) Build() *InfixExpression {
	return &InfixExpression{Token: b.token, Left: b.left, Right: b.right, Operator: b.operator}
}

type PrefixExpressionBuilder struct {
	token    token.Token
	operator string
	right    Expression
}

func (b *PrefixExpressionBuilder) Operator(op string) *PrefixExpressionBuilder {
	b.operator = op
	return b
}

func (b *PrefixExpressionBuilder) Right(expr Expression) *PrefixExpressionBuilder {
	b.right = expr
	return b
}
func (b *PrefixExpressionBuilder) Token(token token.Token) *PrefixExpressionBuilder {
	b.token = token
	return b
}

func (b *PrefixExpressionBuilder) Build() *PrefixExpression {
	return &PrefixExpression{Token: b.token, Right: b.right, Operator: b.operator}
}

func NewExpressionBuilder() ExpressionBuilder {
	return ExpressionBuilderImpl{}
}

type ExpressionBuilderImpl struct{}

func (b ExpressionBuilderImpl) InfixExpression() *InfixExpressionBuilder {
	return &InfixExpressionBuilder{}
}

func (b ExpressionBuilderImpl) PrefixExpression() *PrefixExpressionBuilder {
	return &PrefixExpressionBuilder{}
}

func (b ExpressionBuilderImpl) Build() *Expression {
	return nil
}
