#include <stdio.h>
#include <stdlib.h>
#include "Builtins.h"


struct class_Printer_struct;
typedef struct class_Printer_struct* class_Printer;

typedef struct obj_Printer_struct {
	class_Printer clazz;
} * obj_Printer;


struct class_Node_struct;
typedef struct class_Node_struct* class_Node;

typedef struct obj_Node_struct {
	class_Node clazz;
	obj_Int car;
	obj_Obj cdr;
} * obj_Node;

struct class_Printer_struct the_class_Printer_struct;

struct class_Printer_struct {
	obj_Printer (*constructor) ();
	obj_String (*STRING) (obj_Obj);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
	obj_Nothing (*println) (obj_Printer,obj_Obj);
	obj_Nothing (*feed) (obj_Printer);
};

extern class_Printer the_class_Printer;

struct class_Node_struct the_class_Node_struct;

struct class_Node_struct {
	obj_Node (*constructor) (obj_Obj,obj_Int);
	obj_String (*STRING) (obj_Node);
	obj_Nothing (*PRINT) (obj_Obj);
	obj_Boolean (*EQUALS) (obj_Obj,obj_Obj);
	obj_Nothing (*setCar) (obj_Node,obj_Int);
	obj_Int (*head) (obj_Node);
	obj_Node (*tail) (obj_Node);
	obj_Int (*len) (obj_Node);
	obj_Nothing (*bubbleSort) (obj_Node);
	obj_Nothing (*swap) (obj_Node,obj_Node,obj_Node);
};

extern class_Node the_class_Node;

obj_Printer new_Printer() {
	obj_Printer new_thing = (obj_Printer) malloc(sizeof(struct obj_Printer_struct));
	new_thing->clazz = the_class_Printer;
	return new_thing;
}

obj_Printer Printer_method_println(obj_Printer this, obj_Obj obj) {
obj_Nothing tmp_387 = obj->clazz->PRINT(obj);
	tmp_387;
obj_Nothing tmp_388 = the_class_Printer->feed(the_class_Printer);
	tmp_388;
}

obj_Printer Printer_method_feed(obj_Printer this) {
obj_String tmp_389 = str_literal("\n");
obj_Nothing tmp_390 = tmp_389->clazz->PRINT(tmp_389);
	tmp_390;
}

struct class_Printer_struct the_class_Printer_struct = {
new_Printer,
Obj_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
Printer_method_println,
Printer_method_feed,
};

class_Printer the_class_Printer = &the_class_Printer_struct;

obj_Node new_Node(obj_Obj cdr,obj_Int car) {
	obj_Node new_thing = (obj_Node) malloc(sizeof(struct obj_Node_struct));
	new_thing->clazz = the_class_Node;
new_thing->car = car;
new_thing->cdr = cdr;
	return new_thing;
}

obj_Node Node_method_STRING(obj_Node this) {
obj_String tmp_391 = this->car->clazz->STRING(this->car);
obj_String tmp_392 = str_literal(",");
obj_String tmp_393 = tmp_391->clazz->PLUS(tmp_391, tmp_392);
obj_String tmp_394 = this->cdr->clazz->STRING(this->cdr);
obj_String tmp_395 = tmp_393->clazz->PLUS(tmp_393, tmp_394);
return tmp_395;
}

obj_Node Node_method_setCar(obj_Node this, obj_Int i) {
this->car = i;
}

obj_Node Node_method_head(obj_Node this) {
return this->car;
}

obj_Node Node_method_tail(obj_Node this) {
return this;
}

obj_Node Node_method_len(obj_Node this) {
obj_Int l;
obj_Int tmp_396 = int_literal(1);
l = tmp_396;
obj_Node temp;
temp = this;
goto label_12;
label_13: ;
obj_Int tmp_397 = int_literal(1);
obj_Int tmp_398 = l->clazz->PLUS(l, tmp_397);
l = tmp_398;
obj_Node tmp_399 = temp->clazz->tail(temp);
temp = tmp_399;
obj_Node tmp_400 = temp->clazz->tail(temp);
obj_Boolean tmp_401 = tmp_400->clazz->EQUALS(tmp_400, temp);
label_12: ;
if (1 == tmp_401->value) {
	goto label_13;
}
return l;
}

