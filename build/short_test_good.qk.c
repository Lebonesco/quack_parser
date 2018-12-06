#include <stdio.h>
#include <stdlib.h>
#include "Builtins.h"


int main() {
obj_Int a;
obj_Int tmp_358 = int_literal(5);
a = tmp_358;
obj_Int b;
obj_Int tmp_359 = int_literal(2);
b = tmp_359;
obj_Int x;
obj_Int tmp_360 = a->clazz->PLUS(a, b);
x = tmp_360;
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
