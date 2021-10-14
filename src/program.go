package main

import (
	"bufio"
	"fmt"
	"os"
)

func Assert(condition bool, format string, args ...interface{}) {
	if !condition {
		fmt.Printf(format+"\n", args...)
		os.Exit(1)
	}
}

func ParseCommands(fpath string) []command {
	var commands []command

	// read file
	file, err := os.Open(fpath)
	Assert(err == nil, "ERROR: %s", err)

	// initialize a new scanner, and split by lines
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var ipStack []int

	// iterate through all lines
	row := 0
	ip := 0

	for scanner.Scan() {
		line := scanner.Text()

		for col, ch := range line {
			// iterate through all characters
			tokenId := commandsTable[byte(ch)]

			// if token is 0, then do nothing with it
			if tokenId == 0 {
				continue
			}

			// add token to tokens slice
			token := command{
				id:  tokenId,
				row: row,
				col: col,
			}

			if token.id == COMMAND_JMPFW {
				ipStack = append(ipStack, ip)
			} else if token.id == COMMAND_JMPBW {
				// pop last element from stack
				ipStackSize := len(ipStack)
				refIp := ipStack[ipStackSize-1]
				ipStack = ipStack[:ipStackSize-1]

				// link commands
				commands[refIp].refIp = ip
				token.refIp = refIp
			}

			commands = append(commands, token)

			ip += 1
		}
	}

	return commands
}

func Simulate(commands []command, memSize uint) {
	// data buffer for simulion
	mem := make([]byte, memSize)
	ptr := 0

	// iterate through all commands
	for ip := 0; ip < len(commands); ip++ {
		op := commands[ip]

		switch op.id {
		case COMMAND_INCR_PTR:
			ptr++
			break

		case COMMAND_DECR_PTR:
			ptr--
			break

		case COMMAND_INCR_DPTR:
			mem[ptr]++
			break

		case COMMAND_DECR_DPTR:
			mem[ptr]--
			break

		case COMMAND_PCHAR:
			fmt.Printf("%c", mem[ptr])
			break

		case COMMAND_GCHAR:
			// read char from stdin
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')

			mem[ptr] = input[0]
			break

		case COMMAND_JMPFW:
			if mem[ptr] == 0 {
				ip = op.refIp
			}
			break

		case COMMAND_JMPBW:
			if mem[ptr] != 0 {
				ip = op.refIp
			}
			break

		default:
			panic("Unreachable")
		}
	}
}
