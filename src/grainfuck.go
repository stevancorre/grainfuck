package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// flag setup
var (
	build      *bool = flag.Bool("build", false, "Enables compilation")
	run        *bool = flag.Bool("run", false, "Enables simulation")
	memorySize *uint = flag.Uint("mem", 30_000, "Set the memory size")
)

func init() {
	// custom usage
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])

		flag.PrintDefaults()

		fmt.Fprintf(os.Stderr, "\nNote: if you enable both simulation and compilation mode, it's going to first compile the program, then run the executable\n")
	}
}

// print an error to the standard output
func printError(message string) {
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", message)
}

// print the usage of this tool, then exit with code 1
func printUsageAndExit() {
	flag.Usage()
	os.Exit(1)
}

func main() {
	// parse flags (like -compile etc..)
	flag.Parse()

	// check if either simulation or compilation mode was choosen
	if !*run && !*build {
		printError("You need to select at least one execution mode")
		printUsageAndExit()
	}

	// check if any file was provided
	if len(flag.Args()) != 1 {
		printError("No file was provided")
		printUsageAndExit()
	}

	// get the provided file path and check its extension
	// no need to check if it exists, CompileProgram() is going to handle this case
	fpath := flag.Arg(0)
	fext := filepath.Ext(strings.TrimSpace(fpath))

	if fext != ".b" && fext != ".bf" {
		printError("Target file needs to be a brainfuck file (.bf or .b)")
		printUsageAndExit()
	}

	tokens := ParseCommands(fpath)

	if *run {
		Simulate(tokens, *memorySize)
	}
}
