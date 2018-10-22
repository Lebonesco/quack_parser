package main

import (
	"testing"
	"github.com/Lebonesco/quack_parser/lexer"
	"github.com/Lebonesco/quack_parser/parser"
	"github.com/Lebonesco/quack_parser/ast"
	"io/ioutil"
)

const DIR = "./samples"

func TestFiles(t *testing.T) {
	files, err := ioutil.ReadDir(DIR)
    if err != nil {
        panic(err)
    }

	for i, file := range files {
		t.Log(i, file.Name())
		data, err := ioutil.ReadFile(DIR+"/"+file.Name())
		if err != nil {
			t.Errorf(err.Error())
			continue
		}

		l := lexer.NewLexer([]byte(data))
		p := parser.NewParser()
		res, err := p.Parse(l)
		if err != nil {
			t.Fatalf(err.Error())
		}

		program, _ := res.(*ast.Program)
		_ = program.Statements[0]
	}
}