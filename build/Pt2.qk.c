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
obj_Int w;
obj_Int tmp_44 = int_literal(5);
obj_Int tmp_45 = int_literal(2);
obj_Int tmp_46 = tmp_44->clazz->TIMES(tmp_44, tmp_45);
obj_Int tmp_47 = int_literal(7);
obj_Int tmp_48 = int_literal(2);
obj_Int tmp_49 = tmp_47->clazz->DIVIDE(tmp_47, tmp_48);
obj_Int tmp_50 = tmp_46->clazz->PLUS(tmp_46, tmp_49);
w = tmp_50;
obj_Int z;
obj_Int tmp_51 = new_thing->x->clazz->PLUS(new_thing->x, new_thing->y);
z = tmp_51;
	return new_thing;
}

obj_Pt Pt_method__get_x(obj_Pt this) {
return this->x;
}

obj_Pt Pt_method__get_y(obj_Pt this) {
return this->y;
}

obj_Pt Pt_method_translate(obj_Pt this, obj_Int dy, obj_Int dx) {
obj_Int tmp_52 = this->x->clazz->PLUS(this->x, dx);
this->x = tmp_52;
obj_Int tmp_53 = this->y->clazz->PLUS(this->y, dy);
this->y = tmp_53;
}

obj_Pt Pt_method_PLUS(obj_Pt this, obj_Pt other) {
obj_Int tmp_54 = other->y;
obj_Int tmp_55 = this->y->clazz->PLUS(this->y, tmp_54);
obj_Int tmp_56 = other->x;
obj_Int tmp_57 = this->x->clazz->PLUS(this->x, tmp_56);
obj_Pt tmp_58 = the_class_Pt->constructor(tmp_55,tmp_57);
return tmp_58;
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
obj_Int tmp_59 = int_literal(42);
obj_Int tmp_60 = int_literal(13);
obj_Pt tmp_61 = the_class_Pt->constructor(tmp_59,tmp_60);
obj_Nothing tmp_62 = tmp_61->clazz->PRINT(tmp_61);
tmp_62;
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