obj_Node Node_method_bubbleSort(obj_Node this) {
obj_Node node;
node = this;
obj_Int tmp_402 = the_class_Node->len(the_class_Node);
obj_Int tmp_403 = int_literal(1);
obj_Boolean tmp_404 = tmp_402->clazz->EQUALS(tmp_402, tmp_403);
if (1 == tmp_404->value) {
	goto label_14;
} else {
goto exit_8;
}
label_14: ;
return;
goto exit_8;

exit_8: ;

obj_Int i;
obj_Int tmp_405 = the_class_Node->len(the_class_Node);
i = tmp_405;
goto label_15;
label_16: ;
goto label_17;
label_18: ;
obj_Boolean decision;
decision = lit_false;
if (1 == decision->value) {
	goto label_19;
} else {
goto exit_9;
}
label_19: ;
goto exit_9;

exit_9: ;

obj_Int tmp_406 = node->clazz->len(node);
obj_Int tmp_407 = int_literal(1);
obj_Boolean tmp_408 = tmp_406->clazz->EQUALS(tmp_406, tmp_407);
label_17: ;
if (1 == tmp_408->value) {
	goto label_18;
}
node = this;
obj_Int tmp_409 = int_literal(1);
obj_Int tmp_410 = i->clazz->MINUS(i, tmp_409);
i = tmp_410;
obj_Int tmp_411 = int_literal(0);
obj_Boolean tmp_412 = i->clazz->MORE(i, tmp_411);
label_15: ;
if (1 == tmp_412->value) {
	goto label_16;
}
return;
}

obj_Node Node_method_swap(obj_Node this, obj_Node b, obj_Node a) {
obj_Int temp;
obj_Int tmp_413 = a->car;
temp = tmp_413;
obj_Int tmp_414 = b->car;
obj_Nothing tmp_415 = a->clazz->setCar(a,tmp_414);
	tmp_415;
obj_Nothing tmp_416 = b->clazz->setCar(b,temp);
	tmp_416;
}

struct class_Node_struct the_class_Node_struct = {
new_Node,
Node_method_STRING,
Obj_method_PRINT,
Obj_method_EQUALS,
Node_method_setCar,
Node_method_head,
Node_method_tail,
Node_method_len,
Node_method_bubbleSort,
Node_method_swap,
};

class_Node the_class_Node = &the_class_Node_struct;


