package parser

import (
	"ast"
	"scanner"
	"testing"
)

func TestReturnStatement(t *testing.T) {
	input := "return 15;"

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

	stmt, ok := program.Statements[0].(*ast.ReturnStatement)
	if !ok {
		t.Fatalf("program.Statements[0]  is not *ast.ReturnStatement. Received %T", program.Statements[0])
	}
	if stmt.TokenLiteral() != "return" {
		t.Fatalf("Expected TokenLiteral of ReturnStatement to be `return`. Received %s", stmt.TokenLiteral())
	}
	if stmt.ReturnValue.TokenLiteral() != "15" {
		t.Fatalf("Expected TokenLiteral of ReturnValue to be `15`. Received %s", stmt.ReturnValue.TokenLiteral())
	}

}
