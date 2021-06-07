package main

import (
	"fmt"
	"log"
	"os"
	"panda/eval"
	"panda/lexer"
	"panda/parse"
	"panda/repl"
	"strings"

	"github.com/urfave/cli/v2"
)

var debug = false

func main() {
	if debug {
		debugfunc()
		return
	}

	app := &cli.App{
		Name:      "Panda",
		Version:   "0.0.1",
		Usage:     "一个玩具性质的脚本解释器",
		ArgsUsage: "源文件路径",
		Commands: []*cli.Command{
			{
				Name:        "ast",
				Aliases:     []string{"ast"},
				Usage:       "打印出抽象语法树",
				Description: "读取源文件 打印抽象语法树",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "file", Aliases: []string{"f"}, Usage: "源文件路径", Required: true},
				},
				Action: func(c *cli.Context) error {
					return printAst(c.String("file"))
				},
			},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() > 1 {
				file, err := os.Open(c.Args().First())
				if err != nil {
					return err
				}
				defer file.Close()
				lex := lexer.New(file)
				p := parse.New(lex)
				inter := eval.New(p)
				v, err := inter.Eval()
				if err != nil {
					return err
				}
				if v != nil {
					fmt.Printf("%v\n", v)
				}
			} else {
				repl.StartREPL(os.Stdin, os.Stdout)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func printAst(f string) error {
	file, err := os.Open(f)
	if err != nil {
		return err
	}
	defer file.Close()
	lex := lexer.New(file)
	p := parse.New(lex)
	trees := p.ParserAST()
	if p.Errors() != nil {
		return p.Errors()
	}
	for i := range trees.NodeTrees {
		fmt.Println(trees.NodeTrees[i])
	}
	return nil
}

func debugfunc() {
	lex := lexer.New(strings.NewReader(`
	var a = {
		"name": "wpx",
		"age" : 10+20,
		"sex": true
	}
	`))
	p := parse.New(lex)
	inter := eval.New(p)
	v, err := inter.Eval()
	if err != nil {
		panic(err)
	}
	if v != nil {
		fmt.Printf("%v\n", v)
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
