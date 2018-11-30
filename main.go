package main

import (
	"bytes"
	"fmt"
	"github.com/Lebonesco/quack_parser/errors"
	"github.com/Lebonesco/quack_parser/lexer"
	"github.com/Lebonesco/quack_parser/parser"
	"github.com/Lebonesco/quack_parser/token"
	"github.com/Lebonesco/quack_parser/typechecker"
	"github.com/Lebonesco/quack_parser/environment"
	"github.com/Lebonesco/quack_parser/ast"
	//"github.com/Lebonesco/quack_parser/codegen"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func printGood(messsage string) {
	// print good messages
}

func printBad(message string) {
	// print bad messages
}

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
		log.Fatalln("no valid file name or path provided provided for file!")
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
		fmt.Println("Oh no, there were parsing errors!")
		fmt.Println(format(err.(*errors.Error)))
		return
	}

	fmt.Println("AST has successfully been constructed")

	program, _ := res.(*ast.Program)
	env := environment.CreateEnvironment() // create new environment
	_ , typeErr := typechecker.TypeCheck(program, env)
	if typeErr != nil {
		fmt.Println("checking errors")
		fmt.Printf(string(typeErr.Type) + " - " + typeErr.Message.Error())
	}
	fmt.Println("Checking is completed")
	fmt.Println("Starting code compilation...")

	//fileName := "main"

	// code, err := CodeGen(program) 
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	//write code to file

	fmt.Println("compiler is done...")
}
