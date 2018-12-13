package main

import (
	"flag"
	"os"
	"wt_concurrency"
)

func main() {
	flag.Parse()
	tableFilename := flag.Arg(0)

	file, err := os.Open(tableFilename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	program, err := wt_concurrency.ParseProgram(file)
	if err != nil {
		panic(err)
	}

	program.Compile("./artifacts/wt_sequence.cpp")
}
