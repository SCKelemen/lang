package ast

import (
	"token"
)

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (lit *IntegerLiteral) expressionNode()      {}
func (lit *IntegerLiteral) TokenLiteral() string { return lit.Token.Literal }
func (lit *IntegerLiteral) String() string       { return lit.Token.Literal }
