package main

import (
	"flag"
	"os"
	"slow-lang/repl"
)

func main() {
	filePtr := flag.String("src", "", "path to .sl file")
	flag.Parse()

	if *filePtr != "" {
		repl.StartFile(*filePtr, os.Stdout)
		return
	}

	repl.Start(os.Stdin, os.Stdout)
}
