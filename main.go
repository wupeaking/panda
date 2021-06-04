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
		var a = [1, 2, [3, 1, 2]];
		println(a[2][1]);
		`))
		p := parse.New(lex)
		inter := eval.New(p)
		v, err := inter.Eval()
		fmt.Fprintf(os.Stdout, " %v %v\n", v, err)
	}
	repl.StartREPL(os.Stdin, os.Stdout)
}

/*
	var a = 1 +2*3;
	a;
	a = a+12.1;
	var b = 1+a;
	b*a+1+2;
	var calla = function(a, b, c, d) {
		return a+b+c+d;
	};
	var callb = function() {
		return 1+2+3;
	};
	var callc = function() {
		return function(a, b,c) {
			return a*b*c;
		};
	};
	calla(1,1,1,1);
	callb();
	callc()(2,2,3);
	a+callc()(2,2,3);
	function add(left, right) {
		return left+right;
	}
	add(1, 2.1*3);
	var b = 1;
	for(var a = 1; a<10; a=a+1){
		b = b +a;
		b;
		if (a>=5){
			break;
		}
	}
*/
