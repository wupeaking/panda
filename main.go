package main

import (
	"os"
	"panda/repl"
)

func main() {
	repl.StartREPL(os.Stdin, os.Stdout)
}
