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
	obj_Pt ur;
	obj_Pt ll;
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
	obj_Boolean (*EQUAL) (obj_Pt,obj_Obj);
};

extern class_Pt the_class_Pt;

struct class_Rect_struct the_class_Rect_struct;

struct class_Rect_struct {
	obj_Rect (*constructor) (obj_Pt,obj_Pt);
	obj_String (*STRING) (obj_Rect);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
	obj_Rect (*translate) (obj_Rect,obj_Pt);
	obj_Boolean (*EQUAL) (obj_Rect,obj_Obj);
};

extern class_Rect the_class_Rect;

struct class_Square_struct the_class_Square_struct;

struct class_Square_struct {
	obj_Square (*constructor) (obj_Int,obj_Pt);
	obj_String (*STRING) (obj_Rect);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
	obj_Rect (*translate) (obj_Rect,obj_Pt);
	obj_Boolean (*EQUAL) (obj_Rect,obj_Obj);
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
obj_String tmp_248 = str_literal("(");
obj_String tmp_249 = this->x->clazz->STRING(this->x);
obj_String tmp_250 = tmp_248->clazz->PLUS(tmp_248, tmp_249);
obj_String tmp_251 = str_literal(",");
obj_String tmp_252 = tmp_250->clazz->PLUS(tmp_250, tmp_251);
obj_String tmp_253 = this->y->clazz->STRING(this->y);
obj_String tmp_254 = tmp_252->clazz->PLUS(tmp_252, tmp_253);
obj_String tmp_255 = str_literal(")");
obj_String tmp_256 = tmp_254->clazz->PLUS(tmp_254, tmp_255);
return tmp_256;
}

obj_Pt Pt_method_PLUS(obj_Pt this, obj_Pt other) {
obj_Int tmp_257 = other->y;
obj_Int tmp_258 = this->y->clazz->PLUS(this->y, tmp_257);
obj_Int tmp_259 = other->x;
obj_Int tmp_260 = this->x->clazz->PLUS(this->x, tmp_259);
obj_Pt tmp_261 = the_class_Pt->constructor(tmp_258,tmp_260);
return tmp_261;
}

obj_Pt Pt_method__x(obj_Pt this) {
return this->x;
}

obj_Pt Pt_method__y(obj_Pt this) {
return this->y;
}

obj_Pt Pt_method_EQUAL(obj_Pt this, obj_Obj other) {
return lit_false;
}

struct class_Pt_struct the_class_Pt_struct = {
new_Pt,
Pt_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
Pt_method_PLUS,
Pt_method__x,
Pt_method__y,
Pt_method_EQUAL,
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
obj_Int tmp_262 = int_literal(1);
obj_Int tmp_263 = int_literal(1);
obj_Pt tmp_264 = the_class_Pt->constructor(tmp_262,tmp_263);
obj_Pt tmp_265 = this->ur->clazz->PLUS(this->ur, tmp_264);
obj_Int tmp_266 = int_literal(1);
obj_Int tmp_267 = int_literal(1);
obj_Pt tmp_268 = the_class_Pt->constructor(tmp_266,tmp_267);
obj_Pt tmp_269 = this->ll->clazz->PLUS(this->ll, tmp_268);
obj_Rect tmp_270 = the_class_Rect->constructor(tmp_265,tmp_269);
return tmp_270;
}

obj_Rect Rect_method_STRING(obj_Rect this) {
obj_Pt lr;
obj_Int tmp_271 = this->ll->clazz->_x(this->ll);
obj_Int tmp_272 = this->ur->clazz->_y(this->ur);
obj_Pt tmp_273 = the_class_Pt->constructor(tmp_271,tmp_272);
lr = tmp_273;
obj_Pt ul;
obj_Int tmp_274 = this->ur->clazz->_y(this->ur);
obj_Int tmp_275 = this->ll->clazz->_x(this->ll);
obj_Pt tmp_276 = the_class_Pt->constructor(tmp_274,tmp_275);
ul = tmp_276;
obj_String tmp_277 = str_literal("(");
obj_String tmp_278 = this->ll->clazz->STRING(this->ll);
obj_String tmp_279 = tmp_277->clazz->PLUS(tmp_277, tmp_278);
obj_String tmp_280 = str_literal(", ");
obj_String tmp_281 = tmp_279->clazz->PLUS(tmp_279, tmp_280);
obj_String tmp_282 = ul->clazz->STRING(ul);
obj_String tmp_283 = tmp_281->clazz->PLUS(tmp_281, tmp_282);
obj_String tmp_284 = str_literal(",");
obj_String tmp_285 = tmp_283->clazz->PLUS(tmp_283, tmp_284);
obj_String tmp_286 = this->ur->clazz->STRING(this->ur);
obj_String tmp_287 = tmp_285->clazz->PLUS(tmp_285, tmp_286);
obj_String tmp_288 = str_literal(",");
obj_String tmp_289 = tmp_287->clazz->PLUS(tmp_287, tmp_288);
obj_String tmp_290 = lr->clazz->STRING(lr);
obj_String tmp_291 = tmp_289->clazz->PLUS(tmp_289, tmp_290);
obj_String tmp_292 = str_literal(")");
obj_String tmp_293 = tmp_291->clazz->PLUS(tmp_291, tmp_292);
return tmp_293;
}

