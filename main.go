package main

import (
	"fmt"
	"os"
	"panda/eval"
	"panda/lexer"
	"panda/parse"
	"panda/repl"
	"strings"
)

var debug = true

func main() {
	if debug {
		lex := lexer.New(strings.NewReader(`
		var a = 1 +2*3;
		a;
		a = 12;
		`))
		p := parse.New(lex)
		inter := eval.New(p)
		fmt.Fprintf(os.Stdout, " %v\n", inter.Eval())
	}
	repl.StartREPL(os.Stdin, os.Stdout)
}
