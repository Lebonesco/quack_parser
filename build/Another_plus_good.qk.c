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
	obj_Int (*static_types) (obj_Test,obj_Int);
};

extern class_Test the_class_Test;

obj_Test new_Test() {
	obj_Test new_thing = (obj_Test) malloc(sizeof(struct obj_Test_struct));
	new_thing->clazz = the_class_Test;
	return new_thing;
}

obj_Test Test_method_static_types(obj_Test this, obj_Int x) {
obj_Int z;
obj_Int c;
obj_Int y;
obj_Int tmp_1 = int_literal(0);
obj_Boolean tmp_2 = x->clazz->LESS(x, tmp_1);
if (1 == tmp_2->value) {
	goto label_1;
} else {
obj_Int tmp_3 = int_literal(5);
y = tmp_3;
obj_Int tmp_4 = int_literal(7);
z = tmp_4;
obj_Int tmp_5 = y->clazz->PLUS(y, z);
c = tmp_5;
goto exit_1;
}
label_1: ;
obj_Int tmp_6 = int_literal(1);
y = tmp_6;
obj_Int tmp_7 = int_literal(1);
z = tmp_7;
obj_Int tmp_8 = y->clazz->PLUS(y, z);
c = tmp_8;
goto exit_1;

exit_1: ;

return c;
}

struct class_Test_struct the_class_Test_struct = {
new_Test,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
Test_method_static_types,
};

class_Test the_class_Test = &the_class_Test_struct;


int main() {
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
