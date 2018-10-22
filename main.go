package main

import (
	"github.com/Lebonesco/quack_parser/lexer"
	"github.com/Lebonesco/quack_parser/parser"
	"github.com/Lebonesco/quack_parser/errors"
	"fmt"
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"log"
)

func format(e *errors.Error) string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "Value: '%s' in line %d\n", e.ErrorToken.Lit, e.ErrorToken.Pos.Line)
	fmt.Fprintf(w, "Expected one of: ")
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
	} else {
		fmt.Println("Yay, there were no errors!")
	}

	_ = res
	fmt.Println("parser is done...")
}