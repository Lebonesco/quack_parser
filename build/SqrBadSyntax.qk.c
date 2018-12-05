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
obj_String tmp_126 = str_literal("(");
obj_String tmp_127 = this->x->clazz->STRING(this->x);
obj_String tmp_128 = tmp_126->clazz->PLUS(tmp_126, tmp_127);
obj_String tmp_129 = str_literal(",");
obj_String tmp_130 = tmp_128->clazz->PLUS(tmp_128, tmp_129);
obj_String tmp_131 = this->y->clazz->STRING(this->y);
obj_String tmp_132 = tmp_130->clazz->PLUS(tmp_130, tmp_131);
obj_String tmp_133 = str_literal(")");
obj_String tmp_134 = tmp_132->clazz->PLUS(tmp_132, tmp_133);
return tmp_134;
}

obj_Pt Pt_method_PLUS(obj_Pt this, obj_Pt other) {
obj_Int tmp_135 = other->y;
obj_Int tmp_136 = this->y->clazz->PLUS(this->y, tmp_135);
obj_Int tmp_137 = other->x;
obj_Int tmp_138 = this->x->clazz->PLUS(this->x, tmp_137);
obj_Pt tmp_139 = the_class_Pt->constructor(tmp_136,tmp_138);
return tmp_139;
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
obj_Int tmp_140 = int_literal(2);
obj_Int tmp_141 = int_literal(1);
obj_Pt tmp_142 = the_class_Pt->constructor(tmp_140,tmp_141);
obj_Pt tmp_143 = this->ur->clazz->PLUS(this->ur, tmp_142);
obj_Int tmp_144 = int_literal(1);
obj_Int tmp_145 = int_literal(1);
obj_Pt tmp_146 = the_class_Pt->constructor(tmp_144,tmp_145);
obj_Pt tmp_147 = this->ll->clazz->PLUS(this->ll, tmp_146);
obj_Rect tmp_148 = the_class_Rect->constructor(tmp_143,tmp_147);
return tmp_148;
}

obj_Rect Rect_method_STRING(obj_Rect this) {
obj_Pt lr;
obj_Int tmp_149 = this->ll->clazz->_x(this->ll);
obj_Int tmp_150 = this->ur->clazz->_y(this->ur);
obj_Pt tmp_151 = the_class_Pt->constructor(tmp_149,tmp_150);
lr = tmp_151;
obj_Pt ul;
obj_Int tmp_152 = this->ur->clazz->_y(this->ur);
obj_Int tmp_153 = this->ll->clazz->_x(this->ll);
obj_Pt tmp_154 = the_class_Pt->constructor(tmp_152,tmp_153);
ul = tmp_154;
obj_String tmp_155 = str_literal("(");
obj_String tmp_156 = this->ll->clazz->STRING(this->ll);
obj_String tmp_157 = tmp_155->clazz->PLUS(tmp_155, tmp_156);
obj_String tmp_158 = str_literal(", ");
obj_String tmp_159 = tmp_157->clazz->PLUS(tmp_157, tmp_158);
obj_String tmp_160 = ul->clazz->STRING(ul);
obj_String tmp_161 = tmp_159->clazz->PLUS(tmp_159, tmp_160);
obj_String tmp_162 = str_literal(",");
obj_String tmp_163 = tmp_161->clazz->PLUS(tmp_161, tmp_162);
obj_String tmp_164 = this->ur->clazz->STRING(this->ur);
obj_String tmp_165 = tmp_163->clazz->PLUS(tmp_163, tmp_164);
obj_String tmp_166 = str_literal(",");
obj_String tmp_167 = tmp_165->clazz->PLUS(tmp_165, tmp_166);
obj_String tmp_168 = lr->clazz->STRING(lr);
obj_String tmp_169 = tmp_167->clazz->PLUS(tmp_167, tmp_168);
obj_String tmp_170 = str_literal(")");
obj_String tmp_171 = tmp_169->clazz->PLUS(tmp_169, tmp_170);
return tmp_171;
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
obj_Int tmp_172 = new_thing->ll->clazz->_y(new_thing->ll);
obj_Int tmp_173 = tmp_172->clazz->PLUS(tmp_172, side);
obj_Int tmp_174 = new_thing->ll->clazz->_x(new_thing->ll);
obj_Int tmp_175 = tmp_174->clazz->PLUS(tmp_174, side);
obj_Pt tmp_176 = the_class_Pt->constructor(tmp_173,tmp_175);
new_thing->ur = tmp_176;
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
obj_Int tmp_177 = int_literal(5);
obj_Int tmp_178 = int_literal(3);
obj_Int tmp_179 = int_literal(3);
obj_Pt tmp_180 = the_class_Pt->constructor(tmp_178,tmp_179);
obj_Square tmp_181 = the_class_Square->constructor(tmp_177,tmp_180);
a_square = (obj_Rect) tmp_181;
obj_Int tmp_182 = int_literal(2);
obj_Int tmp_183 = int_literal(2);
obj_Pt tmp_184 = the_class_Pt->constructor(tmp_182,tmp_183);
obj_Rect tmp_185 = a_square->clazz->translate(a_square,tmp_184);
a_square = tmp_185;
obj_Nothing tmp_186 = a_square->clazz->PRINT(a_square);
tmp_186;
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
