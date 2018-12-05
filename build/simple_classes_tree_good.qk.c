#include <stdio.h>
#include <stdlib.h>
#include "Builtins.h"


struct class_A_struct;
typedef struct class_A_struct* class_A;

typedef struct obj_A_struct {
	class_A clazz;
} * obj_A;


struct class_B_struct;
typedef struct class_B_struct* class_B;

typedef struct obj_B_struct {
	class_B clazz;
} * obj_B;


struct class_C_struct;
typedef struct class_C_struct* class_C;

typedef struct obj_C_struct {
	class_C clazz;
} * obj_C;


struct class_D_struct;
typedef struct class_D_struct* class_D;

typedef struct obj_D_struct {
	class_D clazz;
} * obj_D;

struct class_A_struct the_class_A_struct;

struct class_A_struct {
	obj_A (*constructor) ();
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
};

extern class_A the_class_A;

struct class_B_struct the_class_B_struct;

struct class_B_struct {
	obj_B (*constructor) ();
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
};

extern class_B the_class_B;

struct class_C_struct the_class_C_struct;

struct class_C_struct {
	obj_C (*constructor) ();
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
};

extern class_C the_class_C;

struct class_D_struct the_class_D_struct;

struct class_D_struct {
	obj_D (*constructor) ();
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
};

extern class_D the_class_D;

obj_A new_A() {
	obj_A new_thing = (obj_A) malloc(sizeof(struct obj_A_struct));
	new_thing->clazz = the_class_A;
	return new_thing;
}

struct class_A_struct the_class_A_struct = {
new_A,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
};

class_A the_class_A = &the_class_A_struct;

obj_B new_B() {
	obj_B new_thing = (obj_B) malloc(sizeof(struct obj_B_struct));
	new_thing->clazz = the_class_B;
	return new_thing;
}

struct class_B_struct the_class_B_struct = {
new_B,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
};

class_B the_class_B = &the_class_B_struct;

obj_C new_C() {
	obj_C new_thing = (obj_C) malloc(sizeof(struct obj_C_struct));
	new_thing->clazz = the_class_C;
	return new_thing;
}

struct class_C_struct the_class_C_struct = {
new_C,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
};

class_C the_class_C = &the_class_C_struct;

obj_D new_D() {
	obj_D new_thing = (obj_D) malloc(sizeof(struct obj_D_struct));
	new_thing->clazz = the_class_D;
	return new_thing;
}

struct class_D_struct the_class_D_struct = {
new_D,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
};

class_D the_class_D = &the_class_D_struct;


int main() {
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
