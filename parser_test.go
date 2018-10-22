package main

import (
	"testing"
	"github.com/Lebonesco/quack_parser/lexer"
	"github.com/Lebonesco/quack_parser/token"
	"github.com/Lebonesco/quack_parser/parser"
	"github.com/Lebonesco/quack_parser/ast"
)

type Test struct {
	expectedType    token.Type
	expectedLiteral string
}

func TestScannerToken(t *testing.T) {
	tests := []Test{
		{token.TokMap.Type("let"), "let"},
		{token.TokMap.Type("ident"), "five"},
		{token.TokMap.Type("assign"), "="},
		{token.TokMap.Type("string_literal"), "\"test\""},
		{token.TokMap.Type("semicolon"), ";"},
		{token.TokMap.Type("let"), "let"},
		{token.TokMap.Type("ident"), "ten"},
		{token.TokMap.Type("assign"), "="},
		{token.TokMap.Type("int"), "10"},
		{token.TokMap.Type("semicolon"), ";"},
		{token.TokMap.Type("let"), "let"},
		{token.TokMap.Type("ident"), "add"},
		{token.TokMap.Type("assign"), "="},
		{token.TokMap.Type("ident"), "fn"},
		{token.TokMap.Type("lparen"), "("},
		{token.TokMap.Type("ident"), "x"},
		{token.TokMap.Type("comma"), ","},
		{token.TokMap.Type("ident"), "y"},
		{token.TokMap.Type("rparen"), ")"},
		{token.TokMap.Type("lbrace"), "{"},
		{token.TokMap.Type("ident"), "x"},
		{token.TokMap.Type("plus"), "+"},
		{token.TokMap.Type("ident"), "y"},
		{token.TokMap.Type("semicolon"), ";"},
		{token.TokMap.Type("rbrace"), "}"},
		{token.TokMap.Type("semicolon"), ";"},
		{token.TokMap.Type("let"), "let"},
		{token.TokMap.Type("ident"), "result"},
		{token.TokMap.Type("assign"), "="},
		{token.TokMap.Type("ident"), "add"},
		{token.TokMap.Type("lparen"), "("},
		{token.TokMap.Type("ident"), "five"},
		{token.TokMap.Type("comma"), ","},
		{token.TokMap.Type("ident"), "ten"},
		{token.TokMap.Type("rparen"), ")"},
		{token.TokMap.Type("semicolon"), ";"},
		{token.TokMap.Type("int"), "5"},
		{token.TokMap.Type("lt"), "<"},
		{token.TokMap.Type("int"), "10"},
		{token.TokMap.Type("gt"), ">"},
		{token.TokMap.Type("int"), "5"},
		{token.TokMap.Type("semicolon"), ";"},
		{token.TokMap.Type("if"), "if"},
		{token.TokMap.Type("lparen"), "("},
		{token.TokMap.Type("int"), "5"},
		{token.TokMap.Type("lt"), "<"},
		{token.TokMap.Type("int"), "10"},
		{token.TokMap.Type("rparen"), ")"},
		{token.TokMap.Type("lbrace"), "{"},
		{token.TokMap.Type("return"), "return"},
		{token.TokMap.Type("true"), "true"},
		{token.TokMap.Type("semicolon"), ";"},
		{token.TokMap.Type("rbrace"), "}"},
		{token.TokMap.Type("else"), "else"},
		{token.TokMap.Type("lbrace"), "{"},
		{token.TokMap.Type("return"), "return"},
		{token.TokMap.Type("false"), "false"},
		{token.TokMap.Type("semicolon"), ";"},
		{token.TokMap.Type("rbrace"), "}"},
		{token.TokMap.Type("int"), "10"},
		{token.TokMap.Type("eq"), "=="},
		{token.TokMap.Type("int"), "10"},
		{token.TokMap.Type("semicolon"), ";"},
		{token.TokMap.Type("class"), "class"},
		{token.TokMap.Type("ident"), "Pt"},
		{token.TokMap.Type("lparen"), "("},
		{token.TokMap.Type("ident"), "x"},
		{token.TokMap.Type("colon"), ":"},
		{token.TokMap.Type("ident"), "Int"},
		{token.TokMap.Type("comma"), ","},
		{token.TokMap.Type("ident"), "y"},
		{token.TokMap.Type("colon"), ":"},
		{token.TokMap.Type("ident"), "Int"},
		{token.TokMap.Type("rparen"), ")"},
		{token.TokMap.Type("lbrace"), "{"},
		{token.TokMap.Type("ident"), "thisx"},
		{token.TokMap.Type("assign"), "="},
		{token.TokMap.Type("ident"), "x"},
		{token.TokMap.Type("semicolon"), ";"},
		{token.TokMap.Type("ident"), "thisy"},
		{token.TokMap.Type("assign"), "="},
		{token.TokMap.Type("ident"), "y"},
		{token.TokMap.Type("semicolon"), ";"},
		{token.TokMap.Type("def"), "def"},
		{token.TokMap.Type("ident"), "_x"},
		{token.TokMap.Type("lparen"), "("},
		{token.TokMap.Type("rparen"), ")"},
		{token.TokMap.Type("colon"), ":"},
		{token.TokMap.Type("ident"), "Int"},
		{token.TokMap.Type("lbrace"), "{"},
		{token.TokMap.Type("return"), "return"},
		{token.TokMap.Type("ident"), "thisy"},
		{token.TokMap.Type("semicolon"), ";"},
		{token.TokMap.Type("rbrace"), "}"},
		{token.TokMap.Type("rbrace"), "}"},
		{token.TokMap.Type("$"), ""}, // end token
	}

	runTest(tests, INPUT1, t)
}

