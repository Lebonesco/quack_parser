#include <stdio.h>
#include <stdlib.h>
#include "Builtins.h"


int main() {
obj_Int _;
obj_Int tmp_327 = int_literal(42);
_ = tmp_327;
obj_Int __;
__ = _;
obj_Int ____this_is_a_really_bad__id____but__its__legal;
____this_is_a_really_bad__id____but__its__legal = __;
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
