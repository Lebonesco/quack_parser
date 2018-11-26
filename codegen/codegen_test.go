package codegen

import (
	"testing"
	"fmt"
	"strings"
	"github.com/Lebonesco/quack_parser/ast"
	"github.com/Lebonesco/quack_parser/lexer"
	"github.com/Lebonesco/quack_parser/parser"
	"github.com/Lebonesco/quack_parser/typechecker"
)

func TestSmall(t *testing.T) {
	tests := []struct{
		src string
		res string
	}{
		{
			src: ``,
			res: `int main() { return 0; }`},
		{
			src: `5;`,
			res: `int main() { int_literal(5); return 0; }`},
		{
			src: `friend + friend;`,
			res: `int main() { }`
		}

	}

	for i, test := range tests {
		fmt.Println("test: ", i)
		l := lexer.NewLexer([]byte(test.src))
		p := parser.NewParser()
		res, err := p.Parse(l)
		if err != nil {
			t.Log(" parse error", err.Error())
			continue
		}

		program, _ := res.(*ast.Program)

		env := typechecker.CreateEnvironment() // create new environment
		_, typeErr := typechecker.TypeCheck(program, env)
		if typeErr != nil {
			t.Errorf(string(typeErr.Type) + " - " + typeErr.Message.Error())
			continue
		}

		code, err := CodeGen(program)
		if err != nil {
			t.Errorf(err.Error())
		}
		// remove extra spaces
		for _, rep := range []string{" ", "\n", "\t"} {
			code = strings.Replace(code, rep, "", -1)
			test.res = strings.Replace(test.res, rep, "", -1)
		}

		if code != test.res {
			t.Errorf("not match between %s \n %s", code, test.res)
		}
	}
}