package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// flag setup
var (
	build      *bool   = flag.Bool("build", false, "Enables compilation")
	run        *bool   = flag.Bool("run", false, "Enables simulation")
	outputPath *string = flag.String("o", "", "Set the output build path")
	memorySize *uint   = flag.Uint("mem", 30_000, "Set the memory size")
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

	program := ParseCommands(fpath)

	if *build {
		// TODO: check if output path is a file or smth

		// get working directory, and join with output path is output
		// path isn't absolute
		wd, _ := os.Getwd()
		if !filepath.IsAbs(*outputPath) {
			*outputPath = filepath.Join(wd, *outputPath)

			// if output path doesn't exists, or isn't a directory
			// we want to considere the output path as a file
			dir, err := os.Stat(*outputPath)
			if err != nil || !dir.IsDir() {
				fpath = ""
			}
		}

		// remove extension and join with file name
		fpath = strings.TrimSuffix(fpath, path.Ext(fpath))
		opath := path.Join(*outputPath, filepath.Base(fpath))

		// compile to asm, then build using nasm
		CompileProgram(opath, program, *memorySize)
		BuildProgram(opath)

		if *run {
			// build and run
			RunProgram(opath)
		}
	} else if *run {
		// run
		SimulateProgram(program, *memorySize)
	}
}
