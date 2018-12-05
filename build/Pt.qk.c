#include <stdio.h>
#include <stdlib.h>
#include "Builtins.h"


struct class_Pt_struct;
typedef struct class_Pt_struct* class_Pt;

typedef struct obj_Pt_struct {
	class_Pt clazz;
	obj_Int x;
	obj_Int y;
} * obj_Pt;

struct class_Pt_struct the_class_Pt_struct;

struct class_Pt_struct {
	obj_Pt (*constructor) (obj_Int,obj_Int);
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
	obj_Int (*_get_x) (obj_Pt);
	obj_Int (*_get_y) (obj_Pt);
	obj_Nothing (*translate) (obj_Pt,obj_Int,obj_Int);
	obj_Pt (*PLUS) (obj_Pt,obj_Pt);
};

extern class_Pt the_class_Pt;

obj_Pt new_Pt(obj_Int y,obj_Int x) {
	obj_Pt new_thing = (obj_Pt) malloc(sizeof(struct obj_Pt_struct));
	new_thing->clazz = the_class_Pt;
new_thing->x = x;
new_thing->y = y;
	return new_thing;
}

obj_Pt Pt_method__get_x(obj_Pt this) {
return this->x;
}

obj_Pt Pt_method__get_y(obj_Pt this) {
return this->y;
}

obj_Pt Pt_method_translate(obj_Pt this, obj_Int dy, obj_Int dx) {
obj_Int tmp_37 = this->x->clazz->PLUS(this->x, dx);
this->x = tmp_37;
obj_Int tmp_38 = this->y->clazz->PLUS(this->y, dy);
this->y = tmp_38;
}

obj_Pt Pt_method_PLUS(obj_Pt this, obj_Pt other) {
obj_Int tmp_39 = other->y;
obj_Int tmp_40 = this->y->clazz->PLUS(this->y, tmp_39);
obj_Int tmp_41 = other->x;
obj_Int tmp_42 = this->x->clazz->PLUS(this->x, tmp_41);
obj_Pt tmp_43 = the_class_Pt->constructor(tmp_40,tmp_42);
return tmp_43;
}

struct class_Pt_struct the_class_Pt_struct = {
new_Pt,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
Pt_method__get_x,
Pt_method__get_y,
Pt_method_translate,
Pt_method_PLUS,
};

class_Pt the_class_Pt = &the_class_Pt_struct;


int main() {
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
