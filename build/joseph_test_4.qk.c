#include <stdio.h>
#include <stdlib.h>
#include "Builtins.h"


struct class_C1_struct;
typedef struct class_C1_struct* class_C1;

typedef struct obj_C1_struct {
	class_C1 clazz;
} * obj_C1;


struct class_C2_struct;
typedef struct class_C2_struct* class_C2;

typedef struct obj_C2_struct {
	class_C2 clazz;
} * obj_C2;


struct class_C3_struct;
typedef struct class_C3_struct* class_C3;

typedef struct obj_C3_struct {
	class_C3 clazz;
} * obj_C3;

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
	obj_Obj (*method1) (obj_C2,obj_C2);
};

extern class_C2 the_class_C2;

struct class_C3_struct the_class_C3_struct;

struct class_C3_struct {
	obj_C3 (*constructor) ();
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
	obj_Int (*method1) (obj_C3,obj_C2);
};

extern class_C3 the_class_C3;

obj_C1 new_C1() {
	obj_C1 new_thing = (obj_C1) malloc(sizeof(struct obj_C1_struct));
	new_thing->clazz = the_class_C1;
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
	return new_thing;
}

obj_C2 C2_method_method1(obj_C2 this, obj_C2 val) {
obj_Int tmp_343 = int_literal(0);
return tmp_343;
}

struct class_C2_struct the_class_C2_struct = {
new_C2,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
C2_method_method1,
};

class_C2 the_class_C2 = &the_class_C2_struct;

obj_C3 new_C3() {
	obj_C3 new_thing = (obj_C3) malloc(sizeof(struct obj_C3_struct));
	new_thing->clazz = the_class_C3;
	return new_thing;
}

obj_C3 C3_method_method1(obj_C3 this, obj_C1 val) {
obj_Int tmp_344 = int_literal(0);
return tmp_344;
}

struct class_C3_struct the_class_C3_struct = {
new_C3,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
C3_method_method1,
};

class_C3 the_class_C3 = &the_class_C3_struct;


int main() {
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
