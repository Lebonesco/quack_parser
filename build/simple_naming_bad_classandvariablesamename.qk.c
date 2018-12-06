#include <stdio.h>
#include <stdlib.h>
#include "Builtins.h"


struct class_R_struct;
typedef struct class_R_struct* class_R;

typedef struct obj_R_struct {
	class_R clazz;
	obj_Int R;
} * obj_R;

struct class_R_struct the_class_R_struct;

struct class_R_struct {
	obj_R (*constructor) ();
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
};

extern class_R the_class_R;

obj_R new_R() {
	obj_R new_thing = (obj_R) malloc(sizeof(struct obj_R_struct));
	new_thing->clazz = the_class_R;
obj_Int tmp_364 = int_literal(5);
new_thing->R = tmp_364;
	return new_thing;
}

struct class_R_struct the_class_R_struct = {
new_R,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
};

class_R the_class_R = &the_class_R_struct;


int main() {
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
