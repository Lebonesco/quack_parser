package main

import (
	"testing"
	"github.com/Lebonesco/quack_scanner/lexer"
	"github.com/Lebonesco/quack_scanner/token"
)

func Test1(t *testing.T) {
	var TESTS = []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.TokMap.Type("let"), "let"},
		{token.TokMap.Type("ident"), "five"},
		{token.TokMap.Type("assign"), "="},
		{token.TokMap.Type("int"), "5"},
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

	l := lexer.NewLexer([]byte(INPUT))

	for i, tt := range TESTS {
		tok := l.Scan()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%d, got=%d",
				i, tt.expectedType, tok.Type)
		}

		if string(tok.Lit) != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, string(tok.Lit))
		}
	}
}
