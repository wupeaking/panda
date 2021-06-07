package main

import (
	"fmt"
	"os"
	"panda/eval"
	"panda/lexer"
	"panda/parse"
	"panda/repl"
)

var debug = true

func main() {
	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer file.Close()
		lex := lexer.New(file)
		p := parse.New(lex)
		inter := eval.New(p)
		v, err := inter.Eval()
		if err != nil {
			panic(err)
		}
		if v != nil {
			fmt.Printf("%v\n", v)
		}
	} else {
		repl.StartREPL(os.Stdin, os.Stdout)
	}
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
