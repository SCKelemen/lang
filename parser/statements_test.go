package parser

import (
	"ast"
	"fmt"
	"scanner"
	"strings"
	"testing"
	"token"
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

func TestRest(t *testing.T) {
	input := `
let x = 1;
let y = x;
return 5;
return 10;
type alias = x;
`

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

	if len(program.Statements) != 5 {
		t.Fatalf("Program has an unexpected number of statements. Expected 5, received %d", len(program.Statements))
	}

	builder := &ast.AstBuilder{}
	stmt := builder.LetStatement("x", builder.InfixExpression("="))
	stmt := ast.LetDeclarationStatement{Token: token.Token{TokenKind: token.LET, Literal: "let"}, Name: builder.Identifier("x")}
	okay, messages := validateLetStatement(stmt, program.Statements[0].(*ast.LetDeclarationStatement))
	if !okay {
		errors := []string{"Failed to correctly parse statement 1."}
		errors = append(errors, messages...)

		t.Fatal(strings.Join(errors, "; "))
	}
}

func validateReturnStatement(expected, actual *ast.ReturnStatement) (bool, []string) {
	messages := []string{}
	ok := true
	if actual.Token.TokenKind != expected.Token.TokenKind {
		ok = false
		messages = append(messages, fmt.Sprintf("Unexpected TokenKind in *ReturnStatement.Token. Expected %s, received %s",
			expected.Token.TokenKind, actual.Token.TokenKind))

	}

	if actual.Token.Literal != expected.Token.Literal {
		ok = false
		messages = append(messages, fmt.Sprintf("Unexpected Literal in *ReturnStatement.Token. Expected %s, received %s",
			expected.Token.Literal, actual.Token.Literal))
	}

	actual_expr, _ok := actual.ReturnValue.(ast.Expression)
	if !_ok {
		ok = _ok
		messages = append(messages, "Unable to cast Actual.ReturnValue to *ast.Expression.")
	}
	expected_expr, _ok := expected.ReturnValue.(ast.Expression)
	if !_ok {
		ok = _ok
		messages = append(messages, "Unable to cast Expected.ReturnValue to *ast.Expression. Your tests are fucked.")
	}

	if actual_expr.TokenLiteral() != expected_expr.TokenLiteral() {
		ok = false
		messages = append(messages, fmt.Sprintf("Unexpected Literal in *ReturnStatement.ReturnValue. Expected %s, received %s",
			expected_expr.TokenLiteral(), actual_expr.TokenLiteral()))
	}

	return ok, messages
}

func validateLetStatement(expected ast.LetDeclarationStatement, actual *ast.LetDeclarationStatement) (bool, []string) {
	messages := []string{}
	ok := true
	if actual.Token.TokenKind != expected.Token.TokenKind {
		ok = false
		messages = append(messages, fmt.Sprintf("Unexpected TokenKind in *LetDeclarationStatement.Token. Expected %s, received %s",
			expected.Token.TokenKind, actual.Token.TokenKind))

	}

	if actual.Token.Literal != expected.Token.Literal {
		ok = false
		messages = append(messages, fmt.Sprintf("Unexpected Literal in *LetDeclarationStatement.Token. Expected %s, received %s",
			expected.Token.Literal, actual.Token.Literal))
	}

	if actual.Name.Value != expected.Name.Value {
		ok = false
		messages = append(messages, fmt.Sprintf("Unexpected Identifier in *LetDeclarationStatement.Name. Expected %s, received %s",
			expected.Name.Value, actual.Name.Value))
	}

	if actual.Value.TokenLiteral() != expected.Value.TokenLiteral() {
		ok = false
		messages = append(messages, fmt.Sprintf("Unexpected TokenLiteral in *LetDeclarationStatement.Value. Expected %s, received %s",
			expected.Value.TokenLiteral(), actual.Value.TokenLiteral()))
	}

	return ok, messages
}

func ExpressionBuilderTest(t *testing.T) {
	stmt := "let x = 6;"
	builder := ast.NewExpressionBuilder()
	builder.InfixExpression().Operator("=").Right(Expression{})
}
