package main

import (
	"fmt"
	"github.com/Lebonesco/quack_parser/ast"
	"github.com/Lebonesco/quack_parser/lexer"
	"github.com/Lebonesco/quack_parser/parser"
	"github.com/Lebonesco/quack_parser/typechecker"
	"io/ioutil"
	"testing"
)

const DIR = "./samples"

func TestFiles(t *testing.T) {
	files, err := ioutil.ReadDir(DIR)
	if err != nil {
		panic(err)
	}

	for i, file := range files {
		fmt.Printf("Testing file %d/%d - %s\n", i+1, len(files), file.Name())
		data, err := ioutil.ReadFile(DIR + "/" + file.Name())
		if err != nil {
			t.Fatalf(err.Error())
			continue
		}

		l := lexer.NewLexer([]byte(data))
		p := parser.NewParser()
		res, err := p.Parse(l)
		if err != nil {
			t.Log(err.Error())
			//t.Log("\n------------------------------------------------------------------")
			t.Log(file.Name(), "error...")
		}

		program, ok := res.(*ast.Program)
		if !ok {
			//t.Fatalf(err.Error())
		}

		//t.Log("Passed grammatical check")
		env := typechecker.CreateEnvironment()
		_, err = typechecker.TypeCheck(program, env)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
}
