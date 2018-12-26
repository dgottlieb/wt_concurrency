package main

import (
	"flag"
	"os"
	"wt_concurrency"
)

func main() {
	flag.Parse()
	emittedFile := flag.Arg(0)

	file, err := os.Open(emittedFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = wt_concurrency.ParseOutput(file)
	if err != nil {
		panic(err)
	}
}
