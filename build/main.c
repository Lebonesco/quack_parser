#include <stdio.h>
#include <stdlib.h>
#include "Builtins.h"


struct class_Pt_struct;
typedef struct class_Pt_struct* class_Pt;

typedef struct obj_Pt_struct {
	class_Pt clazz;
	obj_Int y;
	obj_Int x;
} * obj_Pt;


struct class_Rect_struct;
typedef struct class_Rect_struct* class_Rect;

typedef struct obj_Rect_struct {
	class_Rect clazz;
	obj_Pt ll;
	obj_Pt ur;
} * obj_Rect;


struct class_Square_struct;
typedef struct class_Square_struct* class_Square;

typedef struct obj_Square_struct {
	class_Square clazz;
	obj_Pt ur;
	obj_Pt ll;
} * obj_Square;

struct class_Pt_struct the_class_Pt_struct;

struct class_Pt_struct {
	obj_Pt (*constructor) (obj_Int,obj_Int);
	obj_String (*STRING) (obj_Pt);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
	obj_Pt (*PLUS) (obj_Pt,obj_Pt);
	obj_Int (*_x) (obj_Pt);
	obj_Int (*_y) (obj_Pt);
};

extern class_Pt the_class_Pt;

struct class_Rect_struct the_class_Rect_struct;

struct class_Rect_struct {
	obj_Rect (*constructor) (obj_Pt,obj_Pt);
	obj_String (*STRING) (obj_Rect);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
	obj_Rect (*translate) (obj_Rect,obj_Pt);
};

extern class_Rect the_class_Rect;

struct class_Square_struct the_class_Square_struct;

struct class_Square_struct {
	obj_Square (*constructor) (obj_Int,obj_Pt);
	obj_String (*STRING) (obj_Rect);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
	obj_Rect (*translate) (obj_Rect,obj_Pt);
};

extern class_Square the_class_Square;

obj_Pt new_Pt(obj_Int y,obj_Int x) {
	obj_Pt new_thing = (obj_Pt) malloc(sizeof(struct obj_Pt_struct));
	new_thing->clazz = the_class_Pt;
new_thing->x = x;
new_thing->y = y;
	return new_thing;
}

obj_Pt Pt_method_STRING(obj_Pt this) {
obj_String tmp_1 = str_literal("(");
obj_String tmp_2 = this->x->clazz->STRING(this->x);
obj_String tmp_3 = tmp_1->clazz->PLUS(tmp_1, tmp_2);
obj_String tmp_4 = str_literal(",");
obj_String tmp_5 = tmp_3->clazz->PLUS(tmp_3, tmp_4);
obj_String tmp_6 = this->y->clazz->STRING(this->y);
obj_String tmp_7 = tmp_5->clazz->PLUS(tmp_5, tmp_6);
obj_String tmp_8 = str_literal(")");
obj_String tmp_9 = tmp_7->clazz->PLUS(tmp_7, tmp_8);
return tmp_9;
}

obj_Pt Pt_method_PLUS(obj_Pt this, obj_Pt other) {
obj_Int tmp_10 = other->y;
obj_Int tmp_11 = this->y->clazz->PLUS(this->y, tmp_10);
obj_Int tmp_12 = other->x;
obj_Int tmp_13 = this->x->clazz->PLUS(this->x, tmp_12);
obj_Pt tmp_14 = the_class_Pt->constructor(tmp_11,tmp_13);
return tmp_14;
}

obj_Pt Pt_method__x(obj_Pt this) {
return this->x;
}

obj_Pt Pt_method__y(obj_Pt this) {
return this->y;
}

struct class_Pt_struct the_class_Pt_struct = {
new_Pt,
Pt_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
Pt_method_PLUS,
Pt_method__x,
Pt_method__y,
};

class_Pt the_class_Pt = &the_class_Pt_struct;

obj_Rect new_Rect(obj_Pt ur,obj_Pt ll) {
	obj_Rect new_thing = (obj_Rect) malloc(sizeof(struct obj_Rect_struct));
	new_thing->clazz = the_class_Rect;
new_thing->ll = ll;
new_thing->ur = ur;
	return new_thing;
}

