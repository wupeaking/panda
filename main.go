package main

import (
	"os"
	"panda/repl"
	"strings"
)

func main() {

	repl.StartREPL(strings.NewReader("1+2*3\n"), os.Stdout)
}