func TestScannerStrings(t *testing.T) {
	tests := []Test{
		{token.TokMap.Type("let"), "let"},
		{token.TokMap.Type("ident"), "five"},
		{token.TokMap.Type("assign"), "="},
		{token.TokMap.Type("INVALID"), "\"te\n"},
		{token.TokMap.Type("ident"), "st"},
		{token.TokMap.Type("INVALID"), "\";"},
		{token.TokMap.Type("$"), ""}, // end token
	}

	runTest(tests, INPUT2, t)
}

func TestScannerComments(t *testing.T) {
	tests := []Test{
		{token.TokMap.Type("INVALID"), "/*"},
		{token.TokMap.Type("$"), ""}, // end token
	}

	runTest(tests, INPUT3, t)
}

func TestScannerTripleQuote(t *testing.T) {
	tests := []Test{
		{token.TokMap.Type("$"), ""}, // end token
	}

	runTest(tests, INPUT4, t)
}

func TestScannerEscape(t *testing.T) {
	tests := []Test{
		{token.TokMap.Type("string_escape_error"), "\"invalid \\q escape character\""},
	}

	runTest(tests, INPUT5, t)
}

func TestEndString(t *testing.T) {
	tests := []Test{
		{token.TokMap.Type("ident"), "s"},
		{token.TokMap.Type("colon"), ":"},
		{token.TokMap.Type("ident"), "String"},
		{token.TokMap.Type("assign"), "="},
		{token.TokMap.Type("INVALID"), "\"This cquote runs off the end of the file"},
	}

	runTest(tests, `s: String = "This cquote runs off the end of the file`, t)
}

// pass input through token checker
func runTest(tests []Test, input string, t *testing.T) {
	l := lexer.NewLexer([]byte(input))
	for i, tt := range tests {
		tok := l.Scan()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected='%s', got='%s' at line %d, column %d",
				i, token.TokMap.Id(tt.expectedType), token.TokMap.Id(tok.Type), tok.Pos.Line, tok.Pos.Column)
		}

		if string(tok.Lit) != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected='%q', got='%q' at line %d, column %d",
				i, tt.expectedLiteral, string(tok.Lit), tok.Pos.Line, tok.Pos.Column)
		}
	}
}

