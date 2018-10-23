package main

import (
	"testing"
	"github.com/Lebonesco/quack_parser/lexer"
	"github.com/Lebonesco/quack_parser/parser"
	"github.com/Lebonesco/quack_parser/ast"
	"io/ioutil"
	"fmt"
)

const DIR = "./samples"

func TestFiles(t *testing.T) {
	files, err := ioutil.ReadDir(DIR)
    if err != nil {
        panic(err)
    }

	for i, file := range files {
		fmt.Printf("Testing file %d/%d - %s\n", i, len(files), file.Name())
		data, err := ioutil.ReadFile(DIR+"/"+file.Name())
		if err != nil {
			t.Fatalf(err.Error())
			continue
		}

		l := lexer.NewLexer([]byte(data))
		p := parser.NewParser()
		res, err := p.Parse(l)
		if err != nil {
			t.Fatalf(err.Error())
		}

		_, _ = res.(*ast.Program)
	}
}