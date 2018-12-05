#include <stdio.h>
#include <stdlib.h>
#include "Builtins.h"


struct class_C1_Fer_struct;
typedef struct class_C1_Fer_struct* class_C1_Fer;

typedef struct obj_C1_Fer_struct {
	class_C1_Fer clazz;
} * obj_C1_Fer;

struct class_C1_Fer_struct the_class_C1_Fer_struct;

struct class_C1_Fer_struct {
	obj_C1_Fer (*constructor) (obj_Boolean,obj_String,obj_Int);
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
};

extern class_C1_Fer the_class_C1_Fer;

obj_C1_Fer new_C1_Fer(obj_Boolean y,obj_String z,obj_Int x) {
	obj_C1_Fer new_thing = (obj_C1_Fer) malloc(sizeof(struct obj_C1_Fer_struct));
	new_thing->clazz = the_class_C1_Fer;
goto label_5;
label_6: ;
obj_Boolean w;
w = lit_true;
obj_Int tmp_325 = int_literal(5);
obj_Boolean tmp_326 = x->clazz->MORE(x, tmp_325);
label_5: ;
if (1 == tmp_326->value) {
	goto label_6;
}
	return new_thing;
}

struct class_C1_Fer_struct the_class_C1_Fer_struct = {
new_C1_Fer,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
};

class_C1_Fer the_class_C1_Fer = &the_class_C1_Fer_struct;


int main() {
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
