package ast

import (
	"bytes"
	"token"
)

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

type LetDeclarationStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetDeclarationStatement) statementNode()       {}
func (ls *LetDeclarationStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetDeclarationStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteRune(';')

	return out.String()
}

type ReturnStatement struct {
	Token       token.Token // 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral())
	out.WriteRune(' ')

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteRune(';')

	return out.String()
}

// ExpressionStatement is required for
// side-effecting code such as
// counter++;
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
