#include <stdio.h>
#include <stdlib.h>
#include "Builtins.h"


struct class_Add_struct;
typedef struct class_Add_struct* class_Add;

typedef struct obj_Add_struct {
	class_Add clazz;
	obj_Int x;
	obj_Int y;
} * obj_Add;

struct class_Add_struct the_class_Add_struct;

struct class_Add_struct {
	obj_Add (*constructor) (obj_Int,obj_Int);
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
	obj_Add (*Plus) (obj_Add,obj_Add);
};

extern class_Add the_class_Add;

obj_Add new_Add(obj_Int y,obj_Int x) {
	obj_Add new_thing = (obj_Add) malloc(sizeof(struct obj_Add_struct));
	new_thing->clazz = the_class_Add;
new_thing->x = x;
new_thing->y = y;
	return new_thing;
}

obj_Add Add_method_Plus(obj_Add this, obj_Add my) {
obj_Int tmp_28 = int_literal(1);
obj_Int tmp_29 = my->x;
obj_Int tmp_30 = my->y;
obj_Int tmp_31 = tmp_29->clazz->PLUS(tmp_29, tmp_30);
obj_Add tmp_32 = the_class_Add->constructor(tmp_28,tmp_31);
return tmp_32;
}

struct class_Add_struct the_class_Add_struct = {
new_Add,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
Add_method_Plus,
};

class_Add the_class_Add = &the_class_Add_struct;


int main() {
obj_Add an_add;
obj_Int tmp_33 = int_literal(5);
obj_Int tmp_34 = int_literal(4);
obj_Add tmp_35 = the_class_Add->constructor(tmp_33,tmp_34);
an_add = tmp_35;
obj_Nothing tmp_36 = an_add->clazz->PRINT(an_add);
tmp_36;
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
