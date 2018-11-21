# Quack Parser with AST

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
gocc quack.bnf

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
Testing file 3/77 - Another_plus_types_good.qk
Testing file 4/77 - Comparison_TRUE_FALSE_bad.qk
Testing file 5/77 - CrazyStrings.qk
Testing file 6/77 - GoodWalk.qk
Testing file 7/77 - Inheritance_Types_bad.qk
Testing file 8/77 - LexChallenge.qk
Testing file 9/77 - Plus_types_bad.qk
Testing file 10/77 - Plus_types_good.qk
Testing file 11/77 - Pt.qk
Testing file 12/77 - Pt2.qk
Testing file 13/77 - Pt_missing_fields.qk
Testing file 14/77 - Sqr.qk
Testing file 15/77 - SqrBadLex.qk
Testing file 16/77 - SqrBadSyntax.qk
Testing file 17/77 - SqrDecl.qk
Testing file 18/77 - SqrDeclEQ.qk
Testing file 19/77 - TRUE_FALSE_bad.qk
Testing file 20/77 - TypeWalk.qk
Testing file 21/77 - adv_constructor_init.qk
Testing file 22/77 - bad_break.qk
Testing file 23/77 - bad_class.qk
Testing file 24/77 - bad_class_params.qk
Testing file 25/77 - bad_escape.qk
Testing file 26/77 - bad_init.qk
Testing file 27/77 - bad_typecase_invalid_type.qk
Testing file 28/77 - bad_typecase_recast.qk
Testing file 29/77 - binop_sugar.qk
Testing file 30/77 - circular_dependency.qk
Testing file 31/77 - dot_priority.qk
Testing file 32/77 - duplicate_class.qk
Testing file 33/77 - duplicate_method.qk
Testing file 34/77 - funny_id.qk
Testing file 35/77 - good_typecase.qk
Testing file 36/77 - hands.qk
Testing file 37/77 - if_false_init.qk
Testing file 38/77 - if_true_init.qk
Testing file 39/77 - init_before_use_bad.qk
Testing file 40/77 - init_before_use_good.qk
Testing file 41/77 - invalid_super.qk
Testing file 42/77 - invalid_super_type.qk
Testing file 43/77 - joseph_test_1.qk
Testing file 44/77 - joseph_test_2.qk
Testing file 45/77 - joseph_test_3.qk
Testing file 46/77 - joseph_test_4.qk
Testing file 47/77 - method_madness.qk
Testing file 48/77 - method_madness_2.qk
Testing file 49/77 - method_return_bad.qk
Testing file 50/77 - min_assign.qk
Testing file 51/77 - named_types.qk
Testing file 52/77 - not_a_duck.qk
Testing file 53/77 - robot.qk
Testing file 54/77 - schroedinger.qk
Testing file 55/77 - schroedinger2.qk
Testing file 56/77 - short_test_bad.qk
Testing file 57/77 - short_test_good.qk
Testing file 58/77 - simple_classes_tree_bad_alreadydefined.qk
Testing file 59/77 - simple_classes_tree_bad_circular.qk
Testing file 60/77 - simple_classes_tree_bad_nosuchsuper.qk
Testing file 61/77 - simple_classes_tree_good.qk
Testing file 62/77 - simple_inheritingvariables_bad_notdefined.qk
Testing file 63/77 - simple_inheritingvariables_bad_wrongtype.qk
Testing file 64/77 - simple_inheritingvariables_good.qk
Testing file 65/77 - simple_lhs.qk
Testing file 66/77 - simple_method_return_bad_wrongtype.qk
Testing file 67/77 - simple_method_return_good.qk
Testing file 68/77 - simple_naming_bad_classandmethodsamename.qk
Testing file 69/77 - simple_naming_bad_classandvariablesamename.qk
Testing file 70/77 - simple_overridingmethod_bad_numberargs.qk
Testing file 71/77 - simple_overridingmethod_good.qk
Testing file 72/77 - simple_tree_good_anyorder.qk
Testing file 73/77 - subclass_method_return_mismatch.qk
Testing file 74/77 - tiniest.qk
Testing file 75/77 - typing_test.qk
Testing file 76/77 - unknown_return_type.qk
Testing file 77/77 - while_init.qk
--- PASS: TestFiles (0.03s)
    files_test.go:76: bad_class_params.qk  parse error Error in S137: comma(9,,), Pos(offset=11, line=1, column=12), expected one of: ident rparen
), Pos(offset=103, line=3, column=9), expected one of: semicolon plus minus atleast atmost lt gt neq and or mul div eq period
    files_test.go:98: 0
=== RUN   TestScannerToken
--- PASS: TestScannerToken (0.00s)
=== RUN   TestScannerStrings
--- PASS: TestScannerStrings (0.00s)
=== RUN   TestScannerComments
--- PASS: TestScannerComments (0.00s)
=== RUN   TestScannerTripleQuote
--- PASS: TestScannerTripleQuote (0.00s)
=== RUN   TestScannerEscape
--- PASS: TestScannerEscape (0.00s)
=== RUN   TestEndString
--- PASS: TestEndString (0.00s)
=== RUN   TestAssignment
--- PASS: TestAssignment (0.00s)
=== RUN   TestOperators
--- PASS: TestOperators (0.00s)
=== RUN   TestBoolOperations
--- PASS: TestBoolOperations (0.00s)
=== RUN   TestIfStatement
--- PASS: TestIfStatement (0.00s)
=== RUN   TestWhileStatement
--- PASS: TestWhileStatement (0.00s)
=== RUN   TestClass
--- PASS: TestClass (0.00s)
=== RUN   TestTypecase
--- PASS: TestTypecase (0.00s)
=== RUN   TestIdentOperations
--- PASS: TestIdentOperations (0.00s)
=== RUN   TestClassVars
--- PASS: TestClassVars (0.00s)
=== RUN   TestIdents
--- PASS: TestIdents (0.00s)
=== RUN   TestClassVariableCall
--- PASS: TestClassVariableCall (0.00s)
=== RUN   TestMethodCall
--- PASS: TestMethodCall (0.00s)
PASS
ok      github.com/Lebonesco/quack_parser       0.371s

```
