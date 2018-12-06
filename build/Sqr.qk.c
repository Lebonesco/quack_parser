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
	obj_Pt ll;
	obj_Pt ur;
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
obj_String tmp_63 = str_literal("(");
obj_String tmp_64 = this->x->clazz->STRING(this->x);
obj_String tmp_65 = tmp_63->clazz->PLUS(tmp_63, tmp_64);
obj_String tmp_66 = str_literal(",");
obj_String tmp_67 = tmp_65->clazz->PLUS(tmp_65, tmp_66);
obj_String tmp_68 = this->y->clazz->STRING(this->y);
obj_String tmp_69 = tmp_67->clazz->PLUS(tmp_67, tmp_68);
obj_String tmp_70 = str_literal(")");
obj_String tmp_71 = tmp_69->clazz->PLUS(tmp_69, tmp_70);
return tmp_71;
}

obj_Pt Pt_method_PLUS(obj_Pt this, obj_Pt other) {
obj_Int tmp_72 = other->y;
obj_Int tmp_73 = this->y->clazz->PLUS(this->y, tmp_72);
obj_Int tmp_74 = other->x;
obj_Int tmp_75 = this->x->clazz->PLUS(this->x, tmp_74);
obj_Pt tmp_76 = the_class_Pt->constructor(tmp_73,tmp_75);
return tmp_76;
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
obj_Int tmp_77 = int_literal(1);
obj_Int tmp_78 = int_literal(1);
obj_Pt tmp_79 = the_class_Pt->constructor(tmp_77,tmp_78);
obj_Pt tmp_80 = this->ur->clazz->PLUS(this->ur, tmp_79);
obj_Int tmp_81 = int_literal(1);
obj_Int tmp_82 = int_literal(1);
obj_Pt tmp_83 = the_class_Pt->constructor(tmp_81,tmp_82);
obj_Pt tmp_84 = this->ll->clazz->PLUS(this->ll, tmp_83);
obj_Rect tmp_85 = the_class_Rect->constructor(tmp_80,tmp_84);
return tmp_85;
}

obj_Rect Rect_method_STRING(obj_Rect this) {
obj_Pt lr;
obj_Int tmp_86 = this->ll->clazz->_x(this->ll);
obj_Int tmp_87 = this->ur->clazz->_y(this->ur);
obj_Pt tmp_88 = the_class_Pt->constructor(tmp_86,tmp_87);
lr = tmp_88;
obj_Pt ul;
obj_Int tmp_89 = this->ur->clazz->_y(this->ur);
obj_Int tmp_90 = this->ll->clazz->_x(this->ll);
obj_Pt tmp_91 = the_class_Pt->constructor(tmp_89,tmp_90);
ul = tmp_91;
obj_String tmp_92 = str_literal("(");
obj_String tmp_93 = this->ll->clazz->STRING(this->ll);
obj_String tmp_94 = tmp_92->clazz->PLUS(tmp_92, tmp_93);
obj_String tmp_95 = str_literal(", ");
obj_String tmp_96 = tmp_94->clazz->PLUS(tmp_94, tmp_95);
obj_String tmp_97 = ul->clazz->STRING(ul);
obj_String tmp_98 = tmp_96->clazz->PLUS(tmp_96, tmp_97);
obj_String tmp_99 = str_literal(",");
obj_String tmp_100 = tmp_98->clazz->PLUS(tmp_98, tmp_99);
obj_String tmp_101 = this->ur->clazz->STRING(this->ur);
obj_String tmp_102 = tmp_100->clazz->PLUS(tmp_100, tmp_101);
obj_String tmp_103 = str_literal(",");
obj_String tmp_104 = tmp_102->clazz->PLUS(tmp_102, tmp_103);
obj_String tmp_105 = lr->clazz->STRING(lr);
obj_String tmp_106 = tmp_104->clazz->PLUS(tmp_104, tmp_105);
obj_String tmp_107 = str_literal(")");
obj_String tmp_108 = tmp_106->clazz->PLUS(tmp_106, tmp_107);
return tmp_108;
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
obj_Int tmp_109 = new_thing->ll->clazz->_y(new_thing->ll);
obj_Int tmp_110 = tmp_109->clazz->PLUS(tmp_109, side);
obj_Int tmp_111 = new_thing->ll->clazz->_x(new_thing->ll);
obj_Int tmp_112 = tmp_111->clazz->PLUS(tmp_111, side);
obj_Pt tmp_113 = the_class_Pt->constructor(tmp_110,tmp_112);
new_thing->ur = tmp_113;
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
obj_Int tmp_114 = int_literal(5);
obj_Int tmp_115 = int_literal(3);
obj_Int tmp_116 = int_literal(3);
obj_Pt tmp_117 = the_class_Pt->constructor(tmp_115,tmp_116);
obj_Square tmp_118 = the_class_Square->constructor(tmp_114,tmp_117);
a_square = (obj_Rect) tmp_118;
obj_Int tmp_119 = int_literal(2);
obj_Int tmp_120 = int_literal(2);
obj_Pt tmp_121 = the_class_Pt->constructor(tmp_119,tmp_120);
obj_Rect tmp_122 = a_square->clazz->translate(a_square,tmp_121);
a_square = tmp_122;
obj_Nothing tmp_123 = a_square->clazz->PRINT(a_square);
tmp_123;
obj_String tmp_124 = str_literal("HELLO\n");
obj_Nothing tmp_125 = tmp_124->clazz->PRINT(tmp_124);
tmp_125;
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
