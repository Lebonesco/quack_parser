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
			res: `int main() { 
				obj_Int* tmp_1 = int_literal(5); 
				tmp_1;
				return 0; 
				}`},
		{
			src: `
			friend = 8;
			friend + 7;`,
			res: `int main() {
				obj_Int* friend;
				obj_Int* tmp_2 = int_literal(8);
				friend = tmp_2;
				obj_Int* tmp_3 = int_literal(7);
				friend->clazz->PLUS(tmp_3);
				return 0;
			 	}`},
		{
			src: `
				 class Pt(x: Int, y: Int) {
				    this.x = x;
				    this.y = y; 

				    def PRINT( ): Pt {
				       "( ".PRINT(); 
				       this.x.PRINT(); 
				       ", ".PRINT();
				       ")".PRINT(); 
				     }

				     def PLUS(other: Pt) : Pt {
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

			struct class_Pt_struct  the_class_Pt_struct;

			struct class_Pt_struct {
			  /* Method table */
			  obj_Pt (*constructor) (obj_Int, obj_Int );  
			  obj_String (*STR) (obj_Obj);           /* Inherit for now */
			  obj_Pt (*PRINT) (obj_Pt);                 /* Overridden */
			  obj_Boolean (*EQUALS) (obj_Obj, obj_Obj); /* Inherited */
			  obj_Pt (*PLUS) (obj_Pt, obj_Pt);          /* Introduced */
			};

			extern class_Pt the_class_Pt;

			obj_Pt new_Pt(obj_Int x, obj_Int y ) {
			  obj_Pt new_thing = (obj_Pt)
			    malloc(sizeof(struct obj_Pt_struct));
			  new_thing->clazz = the_class_Pt;
			  new_thing->x = x;
			  new_thing->y = y; 
			  return new_thing; 
			}

			obj_Pt Pt_method_PRINT(obj_Pt this) {
			  obj_String tmp_5 = str_literal("(");
			  tmp_5->clazz->PRINT(tmp_5);
			  this->x->clazz->PRINT((obj_Obj) this->x);
			  obj_String tmp_6 =str_literal(",");
			  tmp_6->clazz->PRINT(tmp_6);
			  obj_String tmp_7 = str_literal(")");
			  tmp_7->clazz->PRINT(tmp_7);
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
			  new_Pt, 
			  Obj_method_STR, 
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

		tmp := code
		// remove extra spaces
		for _, rep := range []string{" ", "\n", "\t"} {
			code = strings.Replace(code, rep, "", -1)
			test.res = strings.Replace(test.res, rep, "", -1)
		}

		if code != test.res {
			t.Log(tmp)
			//t.Fatalf("not match between\n %s \n %s", code, test.res)
		}
	}
}