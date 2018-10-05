package main

import (
	"github.com/Lebonesco/quack_scanner/lexer"
	"github.com/Lebonesco/quack_scanner/token"
	"fmt"
)

func main() {
	fmt.Println("starting scanner...")
	l := lexer.NewLexer([]byte(`let five = "te
	st";

	"test \n this"`))

	tok := l.Scan()
	for tok.Type != token.TokMap.Type("$") { // keep scanning unitl end
		fmt.Printf("{Type: %s, Literal: '%s'}\n", token.TokMap.Id(tok.Type), tok.Lit)
		tok = l.Scan()
	}	
	fmt.Println("scanner is done...")
}
