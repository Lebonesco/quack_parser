#include <stdio.h>
#include <stdlib.h>
#include "Builtins.h"


struct class_Foo_struct;
typedef struct class_Foo_struct* class_Foo;

typedef struct obj_Foo_struct {
	class_Foo clazz;
} * obj_Foo;


struct class_Bar_struct;
typedef struct class_Bar_struct* class_Bar;

typedef struct obj_Bar_struct {
	class_Bar clazz;
} * obj_Bar;

struct class_Foo_struct the_class_Foo_struct;

struct class_Foo_struct {
	obj_Foo (*constructor) ();
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
};

extern class_Foo the_class_Foo;

struct class_Bar_struct the_class_Bar_struct;

struct class_Bar_struct {
	obj_Bar (*constructor) ();
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
};

extern class_Bar the_class_Bar;

obj_Foo new_Foo() {
	obj_Foo new_thing = (obj_Foo) malloc(sizeof(struct obj_Foo_struct));
	new_thing->clazz = the_class_Foo;
	return new_thing;
}

struct class_Foo_struct the_class_Foo_struct = {
new_Foo,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
};

class_Foo the_class_Foo = &the_class_Foo_struct;

obj_Bar new_Bar() {
	obj_Bar new_thing = (obj_Bar) malloc(sizeof(struct obj_Bar_struct));
	new_thing->clazz = the_class_Bar;
	return new_thing;
}

struct class_Bar_struct the_class_Bar_struct = {
new_Bar,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
};

class_Bar the_class_Bar = &the_class_Bar_struct;


int main() {
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
