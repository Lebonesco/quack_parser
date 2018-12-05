#include <stdio.h>
#include <stdlib.h>
#include "Builtins.h"


struct class_C1_struct;
typedef struct class_C1_struct* class_C1;

typedef struct obj_C1_struct {
	class_C1 clazz;
	obj_Int X;
} * obj_C1;

struct class_C1_struct the_class_C1_struct;

struct class_C1_struct {
	obj_C1 (*constructor) ();
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
};

extern class_C1 the_class_C1;

obj_C1 new_C1() {
	obj_C1 new_thing = (obj_C1) malloc(sizeof(struct obj_C1_struct));
	new_thing->clazz = the_class_C1;
obj_Int tmp_384 = int_literal(5);
new_thing->X = tmp_384;
goto label_13;
label_14: ;
obj_Int tmp_385 = int_literal(4);
new_thing->X = tmp_385;
label_13: ;
if (1 == lit_true->value) {
	goto label_14;
}
	return new_thing;
}

struct class_C1_struct the_class_C1_struct = {
new_C1,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
};

class_C1 the_class_C1 = &the_class_C1_struct;


int main() {
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
