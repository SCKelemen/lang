package ast

import (
	"bytes"
	"token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Expression interface {
	Node
	expressionNode()
}

type Statement interface {
	Node
	statementNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())

	}

	return out.String()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type TypeDeclarationStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (tds *TypeDeclarationStatement) statementNode()       {}
func (tds *TypeDeclarationStatement) TokenLiteral() string { return tds.Token.Literal }
func (tds *TypeDeclarationStatement) String() string {
	var out bytes.Buffer

	out.WriteString(tds.TokenLiteral())
	out.WriteRune(' ')
	out.WriteString(tds.Name.String())
	out.WriteString(" = ")

	if tds.Value != nil {
		out.WriteString(tds.Value.String())
	}
	out.WriteRune(';')

	return out.String()
}

type Identifier struct {
	Token token.Token // 'ident' token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

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
