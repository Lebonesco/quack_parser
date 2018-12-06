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
	obj_Nothing (*static_types) (obj_Test,obj_Int);
};

extern class_Test the_class_Test;

obj_Test new_Test() {
	obj_Test new_thing = (obj_Test) malloc(sizeof(struct obj_Test_struct));
	new_thing->clazz = the_class_Test;
	return new_thing;
}

obj_Test Test_method_static_types(obj_Test this, obj_Int x) {
obj_Obj z;
obj_Obj c;
obj_Obj y;
obj_Int tmp_9 = int_literal(0);
obj_Boolean tmp_10 = x->clazz->LESS(x, tmp_9);
if (1 == tmp_10->value) {
	goto label_2;
} else {
obj_String tmp_11 = str_literal("hello");
y = (obj_Obj) tmp_11;
obj_String tmp_12 = str_literal(" world");
z = (obj_Obj) tmp_12;
obj_String tmp_13 = y->clazz->PLUS(y, z);
c = (obj_Obj) tmp_13;
goto exit_2;
}
label_2: ;
obj_Int tmp_14 = int_literal(1);
y = (obj_Obj) tmp_14;
obj_Int tmp_15 = int_literal(1);
z = (obj_Obj) tmp_15;
obj_Int tmp_16 = y->clazz->PLUS(y, z);
c = (obj_Obj) tmp_16;
goto exit_2;

exit_2: ;

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