func TestAssignment(t *testing.T) {
	tests := []struct{
		src string
		expectedIdentifier string
		expectedValue interface{}
	}{
		{"let five = 5;", "five", 5},
		{"let x = true;", "x", true},
		{"let y = \"k\";", "y", "k"},
		{"let foo = k;", "foo", "k"},
		{"let z = add(five, ten);", "z", ""},
	}

	for _, tt := range tests {
		l := lexer.NewLexer([]byte(tt.src))
		p := parser.NewParser()
		res, err := p.Parse(l)
		if err != nil {
			t.Fatalf(err.Error())
		}

		program, _ := res.(*ast.Program)
		stmt := program.Statements[0]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "LetStatement" {
		t.Errorf("s.TokenLiteral() not 'LetStatement', got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement, got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got='%s'", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got='%s'", name, letStmt.Name)
		return false
	}
	return true
}

func TestOperators(t *testing.T) {
	tests := []struct{
		src string
		expectedLeft int64
		expectedRight int64
		expectedOp string
	}{
		{`5 + 5;`, 5, 5, `+`},
		{`5 + 5;`, 5, 5, `+`},
		{`5 < 6;`, 5, 6, "<"},
		{`5 <= 6;`, 5, 6, "<="},	
		{`5 > 6;`, 5, 6, ">"},	
		{`5 >= 6;`, 5, 6, ">="},
		{"5 == 3;", 5, 3, "=="},
		{"5 != 3;", 5, 3, "!="},		
	}	

	for _, tt := range tests {
		l := lexer.NewLexer([]byte(tt.src))
		p := parser.NewParser()
		res, err := p.Parse(l)
		if err != nil {
			t.Fatalf(err.Error())
		}

		program, _ := res.(*ast.Program)
		s := program.Statements[0]
		
		if s.TokenLiteral() != "ExpressionStatement" {
			t.Errorf("s.TokenLiteral() not 'ExpressionStatement', got=%s", s.TokenLiteral())
		}

		ExpStmt, ok := s.(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("s not *ast.ExpressionStatement, got=%T", s)
		}

		Exp, ok := ExpStmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Errorf("ExpStmt not have *ast.InfixExpression")
		}

		if Exp.Operator != tt.expectedOp {
			t.Errorf("expected operator %s, got=%s", tt.expectedOp, Exp.Operator)
		}

		left, ok := Exp.Left.(*ast.IntegerLiteral)
		if !ok {
			t.Fatalf("not valid left expression, got=%T", Exp.Left)
		}

		right, ok := Exp.Right.(*ast.IntegerLiteral)
		if !ok {
			t.Fatalf("not valid right expression, got=%T", Exp.Right)
		}

		if right.Value != tt.expectedRight {
			t.Fatalf("not correct right value expected %d, got=%d", tt.expectedRight, left.Value)
		}

		if left.Value != tt.expectedLeft {
			t.Fatalf("not correct left value expected %d, got=%d", tt.expectedLeft, left.Value)
		}		
	}
}

func TestBoolOperations(t *testing.T) {
	tests := []struct{
		src string
		expectedLeft bool
		expectedRight bool
		expectedOp string
	}{
		{`true and true;`, true, true, "and"},
		{`true or false;`, true, false, "or"},		
	}	

	for _, tt := range tests {
		l := lexer.NewLexer([]byte(tt.src))
		p := parser.NewParser()
		res, err := p.Parse(l)
		if err != nil {
			t.Fatalf(err.Error())
		}

		program, _ := res.(*ast.Program)
		s := program.Statements[0]
		
		if s.TokenLiteral() != "ExpressionStatement" {
			t.Errorf("s.TokenLiteral() not 'ExpressionStatement', got=%s", s.TokenLiteral())
		}

		ExpStmt, ok := s.(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("s not *ast.ExpressionStatement, got=%T", s)
		}

		Exp, ok := ExpStmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Errorf("ExpStmt not have *ast.InfixExpression")
		}

		if Exp.Operator != tt.expectedOp {
			t.Errorf("expected operator %s, got=%s", tt.expectedOp, Exp.Operator)
		}

		left, ok := Exp.Left.(*ast.Boolean)
		if !ok {
			t.Fatalf("not valid left expression, got=%T", Exp.Left)
		}

		right, ok := Exp.Right.(*ast.Boolean)
		if !ok {
			t.Fatalf("not valid right expression, got=%T", Exp.Right)
		}

		if right.Value != tt.expectedRight {
			t.Fatalf("not correct right value expected %t, got=%t", tt.expectedRight, left.Value)
		}

		if left.Value != tt.expectedLeft {
			t.Fatalf("not correct left value expected %t, got=%t", tt.expectedLeft, left.Value)
		}		
	}
}

func TestIfStatement(t *testing.T) {
	tests := []struct{
		src string
	}{
		{
		`if (5 < 10) {
			return true;
		} elif (false) {
			return false;
		} else {
			return false;
		}`},
	}	

	for _, tt := range tests {
		l := lexer.NewLexer([]byte(tt.src))
		p := parser.NewParser()
		res, err := p.Parse(l)
		if err != nil {
			t.Fatalf(err.Error())
		}

		program, _ := res.(*ast.Program)
		_ = program.Statements[0]
	}
}

func TestWhileStatement(t *testing.T) {
	tests := []struct{
		src string
	}{
		{
		`while i < 4 {
			let tmp = 4;
		}`},
	}	

	for _, tt := range tests {
		l := lexer.NewLexer([]byte(tt.src))
		p := parser.NewParser()
		res, err := p.Parse(l)
		if err != nil {
			t.Fatalf(err.Error())
		}

		program, _ := res.(*ast.Program)
		_ = program.Statements[0]
	}
}

func TestClass(t *testing.T) {
	tests := []struct{
		src string
	}{
		{
		`class Pt(x: Int, y: Int) {
			"""
			example of a class in quack
			"""
			//this.x = x;
			//this.y = y;
				
			//def _x() : Int { return this.y; }
		}`},
	}	

	for _, tt := range tests {
		l := lexer.NewLexer([]byte(tt.src))
		p := parser.NewParser()
		res, err := p.Parse(l)
		if err != nil {
			t.Fatalf(err.Error())
		}

		program, _ := res.(*ast.Program)
		s := program.Classes[0]
		t.Log(s)


	// 	if s.TokenLiteral() != "LetStatement" {
	// 	t.Errorf("s.TokenLiteral() not 'LetStatement', got=%q", s.TokenLiteral())
	// 	return false
	// }

	// letStmt, ok := s.(*ast.LetStatement)
	// if !ok {
	// 	t.Errorf("s not *ast.LetStatement, got=%T", s)
	// 	return false
	// }

	// if letStmt.Name.Value != name {
	// 	t.Errorf("letStmt.Name.Value not '%s'. got='%s'", name, letStmt.Name.Value)
	// 	return false
	// }

	// if letStmt.Name.Value != name {
	// 	t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got='%s'", name, letStmt.Name)
	// 	return false
	// }
	// return true
	}
}