obj_Rect Rect_method_translate(obj_Rect this, obj_Pt delta) {
obj_Int tmp_15 = int_literal(1);
obj_Int tmp_16 = int_literal(1);
obj_Pt tmp_17 = the_class_Pt->constructor(tmp_15,tmp_16);
obj_Pt tmp_18 = this->ur->clazz->PLUS(this->ur, tmp_17);
obj_Int tmp_19 = int_literal(1);
obj_Int tmp_20 = int_literal(1);
obj_Pt tmp_21 = the_class_Pt->constructor(tmp_19,tmp_20);
obj_Pt tmp_22 = this->ll->clazz->PLUS(this->ll, tmp_21);
obj_Rect tmp_23 = the_class_Rect->constructor(tmp_18,tmp_22);
return tmp_23;
}

obj_Rect Rect_method_STRING(obj_Rect this) {
obj_Pt lr;
obj_Int tmp_24 = this->ll->clazz->_x(this->ll);
obj_Int tmp_25 = this->ur->clazz->_y(this->ur);
obj_Pt tmp_26 = the_class_Pt->constructor(tmp_24,tmp_25);
lr = tmp_26;
obj_Pt ul;
obj_Int tmp_27 = this->ur->clazz->_y(this->ur);
obj_Int tmp_28 = this->ll->clazz->_x(this->ll);
obj_Pt tmp_29 = the_class_Pt->constructor(tmp_27,tmp_28);
ul = tmp_29;
obj_String tmp_30 = str_literal("(");
obj_String tmp_31 = this->ll->clazz->STRING(this->ll);
obj_String tmp_32 = tmp_30->clazz->PLUS(tmp_30, tmp_31);
obj_String tmp_33 = str_literal(", ");
obj_String tmp_34 = tmp_32->clazz->PLUS(tmp_32, tmp_33);
obj_String tmp_35 = ul->clazz->STRING(ul);
obj_String tmp_36 = tmp_34->clazz->PLUS(tmp_34, tmp_35);
obj_String tmp_37 = str_literal(",");
obj_String tmp_38 = tmp_36->clazz->PLUS(tmp_36, tmp_37);
obj_String tmp_39 = this->ur->clazz->STRING(this->ur);
obj_String tmp_40 = tmp_38->clazz->PLUS(tmp_38, tmp_39);
obj_String tmp_41 = str_literal(",");
obj_String tmp_42 = tmp_40->clazz->PLUS(tmp_40, tmp_41);
obj_String tmp_43 = lr->clazz->STRING(lr);
obj_String tmp_44 = tmp_42->clazz->PLUS(tmp_42, tmp_43);
obj_String tmp_45 = str_literal(")");
obj_String tmp_46 = tmp_44->clazz->PLUS(tmp_44, tmp_45);
return tmp_46;
}

struct class_Rect_struct the_class_Rect_struct = {
new_Rect,
Rect_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
Rect_method_translate,
};

class_Rect the_class_Rect = &the_class_Rect_struct;

obj_Square new_Square(obj_Int side,obj_Pt ll) {
	obj_Square new_thing = (obj_Square) malloc(sizeof(struct obj_Square_struct));
	new_thing->clazz = the_class_Square;
new_thing->ll = ll;
obj_Int tmp_47 = new_thing->ll->clazz->_y(new_thing->ll);
obj_Int tmp_48 = tmp_47->clazz->PLUS(tmp_47, side);
obj_Int tmp_49 = new_thing->ll->clazz->_x(new_thing->ll);
obj_Int tmp_50 = tmp_49->clazz->PLUS(tmp_49, side);
obj_Pt tmp_51 = the_class_Pt->constructor(tmp_48,tmp_50);
new_thing->ur = tmp_51;
	return new_thing;
}

struct class_Square_struct the_class_Square_struct = {
new_Square,
Rect_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
Rect_method_translate,
};

class_Square the_class_Square = &the_class_Square_struct;


int main() {
obj_Rect a_square;
obj_Int tmp_52 = int_literal(5);
obj_Int tmp_53 = int_literal(3);
obj_Int tmp_54 = int_literal(3);
obj_Pt tmp_55 = the_class_Pt->constructor(tmp_53,tmp_54);
obj_Square tmp_56 = the_class_Square->constructor(tmp_52,tmp_55);
a_square = (obj_Rect) tmp_56;
obj_Int tmp_57 = int_literal(2);
obj_Int tmp_58 = int_literal(2);
obj_Pt tmp_59 = the_class_Pt->constructor(tmp_57,tmp_58);
obj_Rect tmp_60 = a_square->clazz->translate(a_square,tmp_59);
a_square = tmp_60;
obj_Nothing tmp_61 = a_square->clazz->PRINT(a_square);
tmp_61;
obj_String tmp_62 = str_literal("HELLO\n");
obj_Nothing tmp_63 = tmp_62->clazz->PRINT(tmp_62);
tmp_63;
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
