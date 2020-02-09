package ast

import (
	"bytes"
	"token"
)

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (expr InfixExpression) expressionNode()      {}
func (expr InfixExpression) TokenLiteral() string { return expr.Token.Literal }
func (expr InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteRune('(')
	out.WriteString(expr.Left.String())
	out.WriteRune(' ')
	out.WriteString(expr.Operator)
	out.WriteRune(' ')
	out.WriteString(expr.Right.String())
	out.WriteRune(')')

	return out.String()
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (expr PrefixExpression) expressionNode()      {}
func (expr PrefixExpression) TokenLiteral() string { return expr.Token.Literal }
func (expr PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteRune('(')
	out.WriteString(expr.Operator)
	out.WriteString(expr.Right.String())
	out.WriteRune(')')

	return out.String()
}
