#include <stdio.h>
#include <stdlib.h>
#include "Builtins.h"


struct class_Schroedinger_struct;
typedef struct class_Schroedinger_struct* class_Schroedinger;

typedef struct obj_Schroedinger_struct {
	class_Schroedinger clazz;
	obj_Boolean living;
} * obj_Schroedinger;

struct class_Schroedinger_struct the_class_Schroedinger_struct;

struct class_Schroedinger_struct {
	obj_Schroedinger (*constructor) (obj_Int);
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
};

extern class_Schroedinger the_class_Schroedinger;

obj_Schroedinger new_Schroedinger(obj_Int box) {
	obj_Schroedinger new_thing = (obj_Schroedinger) malloc(sizeof(struct obj_Schroedinger_struct));
	new_thing->clazz = the_class_Schroedinger;
obj_Int tmp_356 = int_literal(0);
obj_Boolean tmp_357 = box->clazz->MORE(box, tmp_356);
if (1 == tmp_357->value) {
	goto label_12;
} else {
new_thing->living = lit_false;
goto exit_8;
}
label_12: ;
new_thing->living = lit_true;
goto exit_8;

exit_8: ;

	return new_thing;
}

struct class_Schroedinger_struct the_class_Schroedinger_struct = {
new_Schroedinger,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
};

class_Schroedinger the_class_Schroedinger = &the_class_Schroedinger_struct;


int main() {
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
