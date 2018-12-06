#include <stdio.h>
#include <stdlib.h>
#include "Builtins.h"


struct class_Schroedinger_struct;
typedef struct class_Schroedinger_struct* class_Schroedinger;

typedef struct obj_Schroedinger_struct {
	class_Schroedinger clazz;
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
obj_Int tmp_363 = int_literal(0);
obj_Boolean tmp_364 = box->clazz->MORE(box, tmp_363);
if (1 == tmp_364->value) {
	goto label_15;
} else {
obj_Int tmp_365 = int_literal(0);
this->dead = tmp_365;
goto exit_9;
}
label_15: ;
obj_Int tmp_366 = int_literal(1);
this->living = tmp_366;
goto exit_9;

exit_9: ;

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
