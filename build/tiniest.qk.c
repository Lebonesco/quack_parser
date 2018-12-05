#include <stdio.h>
#include <stdlib.h>
#include "Builtins.h"


int main() {
obj_Int tmp_367 = int_literal(42);
tmp_367;
obj_Int x;
obj_Int tmp_368 = int_literal(5);
obj_Int tmp_369 = int_literal(6);
obj_Int tmp_370 = tmp_368->clazz->PLUS(tmp_368, tmp_369);
x = tmp_370;
obj_Nothing tmp_371 = x->clazz->PRINT(x);
tmp_371;
obj_String tmp_372 = str_literal("\n");
obj_Nothing tmp_373 = tmp_372->clazz->PRINT(tmp_372);
tmp_373;
obj_Int y;
obj_Int tmp_374 = int_literal(8);
obj_Int tmp_375 = x->clazz->PLUS(x, tmp_374);
y = tmp_375;
obj_Nothing tmp_376 = y->clazz->PRINT(y);
tmp_376;
obj_String tmp_377 = str_literal("\n");
obj_Nothing tmp_378 = tmp_377->clazz->PRINT(tmp_377);
tmp_378;
obj_Int tmp_379 = int_literal(1);
obj_Int tmp_380 = y->clazz->PLUS(y, tmp_379);
x = tmp_380;
obj_Nothing tmp_381 = x->clazz->PRINT(x);
tmp_381;
obj_String tmp_382 = str_literal("\n");
obj_Nothing tmp_383 = tmp_382->clazz->PRINT(tmp_382);
tmp_383;
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
