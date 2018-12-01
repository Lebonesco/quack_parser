#include <stdio.h>
#include <stdlib.h>
#include "Builtins.h"


struct class_Test_struct;
typedef struct class_Test_struct* class_Test;

typedef struct obj_Test_struct {
	class_Test clazz;
} * obj_Test;

struct class_Test_struct the_class_Test_struct;

struct class_Test_struct {
	obj_Test (*constructor) ();
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
	obj_Nothing (*IsThisLegal) (obj_Test);
};

extern class_Test the_class_Test;

obj_Test new_Test() {
	obj_Test new_thing = (obj_Test) malloc(sizeof(struct obj_Test_struct));
	new_thing->clazz = the_class_Test;
	return new_thing;
}

obj_Test Test_method_IsThisLegal(obj_Test this) {
obj_Boolean x;
x = lit_true;
obj_Boolean y;
y = lit_false;
x = lit_true;
obj_Boolean tmp_1 = x->clazz->EQUALS(x, lit_false);
if (1 == tmp_1->value) {
	goto label_1;
} else {
goto exit_1;
}
label_1: ;
obj_String z;
obj_String tmp_2 = str_literal("This is weird");
z = tmp_2;
goto exit_1;

exit_1: ;

}

struct class_Test_struct the_class_Test_struct = {
new_Test,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
Test_method_IsThisLegal,
};

class_Test the_class_Test = &the_class_Test_struct;


int main() {
	return 0;
}
