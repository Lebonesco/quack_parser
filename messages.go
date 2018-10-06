package main 

import (
	"fmt"
	"log"
	"github.com/Lebonesco/quack_scanner/token"
)

var errorCount = 0
const errorLimit = 5

func Bail() {
	log.Fatalln("Too many errors, bailing")
}

func ErrorAt(tok *token.Token, msg string) {
	fmt.Printf("%s - '%s' at line %d, column %d \n", msg, tok.Lit, tok.Pos.Line, tok.Pos.Column)
	errorCount++
	if errorCount > errorLimit {
		Bail()
	}
}

/* Additional diagnostic message, does not count against error limit */
func Note(msg string) {
	// print
}

// any errors?
func Ok() bool {
	return errorCount == 0
}