obj_Rect Rect_method_EQUAL(obj_Rect this, obj_Obj other) {
return lit_false;
}

struct class_Rect_struct the_class_Rect_struct = {
new_Rect,
Rect_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
Rect_method_translate,
Rect_method_EQUAL,
};

class_Rect the_class_Rect = &the_class_Rect_struct;

obj_Square new_Square(obj_Int side,obj_Pt ll) {
	obj_Square new_thing = (obj_Square) malloc(sizeof(struct obj_Square_struct));
	new_thing->clazz = the_class_Square;
new_thing->ll = ll;
obj_Int tmp_294 = new_thing->ll->clazz->_y(new_thing->ll);
obj_Int tmp_295 = tmp_294->clazz->PLUS(tmp_294, side);
obj_Int tmp_296 = new_thing->ll->clazz->_x(new_thing->ll);
obj_Int tmp_297 = tmp_296->clazz->PLUS(tmp_296, side);
obj_Pt tmp_298 = the_class_Pt->constructor(tmp_295,tmp_297);
new_thing->ur = tmp_298;
	return new_thing;
}

struct class_Square_struct the_class_Square_struct = {
new_Square,
Rect_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
Rect_method_translate,
Rect_method_EQUAL,
};

class_Square the_class_Square = &the_class_Square_struct;


int main() {
obj_Rect a_square;
obj_Int tmp_299 = int_literal(5);
obj_Int tmp_300 = int_literal(3);
obj_Int tmp_301 = int_literal(3);
obj_Pt tmp_302 = the_class_Pt->constructor(tmp_300,tmp_301);
obj_Square tmp_303 = the_class_Square->constructor(tmp_299,tmp_302);
a_square = (obj_Rect) tmp_303;
obj_Rect b_square;
obj_Int tmp_304 = int_literal(10);
obj_Int tmp_305 = int_literal(10);
obj_Pt tmp_306 = the_class_Pt->constructor(tmp_304,tmp_305);
obj_Int tmp_307 = int_literal(5);
obj_Int tmp_308 = int_literal(5);
obj_Pt tmp_309 = the_class_Pt->constructor(tmp_307,tmp_308);
obj_Rect tmp_310 = the_class_Rect->constructor(tmp_306,tmp_309);
b_square = tmp_310;
obj_Int tmp_311 = int_literal(2);
obj_Int tmp_312 = int_literal(2);
obj_Pt tmp_313 = the_class_Pt->constructor(tmp_311,tmp_312);
obj_Rect tmp_314 = a_square->clazz->translate(a_square,tmp_313);
a_square = tmp_314;
obj_Nothing tmp_315 = a_square->clazz->PRINT(a_square);
tmp_315;
obj_Boolean tmp_316 = a_square->clazz->EQUALS(a_square, b_square);
if (1 == tmp_316->value) {
	goto label_3;
} else {
obj_String tmp_317 = b_square->clazz->STRING(b_square);
obj_String tmp_318 = str_literal(" is different");
obj_String tmp_319 = tmp_317->clazz->PLUS(tmp_317, tmp_318);
obj_Nothing tmp_320 = tmp_319->clazz->PRINT(tmp_319);
	tmp_320;
goto exit_3;
}
label_3: ;
obj_String tmp_321 = str_literal("They are the same");
obj_Nothing tmp_322 = tmp_321->clazz->PRINT(tmp_321);
	tmp_322;
goto exit_3;

exit_3: ;

fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
