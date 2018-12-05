#include <stdio.h>
#include <stdlib.h>
#include "Builtins.h"


struct class_Point_struct;
typedef struct class_Point_struct* class_Point;

typedef struct obj_Point_struct {
	class_Point clazz;
	obj_Int y;
	obj_Int x;
} * obj_Point;

struct class_Point_struct the_class_Point_struct;

struct class_Point_struct {
	obj_Point (*constructor) (obj_Int,obj_Int);
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Point,obj_Obj);
};

extern class_Point the_class_Point;

obj_Point new_Point(obj_Int y,obj_Int x) {
	obj_Point new_thing = (obj_Point) malloc(sizeof(struct obj_Point_struct));
	new_thing->clazz = the_class_Point;
new_thing->x = x;
new_thing->y = y;
	return new_thing;
}

obj_Point Point_method_EQUALS(obj_Point this, obj_Obj other) {
}

struct class_Point_struct the_class_Point_struct = {
new_Point,
Obj_method_STRING,
Obj_method_PRINT,
Point_method_EQUALS,
};

class_Point the_class_Point = &the_class_Point_struct;


int main() {
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
