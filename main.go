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
		out = strings.NewReader("1+(3-2)-(2+3)\n")
	} else {
		out = os.Stdin
	}

	repl.StartREPL(out, os.Stdout)
}
