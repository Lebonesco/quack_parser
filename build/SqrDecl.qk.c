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
obj_String tmp_187 = str_literal("(");
obj_String tmp_188 = this->x->clazz->STRING(this->x);
obj_String tmp_189 = tmp_187->clazz->PLUS(tmp_187, tmp_188);
obj_String tmp_190 = str_literal(",");
obj_String tmp_191 = tmp_189->clazz->PLUS(tmp_189, tmp_190);
obj_String tmp_192 = this->y->clazz->STRING(this->y);
obj_String tmp_193 = tmp_191->clazz->PLUS(tmp_191, tmp_192);
obj_String tmp_194 = str_literal(")");
obj_String tmp_195 = tmp_193->clazz->PLUS(tmp_193, tmp_194);
return tmp_195;
}

obj_Pt Pt_method_PLUS(obj_Pt this, obj_Pt other) {
obj_Int tmp_196 = other->y;
obj_Int tmp_197 = this->y->clazz->PLUS(this->y, tmp_196);
obj_Int tmp_198 = other->x;
obj_Int tmp_199 = this->x->clazz->PLUS(this->x, tmp_198);
obj_Pt tmp_200 = the_class_Pt->constructor(tmp_197,tmp_199);
return tmp_200;
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
obj_Int tmp_201 = int_literal(1);
obj_Int tmp_202 = int_literal(1);
obj_Pt tmp_203 = the_class_Pt->constructor(tmp_201,tmp_202);
obj_Pt tmp_204 = this->ur->clazz->PLUS(this->ur, tmp_203);
obj_Int tmp_205 = int_literal(1);
obj_Int tmp_206 = int_literal(1);
obj_Pt tmp_207 = the_class_Pt->constructor(tmp_205,tmp_206);
obj_Pt tmp_208 = this->ll->clazz->PLUS(this->ll, tmp_207);
obj_Rect tmp_209 = the_class_Rect->constructor(tmp_204,tmp_208);
return tmp_209;
}

obj_Rect Rect_method_STRING(obj_Rect this) {
obj_Pt lr;
obj_Int tmp_210 = this->ll->clazz->_x(this->ll);
obj_Int tmp_211 = this->ur->clazz->_y(this->ur);
obj_Pt tmp_212 = the_class_Pt->constructor(tmp_210,tmp_211);
lr = tmp_212;
obj_Pt ul;
obj_Int tmp_213 = this->ur->clazz->_y(this->ur);
obj_Int tmp_214 = this->ll->clazz->_x(this->ll);
obj_Pt tmp_215 = the_class_Pt->constructor(tmp_213,tmp_214);
ul = tmp_215;
obj_String tmp_216 = str_literal("(");
obj_String tmp_217 = this->ll->clazz->STRING(this->ll);
obj_String tmp_218 = tmp_216->clazz->PLUS(tmp_216, tmp_217);
obj_String tmp_219 = str_literal(", ");
obj_String tmp_220 = tmp_218->clazz->PLUS(tmp_218, tmp_219);
obj_String tmp_221 = ul->clazz->STRING(ul);
obj_String tmp_222 = tmp_220->clazz->PLUS(tmp_220, tmp_221);
obj_String tmp_223 = str_literal(",");
obj_String tmp_224 = tmp_222->clazz->PLUS(tmp_222, tmp_223);
obj_String tmp_225 = this->ur->clazz->STRING(this->ur);
obj_String tmp_226 = tmp_224->clazz->PLUS(tmp_224, tmp_225);
obj_String tmp_227 = str_literal(",");
obj_String tmp_228 = tmp_226->clazz->PLUS(tmp_226, tmp_227);
obj_String tmp_229 = lr->clazz->STRING(lr);
obj_String tmp_230 = tmp_228->clazz->PLUS(tmp_228, tmp_229);
obj_String tmp_231 = str_literal(")");
obj_String tmp_232 = tmp_230->clazz->PLUS(tmp_230, tmp_231);
return tmp_232;
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
obj_Int tmp_233 = new_thing->ll->clazz->_y(new_thing->ll);
obj_Int tmp_234 = tmp_233->clazz->PLUS(tmp_233, side);
obj_Int tmp_235 = new_thing->ll->clazz->_x(new_thing->ll);
obj_Int tmp_236 = tmp_235->clazz->PLUS(tmp_235, side);
obj_Pt tmp_237 = the_class_Pt->constructor(tmp_234,tmp_236);
new_thing->ur = tmp_237;
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
obj_Int tmp_238 = int_literal(5);
obj_Int tmp_239 = int_literal(3);
obj_Int tmp_240 = int_literal(3);
obj_Pt tmp_241 = the_class_Pt->constructor(tmp_239,tmp_240);
obj_Square tmp_242 = the_class_Square->constructor(tmp_238,tmp_241);
a_square = (obj_Rect) tmp_242;
obj_Int tmp_243 = int_literal(2);
obj_Int tmp_244 = int_literal(2);
obj_Pt tmp_245 = the_class_Pt->constructor(tmp_243,tmp_244);
obj_Rect tmp_246 = a_square->clazz->translate(a_square,tmp_245);
a_square = tmp_246;
obj_Nothing tmp_247 = a_square->clazz->PRINT(a_square);
tmp_247;
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
