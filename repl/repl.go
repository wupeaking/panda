package repl

import (
	"bufio"
	"fmt"
	"io"
	"panda/eval"
	"panda/lexer"
	"panda/parse"
	"strings"
)

func StartREPL(in io.Reader, out io.Writer) {
	inter := eval.New(nil)
	for {
		fmt.Fprintf(out, ">> ")

		input := bufio.NewReader(in)
		stream, err := input.ReadString('\n')
		if err != nil {
			fmt.Fprintf(out, ">> %v", err)
			break
		}
		lex := lexer.New(strings.NewReader(stream))
		// for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
		// 	fmt.Fprintf(out, ">> %v\n", tok)
		// }
		p := parse.New(lex)
		inter.SetParser(p)
		// fmt.Fprintf(out, " %v\n",)

		ret, err := inter.Eval()
		if err != nil {
			fmt.Fprintf(out, " %v\n", err)
		} else {
			fmt.Fprintf(out, " %v\n", ret)
		}

	}
}
