package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"wt_concurrency"
)

func main() {
	compiler := flag.String("compiler", "clang++", "A C++14 compatible binary to use for compiling.")
	wiredtigerHeader := flag.String("include", "./", "The directory `wiredtiger.h` can be found.")
	wiredtigerLib := flag.String("lib", "./", "The directory where WiredTiger libraries are installed.")
	debug := flag.Bool("debug", false, "Run in debug mode. Does not delete artifacts.")

	flag.Parse()
	tableFilename := flag.Arg(0)

	inputTable, err := os.Open(tableFilename)
	if err != nil {
		panic(err)
	}
	defer inputTable.Close()

	if err := os.MkdirAll("./artifacts/", 0775); err != nil {
		panic(err)
	}

	program, err := wt_concurrency.ParseProgram(inputTable)
	if err != nil {
		panic(err)
	}

	sequenceProgramFilename := "./artifacts/wt_sequence.cpp"
	program.Compile(sequenceProgramFilename)
	if *debug == false {
		defer os.Remove(sequenceProgramFilename)
	}

	cmd := exec.Command(
		*compiler,
		fmt.Sprintf("-I%v", *wiredtigerHeader),
		"-I./",
		fmt.Sprintf("-L%v", *wiredtigerLib),
		"-ggdb",
		"-std=c++14",
		"-fPIC",
		sequenceProgramFilename,
		"-l",
		"wiredtiger",
		"-o",
		"./artifacts/a.out",
	)
	errorOutput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf(
			"Compilation failed.\nCmd:\n%v\n%v\n",
			strings.Join(cmd.Args, " "),
			string(errorOutput),
		)
		os.Exit(1)
	}
	if *debug == false {
		defer os.Remove("./artifacts/a.out")
	}

	programProcess := exec.Command("./artifacts/a.out")
	if wiredtigerLib != nil {
		programProcess.Env = append(programProcess.Env, fmt.Sprintf("LD_LIBRARY_PATH=%v", *wiredtigerLib))
	}
	tableBytes, err := programProcess.CombinedOutput()
	if err != nil {
		fmt.Println(string(tableBytes))
		panic(err)
	}

	sequenceOutputFilename := "./artifacts/wt_sequence.table"
	err = ioutil.WriteFile(sequenceOutputFilename, tableBytes, 0644)
	if err != nil {
		panic(err)
	}
	if *debug == false {
		defer os.Remove(sequenceOutputFilename)
	}

	sequenceOutput, err := os.Open(sequenceOutputFilename)
	if err != nil {
		panic(err)
	}
	defer sequenceOutput.Close()

	err = wt_concurrency.ParseOutput(sequenceOutput)
	if err != nil {
		panic(err)
	}
}
