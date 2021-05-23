package main

import (
	"io"
	"os"
	"panda/repl"
	"strings"
)

var debug = false

func main() {
	var out io.Reader
	if debug {
		out = strings.NewReader(`var a = 1+2*3;
		a;
		`)
	} else {
		out = os.Stdin
	}

	repl.StartREPL(out, os.Stdout)
}
