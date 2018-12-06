# Quack Parser with AST

## Directory
* [How to Run Program (with Make)](#how-to-run)
* [How to Run Program (Manual)](#how-to-run-manually)
* [Missing Functionality](#missing-functionality)
* [Test Outputs](#tests)
    * [Bad](#bad-programs-that-produce-errors)
    * [Good](#good-programs-that-do-not-produce-errors)

## How to Run

### Download the code

First make sure that you have Golang installed on your system
and that you have your `$GOPATH` set. It is normally
defaulted to `$HOME/go`

Go is really particular when it comes
to import references having absolute paths.
That's why when you:
`git clone https://github.com/Lebonesco/quack_parser.git`
make sure that it is in the file path:
`$GOPATH/src/github.com/Lebonesco/`

At this point, everything can be run from the `Makefile`

### Build, Test, Run Program
```make file=test.txt```

### Install Dependencies
```make deps```

### Build Program
```make build```

### Run Unit Tests
```make test```

### Run Program
```make run file=text.txt```

### Clean Directory
```make clean```

## How to Run Manually

### Install the Parser Tool (gocc)
```git clone https://github.com/goccmack/gocc.git```

It will download to `$GOPATH/go/bin`
Make sure that this directory is in your `Path`

### Generate the Lexer and Parser
```bash
gocc quack.bnf
```

### Run Program 
To run program with test file/file path:
`go run main.go text.txt`

### Run Unit Tests
To run unit tests:
```
$ go test -v
=== RUN   TestFiles
Testing file 1/77 - Another_plus_good.qk
Testing file 2/77 - Another_plus_types_bad.qk
...
Testing file 77/77 - while_init.qk
--- PASS: TestFiles (0.03s)
--- PASS: TestClassVariableCall (0.00s)
=== RUN   TestMethodCall
--- PASS: TestMethodCall (0.00s)
PASS
ok      github.com/Lebonesco/quack_parser       0.371s

```

## Missing Functionality

* lacks code generation for typecase statements
* unable to generate correct multiline, triple quote strings
* havent fully completed code generation for prefix values

## Tests

### Bad (Programs that produce errors)

```
make run file=./samples/circular_dependency.qk
Type Error: CLASS_CYCLE
```

```
make run file=./samples/Comparison_TRUE_FALSE_bad.qk
Type Error: INCOMPATABLE_TYPES - types Int-x and Boolean- not work for expression '==' on line 8
```

```
make run file=./samples/dot_priority.qk
Type Error: VARIABLE_NOT_INITIALIZED - ident i is not defined on line: 10
```

```
make run file=./samples/duplicate_class.qk
Type Error: DUPLICATE_CLASS - class C1 already exists
```

```
make run file=./samples/duplicate_method.qk
Type Error: ALREADY_INITIALIZED - method name: x already exists in class
```

```
make run file=./samples/GoodWalk.qk
Type Error: type 'Top' in method return signature 'foo' not exist
```

```
make run file=./samples/TypeWalk.qk
Type Error: METHOD_NOT_EXIST - method foo not exist in class Obj
```

```
make run file=./samples/typing_test.qk
Type Error: CLASS_NOT_EXIST - class 'NotExist' doesn't exist
```

```
make run file=./samples/unknown_return_type.qk
Type Error: CLASS_NOT_EXIST - type 'Garbage' in method return signature 'PLUS' not exist
```

```
make run file=./samples/subclass_method_return_mismatch.qk
Type Error: INVALID_RETURN_TYPE - overriding method 'x' in C2 has incompatible return type 'Obj'. parent C1 has type 'Int'
```

```
make run file=./samples/simple_naming_bad_classmethodsamename.qk
Type Error: ALREADY_INITIALIZED - method R can't be the same as the class
```

```
make run file=./samples/simple_inheritingvariables_bad_notdefined.qk
Type Error: CREATE_CLASS_FAIL - variables in C2 incompatible with C1
```

```
make run file=./samples/hands.qk
Type Error: METHOD_NOT_EXIST - method foo not exist in class Hand
```

```
make run file=./samples/Inheritance_Types_bad.qk
Type Error: CREATE_CLASS_FAIL - variables in Blend incompatible with Filter
```

```
make run file=./samples/duplicate_class.qk
Type Error: DUPLICATE_CLASS - class C1 already exists
```

```
make run file=./samples/simple_overridingmethod_bad_numberargs.qk
Type Error: INCORRECT_ARGUMENT_COUNT - child overriding method have incorrect param length 1 vs 0
```

```
make run files=./sample/simple_classes_tree_bad_nosuchsuper.qk
Type Error: CLASS_NOT_EXIST - class not exist
```

```
make run file=./samples/prefix.qk
Gen Error: ./build/main.c:10:11: error: ‘none’ undeclared (first use in this function)
```

```
make run file=./samples/schroedinger.qk
Type Error: VARIABLE_NOT_INITIALIZED - variable this.living not initialized on all paths
```

```
make run file=./samples/Plus_types_bad.qk
Type Error: INCOMPATABLE_TYPES - incorrect argument type Int for Class Add, expected String on line 19
```

```
make run file=./samples/joseph_test_1.qk
Type Error: INVALID_RETURN_TYPE - incorrect return subtype Int in method testMethod, wanted Boolean
```

```
make run file=./samples/joseph_test_6.qk
Type Error: METHOD_NOT_EXIST - method PLUS not exist in class Obj on line 10
```

```
make run file=./samples/bad_class_params.qk
Parse Error:  ',' type: 'comma(9)' in line 1, expected one of: 'ident' 'rparen'
```

```
make run file=./samples/bad_escape.qk
Parse Error: ' type: 'INVALID(0)' in line 3, expected one of:  this quote;
'semicolon' 'plus' 'minus' 'atleast' 'atmost' 'lt' 'gt' 'neq' 'and' 'or' 'mul' 'div' 'eq' 'period'
```

### Good (Programs that do not produce errors)

```
make run file=./samples/funny_id.qk
--- Terminated SuccessFully (woot!) ---
```

```
make run file=./samples/good_typcase.qk
--- Terminated SuccessFully (woot!) ---
```

```
make run file=./samples/Sqr.qk
((4,4), (4,9),(9,9),(9,4))
HELLO
--- Terminated SuccessFully (woot!) ---
```

```
make run file=./samples/SqrDeclEQ.qk
((4,4), (4,9),(9,9),(9,4))((5,5), (5,10),(10,10),(10,5)) is different
--- Terminated SuccessFully (woot!) ---
```

```
make run file=./samples/sort.qk
6342,847,9829,8802,1011,3184,3500,7133,942,2875,2955,5565,2610,5469,4975,209,2229,75,2863,6832,8012,8817,6452,9327,3724,7101,11,1179,8066,1616,7580,835,3827,5657,8696,6843,2987,9889,2865,3872,6706,6717,3860,2575,6510,3114,5580,8835,126,7784,6888,4040,5501,7210,5168,5746,2861,2196,5481,9329,9181,4139,589,6107,4572,8371,3595,4448,6740,8223,7431,6382,9652,7585,3086,5664,229,9548,66,9070,5445,4844,817,2612,4166,2960,2964,333,6584,801,6968,2571,2628,8922,7808,3538,2998,7088,8573,5535,<Object at 25769804896>#########

Sorting...

6342,847,9829,8802,1011,3184,3500,7133,942,2875,2955,5565,2610,5469,4975,209,2229,75,2863,6832,8012,8817,6452,9327,3724,7101,11,1179,8066,1616,7580,835,3827,5657,8696,6843,2987,9889,2865,3872,6706,6717,3860,2575,6510,3114,5580,8835,126,7784,6888,4040,5501,7210,5168,5746,2861,2196,5481,9329,9181,4139,589,6107,4572,8371,3595,4448,6740,8223,7431,6382,9652,7585,3086,5664,229,9548,66,9070,5445,4844,817,2612,4166,2960,2964,333,6584,801,6968,2571,2628,8922,7808,3538,2998,7088,8573,5535,<Object at 25769804896>#########


--- Terminated SuccessFully (woot!) ---
```

```
make run file=./samples/while_init.qk
--- Terminated SuccessFully (woot!) ---
```

```
make run file=./samples/simple_tree_good_anyorder.qk
--- Terminated SuccessFully (woot!) ---
```

```
make run file=./samples/Pt.qk
--- Terminated SuccessFully (woot!) ---
```

```
make run file=./samples/Pt2.qk
<Object at 25769804960>
--- Terminated SuccessFully (woot!) ---
```

```
make run file=./samples/robot.qk
Type Error: CREATE_CLASS_FAIL - variables in SmartRobot incompatible with Robot
```

```
make run file=./samples/Plus_types_good.qk
<Object at 25769804960>
--- Terminated SuccessFully (woot!) ---
```

```
make run file=./samples/joseph_test_5.qk
2
--- Terminated SuccessFully (woot!) ---
```

```
make run file=./samples/joseph_test_2.qk
--- Terminated SuccessFully (woot!) ---
```

```
make run file=./samples/tiniest.qk
11
19
20

--- Terminated SuccessFully (woot!) ---
```