package main

import (
	"github.com/Lebonesco/quack_parser/lexer"
	"github.com/Lebonesco/quack_parser/parser"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"log"
)


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
		fmt.Println(err.Error())
	}

	fmt.Println(res)
	fmt.Println("parser is done...")
}