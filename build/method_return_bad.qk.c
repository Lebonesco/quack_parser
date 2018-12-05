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
	obj_Nothing (*ReturnString) (obj_Test);
	obj_Nothing (*UseString) (obj_Test);
};

extern class_Test the_class_Test;

obj_Test new_Test() {
	obj_Test new_thing = (obj_Test) malloc(sizeof(struct obj_Test_struct));
	new_thing->clazz = the_class_Test;
	return new_thing;
}

obj_Test Test_method_ReturnString(obj_Test this) {
obj_String x;
obj_String tmp_351 = str_literal("hello");
x = tmp_351;
obj_String y;
obj_String tmp_352 = str_literal(" world");
y = tmp_352;
}

obj_Test Test_method_UseString(obj_Test this) {
obj_Obj z;
obj_String tmp_353 = str_literal("");
z = (obj_Obj) tmp_353;
obj_Nothing tmp_354 = the_class_Test->ReturnString(the_class_Test);
z = (obj_Obj) tmp_354;
}

struct class_Test_struct the_class_Test_struct = {
new_Test,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
Test_method_ReturnString,
Test_method_UseString,
};

class_Test the_class_Test = &the_class_Test_struct;


int main() {
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
