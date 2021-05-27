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
		a = a+12;
		a;
		var b = 1+a;
		b*a+1+2;
		var calla = function(a, b, c, d) {
			return a+b+c+d;
		};
		var callb = function() {
			return 1+2+3;
		};
		calla(1,1,1,1);
		`))
		p := parse.New(lex)
		inter := eval.New(p)
		fmt.Fprintf(os.Stdout, " %v\n", inter.Eval())
	}
	repl.StartREPL(os.Stdin, os.Stdout)
}

/*

 */
