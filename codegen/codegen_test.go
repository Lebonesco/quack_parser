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
				obj_Int tmp_4 = friend->clazz->PLUS(friend, tmp_3);
				tmp_4;
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

				  x = 11;
				  y = 32;
				  pt1 = "HELLO";
				  pt1 = Pt(x, y);
				  pt1.PRINT();
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
			  obj_Pt (*constructor) (obj_Int, obj_Int );  
			  obj_String (*STR) (obj_Obj);           
			  obj_Pt (*PRINT) (obj_Pt);                 
			  obj_Boolean (*EQUALS) (obj_Obj, obj_Obj); 
			  obj_Pt (*PLUS) (obj_Pt, obj_Pt);     
			};

			extern class_Pt the_class_Pt;

			obj_Pt new_Pt(obj_Int y, obj_Int x ) {
			  obj_Pt new_thing = (obj_Pt)
			    malloc(sizeof(struct obj_Pt_struct));
			  new_thing->clazz = the_class_Pt;
			  new_thing->y = y; 
			  new_thing->x = x;
			  return new_thing; 
			}

			obj_Pt Pt_method_PRINT(obj_Pt this) {
			  obj_String tmp_5 = str_literal("(");
			  tmp_5->clazz->PRINT(tmp_5);
			  this->x->clazz->PRINT(this->x); 
			  obj_String tmp_6 =str_literal(",");
			  tmp_6->clazz->PRINT(tmp_6);
			  obj_String tmp_7 = str_literal(")");
			  tmp_7->clazz->PRINT(tmp_7);
			}

			obj_Pt Pt_method_PLUS(obj_Pt this, obj_Pt other) {
			  obj_Int tmp_8 = other->y;
			  obj_Int tmp_9 = this->y->clazz->PLUS(this->y, tmp_8);
			  obj_Int tmp_10 = other->x;
			  obj_Int tmp_11 = this->x->clazz->PLUS(this->x, tmp_10);
			  obj_Pt tmp_12 = the_class_Pt->clazz->constructor(tmp_9, tmp_11);
			  return tmp_12; 
			}

			struct  class_Pt_struct  the_class_Pt_struct = {
			  new_Pt, 
			  Obj_method_STR, 
			  Pt_method_PRINT, 
			  Obj_method_EQUALS,
			  Pt_method_PLUS,
			};

			class_Pt the_class_Pt = &the_class_Pt_struct;

				int main() {
					obj_Int* x; 
					obj_Int* tmp_13 = int_literal(11);
					x = tmp_13;
					obj_Int* y;
					obj_Int* tmp_14 = int_literal(32);  
					y = tmp_14;
					obj_Obj* pt1;
					obj_String tmp_15 = (obj_Obj) str_literal("HELLO");
					pt1 = tmp_15;
					obj_Obj tmp_16 = (obj_Obj) the_class_Pt->constructor(x,y);
					pt1 = tmp_16;
					pt1->clazz->PRINT(pt1);
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
			for i := 1; i <= len(code); i++ {
				if code[i] != test.res[i] {
					t.Log(string(code[i-50:i+1]), string(test.res[i-50:i+1]), i)
					break
				}
			}
			//t.Fatalf("not match between\n %s \n %s", code, test.res)
		}
	}
}