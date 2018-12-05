#include <stdio.h>
#include <stdlib.h>
#include "Builtins.h"


struct class_A_struct;
typedef struct class_A_struct* class_A;

typedef struct obj_A_struct {
	class_A clazz;
} * obj_A;

struct class_A_struct the_class_A_struct;

struct class_A_struct {
	obj_A (*constructor) ();
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
	obj_Int (*foo) (obj_A);
};

extern class_A the_class_A;

obj_A new_A() {
	obj_A new_thing = (obj_A) malloc(sizeof(struct obj_A_struct));
	new_thing->clazz = the_class_A;
	return new_thing;
}

obj_A A_method_foo(obj_A this) {
obj_Int tmp_363 = int_literal(42);
return tmp_363;
}

struct class_A_struct the_class_A_struct = {
new_A,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
A_method_foo,
};

class_A the_class_A = &the_class_A_struct;


int main() {
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
