package typechecker

import (
	"testing"
	"github.com/Lebonesco/quack_parser/lexer"
	"github.com/Lebonesco/quack_parser/parser"
	"github.com/Lebonesco/quack_parser/ast"
)

func TestIfStatement(t *testing.T) {
	tests := []struct {
		src string
		Success bool
	}{
		{
		`if (5 < 10) {
			return true;
		} elif (false) {
			return false;
		} else {
			return false;
		}`,
		true},
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
		if err != nil {
			t.Fatalf(err.Error())
		}
	}
}
