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
$ go  test -v
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
    parser_test.go:443: {Pos(offset=0, line=0, column=0) 0xc0000b0d40 0xc00005afc0}
PASS
ok      github.com/Lebonesco/quack_parser       0.373s
```
