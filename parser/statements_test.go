package parser

import (
	"ast"
	"scanner"
	"testing"
)

func TestTypeDeclarationStatement(t *testing.T) {
	input := "type bool = fuck;"

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

	stmt, ok := program.Statements[0].(*ast.TypeDeclarationStatement)
	if !ok {
		t.Fatalf("program.Statements[0]  is not *ast.TypeDeclarationStatement. Received %T", program.Statements[0])
	}
	if stmt.TokenLiteral() != "type" {
		t.Fatalf("Expected TokenLiteral of TypeDeclarationStatement to be `let`. Received %s", stmt.TokenLiteral())
	}
	if stmt.Name.TokenLiteral() != "bool" {
		t.Fatalf("Expected TokenLiteral of Identifier to be `bool`. Received %s", stmt.Name.TokenLiteral())
	}
	if stmt.Value == nil {
		t.Fatal("Expected TokenLiteral of Value to be Expression. Received nil.")
	}
	expr, ok := stmt.Value.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("stmt.Value  is not *ast.InfixExpression. Received %T", stmt.Value)
	}
	if stmt.Value.TokenLiteral() != "=" {
		t.Fatalf("Expected TokenLiteral of Value to be `=`. Received %s", stmt.Value.TokenLiteral())
	}
	if expr.Right.TokenLiteral() != "fuck" {
		t.Fatalf("Expected TokenLiteral of Value to be `fuck`. Received %s", expr.Right.TokenLiteral())
	}

}

func TestLetDeclarationStatement(t *testing.T) {
	input := "let five = 5;"

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

	stmt, ok := program.Statements[0].(*ast.LetDeclarationStatement)
	if !ok {
		t.Fatalf("program.Statements[0]  is not *ast.LetDeclarationStatement. Received %T", program.Statements[0])
	}
	if stmt.TokenLiteral() != "let" {
		t.Fatalf("Expected TokenLiteral of LetDeclarationStatement to be `let`. Received %s", stmt.TokenLiteral())
	}
	if stmt.Name.Value != "five" {
		t.Fatalf("Expected Value of Name to be `five`. Received %s", stmt.Name.Value)
	}
	if stmt.Name.TokenLiteral() != "five" {
		t.Fatalf("Expected Value of Name to be `five`. Received %s", stmt.Name.TokenLiteral())
	}
	if stmt.Value.TokenLiteral() != "5" {
		t.Fatalf("Expected TokenLiteral of ReturnValue to be `5`. Received %s", stmt.Value.TokenLiteral())
	}

}

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
