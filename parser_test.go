package main

import (
	"testing"
	"fmt"
	"github.com/Lebonesco/quack_parser/lexer"
	"github.com/Lebonesco/quack_parser/token"
	"github.com/Lebonesco/quack_parser/parser"
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

func TestParserMath(t *testing.T) {
	tests := []struct{
		src string
		expect int64
	}{
		{"1 + 1", 2},
		{"1 + 2 ", 3},
	}

	p := parser.NewParser()
	pass := true
	for _, ts := range tests {
		s := lexer.NewLexer([]byte(ts.src))
		sum, err := p.Parse(s)
		if err != nil {
			pass = false
			t.Log(err.Error())
		}
		if sum != ts.expect {
			pass = false
			t.Log(fmt.Sprintf("Error: %s = %d. Got %d\n", ts.src, ts.expect, sum))
		}
	}
	if !pass {
		t.Fail()
	}
}
