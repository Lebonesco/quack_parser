#include <stdio.h>
#include <stdlib.h>
#include "Builtins.h"


int main() {
obj_Int tmp_17 = int_literal(42);
tmp_17;
obj_String bog_standard;
obj_String tmp_18 = str_literal("This is a bog standard string");
bog_standard = tmp_18;
obj_String escaped;
obj_String tmp_19 = str_literal("This\tstring\nhas escapes");
escaped = tmp_19;
obj_String triple;
obj_String tmp_20 = str_literal("""This string has a newline""");
triple = tmp_20;
obj_String raw;
obj_String tmp_21 = str_literal("""This\tstring\tdoes\nnot\yhave\nescapes""");
raw = tmp_21;
obj_String tmp_22 = str_literal("This is a bog standard string");
bog_standard = tmp_22;
obj_String tmp_23 = str_literal("This\tstring\nhas escapes");
escaped = tmp_23;
obj_String tmp_24 = str_literal("""This string has a newline""");
triple = tmp_24;
obj_String tmp_25 = str_literal("""This\tstring\tdoes\nnot\yhave\nescapes""");
raw = tmp_25;
obj_String x;
obj_String tmp_26 = str_literal("oh no, comments must be closed.");
x = tmp_26;
obj_String s;
obj_String tmp_27 = str_literal("This cquote runs off the end of the file");
s = tmp_27;
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
