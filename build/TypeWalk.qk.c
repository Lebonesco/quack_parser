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


struct class_C4_struct;
typedef struct class_C4_struct* class_C4;

typedef struct obj_C4_struct {
	class_C4 clazz;
} * obj_C4;

struct class_C1_struct the_class_C1_struct;

struct class_C1_struct {
	obj_C1 (*constructor) ();
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
	obj_Obj (*foo) (obj_C1);
};

extern class_C1 the_class_C1;

struct class_C2_struct the_class_C2_struct;

struct class_C2_struct {
	obj_C2 (*constructor) ();
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
	obj_C1 (*foo) (obj_C2);
};

extern class_C2 the_class_C2;

struct class_C3_struct the_class_C3_struct;

struct class_C3_struct {
	obj_C3 (*constructor) ();
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
	obj_C2 (*foo) (obj_C3);
};

extern class_C3 the_class_C3;

struct class_C4_struct the_class_C4_struct;

struct class_C4_struct {
	obj_C4 (*constructor) ();
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
	obj_C3 (*foo) (obj_C4);
};

extern class_C4 the_class_C4;

obj_C1 new_C1() {
	obj_C1 new_thing = (obj_C1) malloc(sizeof(struct obj_C1_struct));
	new_thing->clazz = the_class_C1;
	return new_thing;
}

obj_C1 C1_method_foo(obj_C1 this) {
obj_Obj tmp_325 = the_class_Obj->constructor();
return tmp_325;
}

struct class_C1_struct the_class_C1_struct = {
new_C1,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
C1_method_foo,
};

class_C1 the_class_C1 = &the_class_C1_struct;

obj_C2 new_C2() {
	obj_C2 new_thing = (obj_C2) malloc(sizeof(struct obj_C2_struct));
	new_thing->clazz = the_class_C2;
	return new_thing;
}

obj_C2 C2_method_foo(obj_C2 this) {
obj_C1 tmp_326 = the_class_C1->constructor();
return tmp_326;
}

struct class_C2_struct the_class_C2_struct = {
new_C2,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
C2_method_foo,
};

class_C2 the_class_C2 = &the_class_C2_struct;

obj_C3 new_C3() {
	obj_C3 new_thing = (obj_C3) malloc(sizeof(struct obj_C3_struct));
	new_thing->clazz = the_class_C3;
	return new_thing;
}

obj_C3 C3_method_foo(obj_C3 this) {
obj_C2 tmp_327 = the_class_C2->constructor();
return tmp_327;
}

struct class_C3_struct the_class_C3_struct = {
new_C3,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
C3_method_foo,
};

class_C3 the_class_C3 = &the_class_C3_struct;

obj_C4 new_C4() {
	obj_C4 new_thing = (obj_C4) malloc(sizeof(struct obj_C4_struct));
	new_thing->clazz = the_class_C4;
	return new_thing;
}

obj_C4 C4_method_foo(obj_C4 this) {
obj_C3 tmp_328 = the_class_C3->constructor();
return tmp_328;
}

struct class_C4_struct the_class_C4_struct = {
new_C4,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
C4_method_foo,
};

class_C4 the_class_C4 = &the_class_C4_struct;


int main() {
obj_C4 x;
obj_C4 tmp_329 = the_class_C4->constructor();
x = tmp_329;
goto label_5;
label_6: ;
obj_C3 tmp_330 = x->clazz->foo(x);
x = (obj_C4) tmp_330;
obj_Boolean tmp_331 = x->clazz->EQUALS(x, x);
label_5: ;
if (1 == tmp_331->value) {
	goto label_6;
}
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
