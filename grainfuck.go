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
	compile  *bool = flag.Bool("compile", false, "Enables compilation")
	simulate *bool = flag.Bool("simulate", false, "Enables simulation")
)

func init() {
	// custom usage
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])

		flag.PrintDefaults()

		fmt.Fprintf(os.Stderr, "\nNote: if you enable both simulation and compilation mode, it's going to first simulate the program, then compile it if no error was detected.")
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

	// check if any file was provided
	if len(flag.Args()) != 1 {
		printUsageAndExit()
	}

	// get the provided file path and check its extension
	var (
		fpath string = flag.Arg(0)
		fext  string = filepath.Ext(strings.TrimSpace(fpath))
	)

	if fext != ".b" && fext != ".bf" {
		printError("Target file needs to be a brainfuck file (.bf or .b)")
		printUsageAndExit()
	}
}

func compileProgram() {

}
