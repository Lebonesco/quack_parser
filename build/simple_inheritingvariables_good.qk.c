#include <stdio.h>
#include <stdlib.h>
#include "Builtins.h"


struct class_C1_struct;
typedef struct class_C1_struct* class_C1;

typedef struct obj_C1_struct {
	class_C1 clazz;
	obj_Int x;
} * obj_C1;


struct class_C2_struct;
typedef struct class_C2_struct* class_C2;

typedef struct obj_C2_struct {
	class_C2 clazz;
	obj_Int x;
} * obj_C2;

struct class_C1_struct the_class_C1_struct;

struct class_C1_struct {
	obj_C1 (*constructor) ();
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
};

extern class_C1 the_class_C1;

struct class_C2_struct the_class_C2_struct;

struct class_C2_struct {
	obj_C2 (*constructor) ();
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
};

extern class_C2 the_class_C2;

obj_C1 new_C1() {
	obj_C1 new_thing = (obj_C1) malloc(sizeof(struct obj_C1_struct));
	new_thing->clazz = the_class_C1;
obj_Int tmp_361 = int_literal(5);
new_thing->x = tmp_361;
	return new_thing;
}

struct class_C1_struct the_class_C1_struct = {
new_C1,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
};

class_C1 the_class_C1 = &the_class_C1_struct;

obj_C2 new_C2() {
	obj_C2 new_thing = (obj_C2) malloc(sizeof(struct obj_C2_struct));
	new_thing->clazz = the_class_C2;
obj_Int tmp_362 = int_literal(8);
new_thing->x = tmp_362;
	return new_thing;
}

struct class_C2_struct the_class_C2_struct = {
new_C2,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
};

class_C2 the_class_C2 = &the_class_C2_struct;


int main() {
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
