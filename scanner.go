package main

import (
	"github.com/Lebonesco/quack_scanner/lexer"
	"github.com/Lebonesco/quack_scanner/token"
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

	tok := l.Scan()
	for tok.Type != token.TokMap.Type("$") { // keep scanning unitl end
		switch tok.Type {
		case token.TokMap.Type("string_escape_error"):
			ErrorAt(tok, "ILLEGAL use of escape character in string")
		case token.TokMap.Type("unknown"):
			ErrorAt(tok, "Unexpected item")
		default:
			fmt.Printf("{Type: %s, Literal: '%s'}\n", token.TokMap.Id(tok.Type), tok.Lit)
		}
		tok = l.Scan()
	}	
	fmt.Println("scanner is done...")
}