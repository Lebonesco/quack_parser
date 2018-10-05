package main

import (
	"testing"
	"github.com/Lebonesco/quack_scanner/lexer"
	"github.com/Lebonesco/quack_scanner/token"
)

type Test struct {
	expectedType    token.Type
	expectedLiteral string
}

func TestToken(t *testing.T) {
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

	l := lexer.NewLexer([]byte(INPUT1))

	for i, tt := range tests {
		tok := l.Scan()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%d, got=%d at line %d, column %d",
				i, tt.expectedType, tok.Type, tok.Pos.Line, tok.Pos.Column)
		}

		if string(tok.Lit) != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q at line %d, column %d",
				i, tt.expectedLiteral, string(tok.Lit), tok.Pos.Line, tok.Pos.Column)
		}
	}
}

func TestStrings(t *testing.T) {
	tests := []Test{
		{token.TokMap.Type("let"), "let"},
		{token.TokMap.Type("ident"), "five"},
		{token.TokMap.Type("assign"), "="},
		{token.TokMap.Type("INVALID"), "\"te\n"},
		{token.TokMap.Type("ident"), "st"},
		{token.TokMap.Type("INVALID"), "\";"},
		{token.TokMap.Type("$"), ""}, // end token
	}

	l := lexer.NewLexer([]byte(INPUT2))

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

func TestComments(t *testing.T) {
	tests := []Test{
		{token.TokMap.Type("INVALID"), "/*"},
		{token.TokMap.Type("$"), ""}, // end token
	}

	l := lexer.NewLexer([]byte(INPUT3))

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

func TestTripleQuote(t *testing.T) {
	tests := []Test{
		{token.TokMap.Type("$"), ""}, // end token
	}

	l := lexer.NewLexer([]byte(INPUT4))

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



