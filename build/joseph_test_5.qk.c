#include <stdio.h>
#include <stdlib.h>
#include "Builtins.h"


struct class_Test_struct;
typedef struct class_Test_struct* class_Test;

typedef struct obj_Test_struct {
	class_Test clazz;
} * obj_Test;

struct class_Test_struct the_class_Test_struct;

struct class_Test_struct {
	obj_Test (*constructor) ();
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
	obj_Int (*method1) (obj_Test,obj_Int,obj_Int);
};

extern class_Test the_class_Test;

obj_Test new_Test() {
	obj_Test new_thing = (obj_Test) malloc(sizeof(struct obj_Test_struct));
	new_thing->clazz = the_class_Test;
	return new_thing;
}

obj_Test Test_method_method1(obj_Test this, obj_Int x, obj_Int y) {
obj_Boolean tmp_345 = y->clazz->LESS(y, x);
if (1 == tmp_345->value) {
	goto label_11;
} else {
return y;
goto exit_7;
}
label_11: ;
return x;
goto exit_7;

exit_7: ;

return lit_true;
}

struct class_Test_struct the_class_Test_struct = {
new_Test,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
Test_method_method1,
};

class_Test the_class_Test = &the_class_Test_struct;


int main() {
obj_Int z;
obj_Test tmp_346 = the_class_Test->constructor();
obj_Int tmp_347 = int_literal(2);
obj_Int tmp_348 = int_literal(1);
obj_Int tmp_349 = tmp_346->clazz->method1(tmp_346,tmp_347,tmp_348);
z = tmp_349;
obj_Nothing tmp_350 = z->clazz->PRINT(z);
tmp_350;
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
