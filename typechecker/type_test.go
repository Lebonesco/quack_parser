package typechecker

import (
	"github.com/Lebonesco/quack_parser/ast"
	"github.com/Lebonesco/quack_parser/lexer"
	"github.com/Lebonesco/quack_parser/parser"
	"github.com/Lebonesco/quack_parser/environment"
	"testing"
)

func TestIfStatement(t *testing.T) {
	tests := []struct {
		src     string
		Success bool
		Error   error
	}{
		{
			`class C1()  extends Obj {
	   def foo():  Top {
	       return C1();    /* CHANGED */
	   }
	}

	class C2() extends C1 {
	   def foo():  C1 {
	        return C1();    /* Conforms to C1.foo() */
	   }
	}

	class C3() extends C2 { 
	   def foo(): C2 {
	        return C2();   /* Conforms to C2.foo() */
	   }
	}

	class C4() extends C3 {
	    def foo() : C3 {
	         return C3();  /* Conforms to C3.foo() */
	    }
	}

//	x = C4();
//	while ( True ) {
//	   x = x.foo();      /* Type system should reject this */
//	}
	`,
			false,
			nil},
	}

	for _, tt := range tests {
		l := lexer.NewLexer([]byte(tt.src))
		p := parser.NewParser()
		res, err := p.Parse(l)
		if err != nil {
			t.Fatalf(err.Error())
		}

		program, _ := res.(*ast.Program)
		if err != nil {
			t.Fatalf(err.Error())
		}

		env := environment.CreateEnvironment()
		_, err = TypeCheck(program, env)
		if err != nil {
			t.Fatalf(err.Error())
		}
	}
}
