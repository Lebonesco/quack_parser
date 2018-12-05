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
	obj_Nothing (*count) (obj_Test,obj_Int);
};

extern class_Test the_class_Test;

obj_Test new_Test() {
	obj_Test new_thing = (obj_Test) malloc(sizeof(struct obj_Test_struct));
	new_thing->clazz = the_class_Test;
	return new_thing;
}

obj_Test Test_method_count(obj_Test this, obj_Int x) {
obj_Int count;
obj_Int tmp_330 = int_literal(0);
count = tmp_330;
obj_Int round;
obj_Int tmp_331 = int_literal(50);
round = tmp_331;
goto label_8;
label_9: ;
obj_Int tmp_332 = int_literal(0);
obj_Boolean tmp_333 = count->clazz->EQUALS(count, tmp_332);
if (1 == tmp_333->value) {
	goto label_10;
} else {
goto exit_6;
}
label_10: ;
obj_Int tmp_334 = int_literal(0);
x = tmp_334;
goto exit_6;

exit_6: ;

obj_Int tmp_335 = int_literal(1);
obj_Int tmp_336 = count->clazz->PLUS(count, tmp_335);
count = tmp_336;
obj_Int tmp_337 = int_literal(1);
obj_Int tmp_338 = x->clazz->PLUS(x, tmp_337);
x = tmp_338;
obj_Boolean tmp_339 = x->clazz->LESS(x, round);
label_8: ;
if (1 == tmp_339->value) {
	goto label_9;
}
}

struct class_Test_struct the_class_Test_struct = {
new_Test,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
Test_method_count,
};

class_Test the_class_Test = &the_class_Test_struct;


int main() {
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
