package parser

import (
	"ast"
	"scanner"
	"testing"
)

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	lxr := scanner.New(input)
	p := New(lxr)

	program := p.ParseProgram()
	errors := p.Errors()
	if len(errors) != 0 {
		t.Errorf("parser had %d errors", len(errors))
		for _, msg := range errors {
			t.Errorf("parser error: %q", msg)
		}
		t.FailNow()
	}

	if len(program.Statements) != 1 {
		t.Fatalf("Program has an unexpected number of statements. Expected 1, received %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0]  is not *ast.ExpressionStatement. Received %T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expression not of type *ast.Identifier, received %T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s, received %s", "foobar", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s, received %s", "foobar", ident.TokenLiteral())
	}
}
