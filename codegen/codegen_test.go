package codegen

// import (
// 	"testing"
// 	"fmt"
// 	"strings"
// 	"github.com/Lebonesco/quack_parser/ast"
// 	"github.com/Lebonesco/quack_parser/lexer"
// 	"github.com/Lebonesco/quack_parser/parser"
// 	"github.com/Lebonesco/quack_parser/typechecker"
// 	"github.com/Lebonesco/quack_parser/environment"
// )

// func TestSmall(t *testing.T) {
// 	tests := []struct{
// 		src string
// 		res string
// 	}{
// 		{
// 			src: ``,
// 			res: ` #include <stdio.h>
//         			#include <stdlib.h>
//         			#include "Builtins.h"

// 					int main() { return 0; }`},
// 		{
// 			src: `5;`,
// 			res: ` #include <stdio.h>
// 		        #include <stdlib.h>
// 		        #include "Builtins.h"

// 					int main() { 
// 						obj_Int* tmp_1 = int_literal(5); 
// 						tmp_1;
// 						return 0; 
// 					}`},
// 		{
// 			src: `
// 			friend = 8;
// 			friend + 7;`,
// 			res: ` #include <stdio.h>
//         #include <stdlib.h>
//         #include "Builtins.h"

// int main() {
// 				obj_Int friend;
// 				obj_Int* tmp_2 = int_literal(8);
// 				friend = tmp_2;
// 				obj_Int* tmp_3 = int_literal(7);
// 				obj_Int tmp_4 = friend->clazz->PLUS(friend, tmp_3);
// 				tmp_4;
// 				return 0;
// 			 	}`},
// 		// {
// 		// 	src: `
// 		// 		 class Pt(x: Int, y: Int) {
// 		// 		    this.x = x;
// 		// 		    this.y = y; 

// 		// 		    def PRINT( ): Pt {
// 		// 		       "( ".PRINT(); 
// 		// 		       this.x.PRINT(); 
// 		// 		       ", ".PRINT();
// 		// 		       ")".PRINT(); 
// 		// 		     }

// 		// 		     def PLUS(other: Pt) : Pt {
// 		// 		         return Pt(this.x+other.x, this.y+other.y); 
// 		// 		     }
// 		// 		  }

// 		// 		  x = 11;
// 		// 		  y = 32;
// 		// 		  pt1 = "HELLO";
// 		// 		  pt1 = Pt(x, y);
// 		// 		  pt1.PRINT();
// 		// 	`,
// 		// 	res: `
// 		// 	#include <stdio.h>
//   //       #include <stdlib.h>
//   //       #include "Builtins.h"

// 		// 	struct class_Pt_struct;
// 		// 	typedef struct class_Pt_struct* class_Pt; 

// 		// 	typedef struct obj_Pt_struct {
// 		// 	  class_Pt  clazz;
// 		// 	  obj_Int x;
// 		// 	  obj_Int y;
// 		// 	} * obj_Pt;

// 		// 	struct class_Pt_struct  the_class_Pt_struct;

// 		// 	struct class_Pt_struct {
// 		// 	  obj_Pt (*constructor) (obj_Int, obj_Int );  
// 		// 	  obj_String (*STRING) (obj_Obj);           
// 		// 	  obj_Pt (*PRINT) (obj_Pt);                 
// 		// 	  obj_Boolean (*EQUALS) (obj_Obj, obj_Obj); 
// 		// 	  obj_Pt (*PLUS) (obj_Pt, obj_Pt);     
// 		// 	};

// 		// 	extern class_Pt the_class_Pt;

// 		// 	obj_Pt new_Pt(obj_Int y, obj_Int x ) {
// 		// 	  obj_Pt new_thing = (obj_Pt)
// 		// 	    malloc(sizeof(struct obj_Pt_struct));
// 		// 	  new_thing->clazz = the_class_Pt;
// 		// 	  new_thing->y = y; 
// 		// 	  new_thing->x = x;
// 		// 	  return new_thing; 
// 		// 	}

// 		// 	obj_Pt Pt_method_PRINT(obj_Pt this) {
// 		// 	  obj_String tmp_5 = str_literal("(");
// 		// 	  tmp_5->clazz->PRINT(tmp_5);
// 		// 	  this->x->clazz->PRINT(this->x); 
// 		// 	  obj_String tmp_6 =str_literal(",");
// 		// 	  tmp_6->clazz->PRINT(tmp_6);
// 		// 	  obj_String tmp_7 = str_literal(")");
// 		// 	  tmp_7->clazz->PRINT(tmp_7);
// 		// 	}

// 		// 	obj_Pt Pt_method_PLUS(obj_Pt this, obj_Pt other) {
// 		// 	  obj_Int tmp_8 = other->y;
// 		// 	  obj_Int tmp_9 = this->y->clazz->PLUS(this->y, tmp_8);
// 		// 	  obj_Int tmp_10 = other->x;
// 		// 	  obj_Int tmp_11 = this->x->clazz->PLUS(this->x, tmp_10);
// 		// 	  obj_Pt tmp_12 = the_class_Pt->constructor(tmp_9, tmp_11);
// 		// 	  return tmp_12; 
// 		// 	}

