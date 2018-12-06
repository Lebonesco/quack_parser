#include <stdio.h>
#include <stdlib.h>
#include "Builtins.h"


struct class_P_struct;
typedef struct class_P_struct* class_P;

typedef struct obj_P_struct {
	class_P clazz;
} * obj_P;


struct class_X_struct;
typedef struct class_X_struct* class_X;

typedef struct obj_X_struct {
	class_X clazz;
} * obj_X;

struct class_P_struct the_class_P_struct;

struct class_P_struct {
	obj_P (*constructor) ();
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
	obj_Int (*s) (obj_P,obj_Int);
};

extern class_P the_class_P;

struct class_X_struct the_class_X_struct;

struct class_X_struct {
	obj_X (*constructor) ();
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
	obj_Int (*s) (obj_X,obj_Int);
};

extern class_X the_class_X;

obj_P new_P() {
	obj_P new_thing = (obj_P) malloc(sizeof(struct obj_P_struct));
	new_thing->clazz = the_class_P;
	return new_thing;
}

obj_P P_method_s(obj_P this, obj_Int x) {
return x;
}

struct class_P_struct the_class_P_struct = {
new_P,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
P_method_s,
};

class_P the_class_P = &the_class_P_struct;

obj_X new_X() {
	obj_X new_thing = (obj_X) malloc(sizeof(struct obj_X_struct));
	new_thing->clazz = the_class_X;
	return new_thing;
}

obj_X X_method_s(obj_X this, obj_Int x) {
obj_Int tmp_365 = int_literal(1);
obj_Int tmp_366 = x->clazz->PLUS(x, tmp_365);
return tmp_366;
}

struct class_X_struct the_class_X_struct = {
new_X,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
X_method_s,
};

class_X the_class_X = &the_class_X_struct;


int main() {
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
