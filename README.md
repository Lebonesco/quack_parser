# Quack Parser with AST

## Directory
* [How to Run Program (with Make)](#how-to-run)
* [How to Run Program (Manual)](#how-to-run-manually)
* [Missing Functionality](#missing-functionality)
* [Test Outputs](#tests)
    * [Bad](#bad)
    * [Good](#good)

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
<a href="good">
### Good (Programs that don't have errors)
</a>