// 		// 	struct  class_Pt_struct  the_class_Pt_struct = {
// 		// 	  new_Pt, 
// 		// 	  Obj_method_STRING, 
// 		// 	  Pt_method_PRINT, 
// 		// 	  Obj_method_EQUALS,
// 		// 	  Pt_method_PLUS,
// 		// 	};

// 		// 	class_Pt the_class_Pt = &the_class_Pt_struct;

// 		// 		int main() {
// 		// 			obj_Int x; 
// 		// 			obj_Int* tmp_13 = int_literal(11);
// 		// 			x = tmp_13;
// 		// 			obj_Int y;
// 		// 			obj_Int* tmp_14 = int_literal(32);  
// 		// 			y = tmp_14;
// 		// 			obj_Obj pt1;
// 		// 			obj_String tmp_15 = str_literal("HELLO");
// 		// 			pt1 = (obj_Obj) tmp_15;
// 		// 			obj_Pt tmp_16 = the_class_Pt->constructor(y, x);
// 		// 			pt1 = (obj_Obj) tmp_16;
// 		// 			pt1->clazz->PRINT(pt1);
// 		// 			return 0;
// 		// 		}
// 		// 	`},
// 			{
// 				src: `
// 					x = 5;
// 					if x < 9 {
// 						x = x + 100;
// 					}

// 					if x > 100 {
// 						y = "hello";
// 					} else {
// 						y = "goodbye";
// 					}
// 					z = y + " " + "Tom";
// 					z.PRINT();
// 				`,
// 				res: `
// 				#include <stdio.h>
//         #include <stdlib.h>
//         #include "Builtins.h"
// 					int main() {
// 						obj_Int x;
// 						obj_Int* tmp_17 = int_literal(5);
// 						x = tmp_17;
// 						obj_Int* tmp_18 = int_literal(9);
// 						obj_Boolean tmp_19 = x->clazz->LESS(x, tmp_18);

// 						if (1 == tmp_19->value) {
// 							goto label_1;
// 						} else {
// 							goto exit_1;
// 						}

// 						label_1: ;
// 							obj_Int* tmp_20 = int_literal(100);
// 							obj_Int tmp_21 = x->clazz->PLUS(x, tmp_20);
// 							x = tmp_21;
// 							goto exit_1;
// 						exit_1: ;

// 						obj_Int* tmp_22 = int_literal(100);
// 						obj_Boolean tmp_23 = x->clazz->MORE(x, tmp_22);
// 						if (1 == tmp_23->value) {
// 							goto label_2;
// 						} else {
// 							obj_String y;
// 							obj_String tmp_24 = str_literal("goodbye");
// 							y = tmp_24;
// 							goto exit_2;
// 						}

// 						label_2: ;
// 							obj_String tmp_25 = str_literal("hello");
// 							y= tmp_25;
// 							goto exit_2;

// 						exit_2: ;

// 						obj_String z;
// 						obj_String tmp_26 = str_literal(" ");
// 						obj_String tmp_27 = y->clazz->PLUS(y, tmp_26);
// 						obj_String tmp_28 = str_literal("Tom");
// 				        obj_String tmp_29 = tmp_27->clazz->PLUS(tmp_27, tmp_28);
// 				        z = tmp_29;
// 				        z->clazz->PRINT(z);
// 				        return 0;
// 					}

// 				`},
// 	}

// 	for i, test := range tests {
// 		fmt.Println("test: ", i)
// 		l := lexer.NewLexer([]byte(test.src))
// 		p := parser.NewParser()
// 		res, err := p.Parse(l)
// 		if err != nil {
// 			t.Log(" parse error", err.Error())
// 			continue
// 		}

// 		program, _ := res.(*ast.Program)

// 		env := environment.CreateEnvironment() // create new environment
// 		_, typeErr := typechecker.TypeCheck(program, env)
// 		if typeErr != nil {
// 			t.Errorf(string(typeErr.Type) + " - " + typeErr.Message.Error())
// 			continue
// 		}

// 		cod, err := CodeGen(program)
// 		if err != nil {
// 			t.Errorf(err.Error())
// 		}

// 		code := cod.String()
// 		test := test.res

// 		tmp := code
// 		// remove extra spaces
// 		for _, rep := range []string{" ", "\n", "\t"} {
// 			code = strings.Replace(code, rep, "", -1)
// 			test = strings.Replace(test, rep, "", -1)
// 		}

// 		if code != test {
// 			t.Log(tmp)
// 			for i := 1; i <= len(code); i++ {
// 				if code[i] != test[i] {
// 					t.Log(string(code[i-5:i+1]), string(test[i-5:i+1]), i)
// 					t.Fatalf("failed")
// 					break
// 				}
// 			}
// 			//t.Fatalf("not match between\n %s \n %s", code, test.res)
// 		}
// 	}
// }