int main() {
obj_Node link;
obj_Obj tmp_417 = the_class_Obj->constructor();
obj_Int tmp_418 = int_literal(5535);
obj_Node tmp_419 = the_class_Node->constructor(tmp_417,tmp_418);
obj_Int tmp_420 = int_literal(8573);
obj_Node tmp_421 = the_class_Node->constructor(tmp_419,tmp_420);
obj_Int tmp_422 = int_literal(7088);
obj_Node tmp_423 = the_class_Node->constructor(tmp_421,tmp_422);
obj_Int tmp_424 = int_literal(2998);
obj_Node tmp_425 = the_class_Node->constructor(tmp_423,tmp_424);
obj_Int tmp_426 = int_literal(3538);
obj_Node tmp_427 = the_class_Node->constructor(tmp_425,tmp_426);
obj_Int tmp_428 = int_literal(7808);
obj_Node tmp_429 = the_class_Node->constructor(tmp_427,tmp_428);
obj_Int tmp_430 = int_literal(8922);
obj_Node tmp_431 = the_class_Node->constructor(tmp_429,tmp_430);
obj_Int tmp_432 = int_literal(2628);
obj_Node tmp_433 = the_class_Node->constructor(tmp_431,tmp_432);
obj_Int tmp_434 = int_literal(2571);
obj_Node tmp_435 = the_class_Node->constructor(tmp_433,tmp_434);
obj_Int tmp_436 = int_literal(6968);
obj_Node tmp_437 = the_class_Node->constructor(tmp_435,tmp_436);
obj_Int tmp_438 = int_literal(801);
obj_Node tmp_439 = the_class_Node->constructor(tmp_437,tmp_438);
obj_Int tmp_440 = int_literal(6584);
obj_Node tmp_441 = the_class_Node->constructor(tmp_439,tmp_440);
obj_Int tmp_442 = int_literal(333);
obj_Node tmp_443 = the_class_Node->constructor(tmp_441,tmp_442);
obj_Int tmp_444 = int_literal(2964);
obj_Node tmp_445 = the_class_Node->constructor(tmp_443,tmp_444);
obj_Int tmp_446 = int_literal(2960);
obj_Node tmp_447 = the_class_Node->constructor(tmp_445,tmp_446);
obj_Int tmp_448 = int_literal(4166);
obj_Node tmp_449 = the_class_Node->constructor(tmp_447,tmp_448);
obj_Int tmp_450 = int_literal(2612);
obj_Node tmp_451 = the_class_Node->constructor(tmp_449,tmp_450);
obj_Int tmp_452 = int_literal(817);
obj_Node tmp_453 = the_class_Node->constructor(tmp_451,tmp_452);
obj_Int tmp_454 = int_literal(4844);
obj_Node tmp_455 = the_class_Node->constructor(tmp_453,tmp_454);
obj_Int tmp_456 = int_literal(5445);
obj_Node tmp_457 = the_class_Node->constructor(tmp_455,tmp_456);
obj_Int tmp_458 = int_literal(9070);
obj_Node tmp_459 = the_class_Node->constructor(tmp_457,tmp_458);
obj_Int tmp_460 = int_literal(66);
obj_Node tmp_461 = the_class_Node->constructor(tmp_459,tmp_460);
obj_Int tmp_462 = int_literal(9548);
obj_Node tmp_463 = the_class_Node->constructor(tmp_461,tmp_462);
obj_Int tmp_464 = int_literal(229);
obj_Node tmp_465 = the_class_Node->constructor(tmp_463,tmp_464);
obj_Int tmp_466 = int_literal(5664);
obj_Node tmp_467 = the_class_Node->constructor(tmp_465,tmp_466);
obj_Int tmp_468 = int_literal(3086);
obj_Node tmp_469 = the_class_Node->constructor(tmp_467,tmp_468);
obj_Int tmp_470 = int_literal(7585);
obj_Node tmp_471 = the_class_Node->constructor(tmp_469,tmp_470);
obj_Int tmp_472 = int_literal(9652);
obj_Node tmp_473 = the_class_Node->constructor(tmp_471,tmp_472);
obj_Int tmp_474 = int_literal(6382);
obj_Node tmp_475 = the_class_Node->constructor(tmp_473,tmp_474);
obj_Int tmp_476 = int_literal(7431);
obj_Node tmp_477 = the_class_Node->constructor(tmp_475,tmp_476);
obj_Int tmp_478 = int_literal(8223);
obj_Node tmp_479 = the_class_Node->constructor(tmp_477,tmp_478);
obj_Int tmp_480 = int_literal(6740);
obj_Node tmp_481 = the_class_Node->constructor(tmp_479,tmp_480);
obj_Int tmp_482 = int_literal(4448);
obj_Node tmp_483 = the_class_Node->constructor(tmp_481,tmp_482);
obj_Int tmp_484 = int_literal(3595);
obj_Node tmp_485 = the_class_Node->constructor(tmp_483,tmp_484);
obj_Int tmp_486 = int_literal(8371);
obj_Node tmp_487 = the_class_Node->constructor(tmp_485,tmp_486);
obj_Int tmp_488 = int_literal(4572);
obj_Node tmp_489 = the_class_Node->constructor(tmp_487,tmp_488);
obj_Int tmp_490 = int_literal(6107);
obj_Node tmp_491 = the_class_Node->constructor(tmp_489,tmp_490);
obj_Int tmp_492 = int_literal(589);
obj_Node tmp_493 = the_class_Node->constructor(tmp_491,tmp_492);
obj_Int tmp_494 = int_literal(4139);
obj_Node tmp_495 = the_class_Node->constructor(tmp_493,tmp_494);
obj_Int tmp_496 = int_literal(9181);
obj_Node tmp_497 = the_class_Node->constructor(tmp_495,tmp_496);
obj_Int tmp_498 = int_literal(9329);
obj_Node tmp_499 = the_class_Node->constructor(tmp_497,tmp_498);
obj_Int tmp_500 = int_literal(5481);
obj_Node tmp_501 = the_class_Node->constructor(tmp_499,tmp_500);
obj_Int tmp_502 = int_literal(2196);
obj_Node tmp_503 = the_class_Node->constructor(tmp_501,tmp_502);
obj_Int tmp_504 = int_literal(2861);
obj_Node tmp_505 = the_class_Node->constructor(tmp_503,tmp_504);
obj_Int tmp_506 = int_literal(5746);
obj_Node tmp_507 = the_class_Node->constructor(tmp_505,tmp_506);
obj_Int tmp_508 = int_literal(5168);
obj_Node tmp_509 = the_class_Node->constructor(tmp_507,tmp_508);
obj_Int tmp_510 = int_literal(7210);
obj_Node tmp_511 = the_class_Node->constructor(tmp_509,tmp_510);
obj_Int tmp_512 = int_literal(5501);
obj_Node tmp_513 = the_class_Node->constructor(tmp_511,tmp_512);
obj_Int tmp_514 = int_literal(4040);
obj_Node tmp_515 = the_class_Node->constructor(tmp_513,tmp_514);
obj_Int tmp_516 = int_literal(6888);
obj_Node tmp_517 = the_class_Node->constructor(tmp_515,tmp_516);
obj_Int tmp_518 = int_literal(7784);
obj_Node tmp_519 = the_class_Node->constructor(tmp_517,tmp_518);
obj_Int tmp_520 = int_literal(126);
obj_Node tmp_521 = the_class_Node->constructor(tmp_519,tmp_520);
obj_Int tmp_522 = int_literal(8835);
obj_Node tmp_523 = the_class_Node->constructor(tmp_521,tmp_522);
obj_Int tmp_524 = int_literal(5580);
obj_Node tmp_525 = the_class_Node->constructor(tmp_523,tmp_524);
obj_Int tmp_526 = int_literal(3114);
obj_Node tmp_527 = the_class_Node->constructor(tmp_525,tmp_526);
obj_Int tmp_528 = int_literal(6510);
obj_Node tmp_529 = the_class_Node->constructor(tmp_527,tmp_528);
obj_Int tmp_530 = int_literal(2575);
obj_Node tmp_531 = the_class_Node->constructor(tmp_529,tmp_530);
obj_Int tmp_532 = int_literal(3860);
obj_Node tmp_533 = the_class_Node->constructor(tmp_531,tmp_532);
obj_Int tmp_534 = int_literal(6717);
obj_Node tmp_535 = the_class_Node->constructor(tmp_533,tmp_534);
obj_Int tmp_536 = int_literal(6706);
obj_Node tmp_537 = the_class_Node->constructor(tmp_535,tmp_536);
obj_Int tmp_538 = int_literal(3872);
obj_Node tmp_539 = the_class_Node->constructor(tmp_537,tmp_538);
obj_Int tmp_540 = int_literal(2865);
obj_Node tmp_541 = the_class_Node->constructor(tmp_539,tmp_540);
obj_Int tmp_542 = int_literal(9889);
obj_Node tmp_543 = the_class_Node->constructor(tmp_541,tmp_542);
obj_Int tmp_544 = int_literal(2987);
obj_Node tmp_545 = the_class_Node->constructor(tmp_543,tmp_544);
obj_Int tmp_546 = int_literal(6843);
obj_Node tmp_547 = the_class_Node->constructor(tmp_545,tmp_546);
obj_Int tmp_548 = int_literal(8696);
obj_Node tmp_549 = the_class_Node->constructor(tmp_547,tmp_548);
obj_Int tmp_550 = int_literal(5657);
obj_Node tmp_551 = the_class_Node->constructor(tmp_549,tmp_550);
obj_Int tmp_552 = int_literal(3827);
obj_Node tmp_553 = the_class_Node->constructor(tmp_551,tmp_552);
obj_Int tmp_554 = int_literal(835);
obj_Node tmp_555 = the_class_Node->constructor(tmp_553,tmp_554);
obj_Int tmp_556 = int_literal(7580);
obj_Node tmp_557 = the_class_Node->constructor(tmp_555,tmp_556);
obj_Int tmp_558 = int_literal(1616);
obj_Node tmp_559 = the_class_Node->constructor(tmp_557,tmp_558);
obj_Int tmp_560 = int_literal(8066);
obj_Node tmp_561 = the_class_Node->constructor(tmp_559,tmp_560);
obj_Int tmp_562 = int_literal(1179);
obj_Node tmp_563 = the_class_Node->constructor(tmp_561,tmp_562);
obj_Int tmp_564 = int_literal(11);
obj_Node tmp_565 = the_class_Node->constructor(tmp_563,tmp_564);
obj_Int tmp_566 = int_literal(7101);
obj_Node tmp_567 = the_class_Node->constructor(tmp_565,tmp_566);
obj_Int tmp_568 = int_literal(3724);
obj_Node tmp_569 = the_class_Node->constructor(tmp_567,tmp_568);
obj_Int tmp_570 = int_literal(9327);
obj_Node tmp_571 = the_class_Node->constructor(tmp_569,tmp_570);
obj_Int tmp_572 = int_literal(6452);
obj_Node tmp_573 = the_class_Node->constructor(tmp_571,tmp_572);
obj_Int tmp_574 = int_literal(8817);
obj_Node tmp_575 = the_class_Node->constructor(tmp_573,tmp_574);
obj_Int tmp_576 = int_literal(8012);
obj_Node tmp_577 = the_class_Node->constructor(tmp_575,tmp_576);
obj_Int tmp_578 = int_literal(6832);
obj_Node tmp_579 = the_class_Node->constructor(tmp_577,tmp_578);
obj_Int tmp_580 = int_literal(2863);
obj_Node tmp_581 = the_class_Node->constructor(tmp_579,tmp_580);
obj_Int tmp_582 = int_literal(75);
obj_Node tmp_583 = the_class_Node->constructor(tmp_581,tmp_582);
obj_Int tmp_584 = int_literal(2229);
obj_Node tmp_585 = the_class_Node->constructor(tmp_583,tmp_584);
obj_Int tmp_586 = int_literal(209);
obj_Node tmp_587 = the_class_Node->constructor(tmp_585,tmp_586);
obj_Int tmp_588 = int_literal(4975);
obj_Node tmp_589 = the_class_Node->constructor(tmp_587,tmp_588);
obj_Int tmp_590 = int_literal(5469);
obj_Node tmp_591 = the_class_Node->constructor(tmp_589,tmp_590);
obj_Int tmp_592 = int_literal(2610);
obj_Node tmp_593 = the_class_Node->constructor(tmp_591,tmp_592);
obj_Int tmp_594 = int_literal(5565);
obj_Node tmp_595 = the_class_Node->constructor(tmp_593,tmp_594);
obj_Int tmp_596 = int_literal(2955);
obj_Node tmp_597 = the_class_Node->constructor(tmp_595,tmp_596);
obj_Int tmp_598 = int_literal(2875);
obj_Node tmp_599 = the_class_Node->constructor(tmp_597,tmp_598);
obj_Int tmp_600 = int_literal(942);
obj_Node tmp_601 = the_class_Node->constructor(tmp_599,tmp_600);
obj_Int tmp_602 = int_literal(7133);
obj_Node tmp_603 = the_class_Node->constructor(tmp_601,tmp_602);
obj_Int tmp_604 = int_literal(3500);
obj_Node tmp_605 = the_class_Node->constructor(tmp_603,tmp_604);
obj_Int tmp_606 = int_literal(3184);
obj_Node tmp_607 = the_class_Node->constructor(tmp_605,tmp_606);
obj_Int tmp_608 = int_literal(1011);
obj_Node tmp_609 = the_class_Node->constructor(tmp_607,tmp_608);
obj_Int tmp_610 = int_literal(8802);
obj_Node tmp_611 = the_class_Node->constructor(tmp_609,tmp_610);
obj_Int tmp_612 = int_literal(9829);
obj_Node tmp_613 = the_class_Node->constructor(tmp_611,tmp_612);
obj_Int tmp_614 = int_literal(847);
obj_Node tmp_615 = the_class_Node->constructor(tmp_613,tmp_614);
obj_Int tmp_616 = int_literal(6342);
obj_Node tmp_617 = the_class_Node->constructor(tmp_615,tmp_616);
link = tmp_617;
obj_Printer p;
obj_Printer tmp_618 = the_class_Printer->constructor();
p = tmp_618;
obj_Nothing tmp_619 = p->clazz->feed(p);
tmp_619;
obj_Nothing tmp_620 = p->clazz->println(p,link);
tmp_620;
obj_Nothing tmp_621 = p->clazz->feed(p);
tmp_621;
obj_String tmp_622 = str_literal("Sorting...");
obj_Nothing tmp_623 = p->clazz->println(p,tmp_622);
tmp_623;
obj_Nothing tmp_624 = link->clazz->bubbleSort(link);
tmp_624;
obj_Nothing tmp_625 = p->clazz->feed(p);
tmp_625;
obj_Nothing tmp_626 = p->clazz->println(p,link);
tmp_626;
obj_Nothing tmp_627 = p->clazz->feed(p);
tmp_627;
fprintf(stdout, "\n--- Terminated SuccessFully (woot!) ---");
	return 0;
}
