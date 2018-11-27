package codegen

import (
	"testing"
	"fmt"
	"strings"
	"github.com/Lebonesco/quack_parser/ast"
	"github.com/Lebonesco/quack_parser/lexer"
	"github.com/Lebonesco/quack_parser/parser"
	"github.com/Lebonesco/quack_parser/typechecker"
	"github.com/Lebonesco/quack_parser/environment"
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
			src: `
			friend = 8;
			friend + 7;`,
			res: `int main() {
				obj_Int* friend;
				friend = int_literal(8);
				friend->clazz->PLUS(int_literal(7));
				return 0;
			 	}`},
		{
			src: `
				 class Pt(x: Int, y: Int) {
				    this.x = x;
				    this.y = y; 

				    def PRINT( ):Pt {
				       "( ".PRINT(); 
				       this.x.PRINT(); 
				       ", ".PRINT();
				       ")".PRINT(); 
				     }

				     def PLUS(other: Pt) {
				         return Pt(this.x+other.x, this.y+other.y); 
				     }
				  }
			`,
			res: `
			struct class_Pt_struct;
			typedef struct class_Pt_struct* class_Pt; 

			typedef struct obj_Pt_struct {
			  class_Pt  clazz;
			  obj_Int x;
			  obj_Int y;
			} * obj_Pt;

			struct  class_Pt_struct  the_class_Pt_struct;  /* So I can use it in PLUS */ 

			struct class_Pt_struct {
			  /* Method table */
			  obj_Pt (*constructor) (obj_Int, obj_Int );  
			  obj_String (*STRING) (obj_Obj);           /* Inherit for now */
			  obj_Pt (*PRINT) (obj_Pt);                 /* Overridden */
			  obj_Boolean (*EQUALS) (obj_Obj, obj_Obj); /* Inherited */
			  obj_Pt (*PLUS) (obj_Pt, obj_Pt);          /* Introduced */
			};

			extern class_Pt the_class_Pt;

			/* Constructor */
			obj_Pt new_Pt(obj_Int x, obj_Int y ) {
			  obj_Pt new_thing = (obj_Pt)
			    malloc(sizeof(struct obj_Pt_struct));
			  new_thing->clazz = the_class_Pt;
			  /* Quack code: 
			    this.x = x;
			    this.y = y; 
			  */
			  new_thing->x = x;
			  new_thing->y = y; 
			  return new_thing; 
			}

			obj_Pt Pt_method_PRINT(obj_Pt this) {
			  obj_String lparen = str_literal("(");
			  lparen->clazz->PRINT(lparen);
			  this->x->clazz->PRINT((obj_Obj) this->x);
			  obj_String comma=str_literal(",");
			  comma->clazz->PRINT(comma);
			  this->y->clazz->PRINT((obj_Obj) this->y);
			  obj_String rparen = str_literal(")");
			  rparen->clazz->PRINT(rparen);
			  return this;
			}

			obj_Pt Pt_method_PLUS(obj_Pt this, obj_Pt other) {
			  obj_Int this_x = this->x;
			  obj_Int other_x = other->x;
			  obj_Int this_y = this->y;
			  obj_Int other_y = other->y; 
			  obj_Int x_sum = this_x->clazz->PLUS(this_x, other_x);
			  obj_Int y_sum = this_y->clazz->PLUS(this_y, other_y); 
			  return the_class_Pt->constructor(x_sum, y_sum); 
			}

			/* The Pt Class (a singleton) */
			struct  class_Pt_struct  the_class_Pt_struct = {
			  new_Pt,     /* Constructor */
			  Obj_method_STRING, 
			  Pt_method_PRINT, 
			  Obj_method_EQUALS,
			  Pt_method_PLUS
			};

			class_Pt the_class_Pt = &the_class_Pt_struct;

				int main() {

					return 0;
				}

			`},
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

		env := environment.CreateEnvironment() // create new environment
		_, typeErr := typechecker.TypeCheck(program, env)
		if typeErr != nil {
			t.Errorf(string(typeErr.Type) + " - " + typeErr.Message.Error())
			continue
		}

		code, err := CodeGen(program)
		if err != nil {
			t.Errorf(err.Error())
		}

		t.Log(code)
		// remove extra spaces
		for _, rep := range []string{" ", "\n", "\t"} {
			code = strings.Replace(code, rep, "", -1)
			test.res = strings.Replace(test.res, rep, "", -1)
		}

		if code != test.res {
			t.Errorf("not match between\n %s \n %s", code, test.res)
		}
	}
}