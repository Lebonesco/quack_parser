package main

import (
	"bytes"
	"fmt"
	"github.com/Lebonesco/quack_parser/errors"
	"github.com/Lebonesco/quack_parser/lexer"
	"github.com/Lebonesco/quack_parser/parser"
	"github.com/Lebonesco/quack_parser/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func format(e *errors.Error) string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "Error: value: '%s' type: '%s' in line %d, expected one of: \n", e.ErrorToken.Lit, token.TokMap.StringType(e.ErrorToken.Type), e.ErrorToken.Pos.Line)
	for _, sym := range e.ExpectedTokens {
		fmt.Fprintf(w, "'%s' ", sym)
	}
	return w.String()
}

func main() {
	fmt.Println("starting scanner...")
	if len(os.Args) < 2 {
		log.Fatalln("no valid file name or path provided provided!")
	}

	path := os.Args[1]
	absPath, _ := filepath.Abs(path)
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		panic(err)
	}

	l := lexer.NewLexer([]byte(data))
	p := parser.NewParser()
	res, err := p.Parse(l)
	if err != nil {
		fmt.Println("Oh no, there were errors!")
		fmt.Println(format(err.(*errors.Error)))
		return
	}

	fmt.Println("AST has successfully been constructed")

	_ = res
	fmt.Println("compiler is done...")
}
