// TODO: enable wrapping

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func Assert(condition bool, format string, args ...interface{}) {
	if !condition {
		log.Printf(format+"\n", args...)
		os.Exit(1)
	}
}

// Parse source file to commands
func ParseCommands(fpath string) []command {
	var commands []command

	// read file
	file, err := ioutil.ReadFile(fpath)
	Assert(err == nil, "ERROR: %s", err)

	src := string(file)

	var ipStack []int

	ip := 0
	for _, ch := range src {
		tokenId := commandsTable[byte(ch)]

		// if token is 0, then do nothing with it
		if tokenId == 0 {
			continue
		}

		token := command{
			id: tokenId,
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

		// add token to tokens slice
		commands = append(commands, token)

		ip += 1
	}

	return commands
}

// Simulate program
func SimulateProgram(commands []command, memSize uint) {
	// data buffer for simulion
	mem := make([]byte, memSize)
	ptr := 0

	// iterate through all commands
	// TODO: pre optimization
	// TODO: do something that looks better
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

// Compile program to assembly
func CompileProgram(opath string, commands []command, memSize uint) {
	// TODO: check error here
	file, _ := os.OpenFile(opath+".asm", os.O_CREATE|os.O_WRONLY, 0644)
	//Assert(err == nil, err.Error())
	datawriter := bufio.NewWriter(file)

	// data
	datawriter.WriteString("section .data\n")
	datawriter.WriteString("    mem: times 30000 db 0		 	; memory\n")
	datawriter.WriteString("    memp: dd 0					 	; position in memory\n")

	// asm
	datawriter.WriteString("section .text\n")

	datawriter.WriteString("extern putchar, getchar\n")
	datawriter.WriteString("global main\n")

	// printchar
	datawriter.WriteString("printChar:")
	datawriter.WriteString("    mov edx, dword[memp]\n")
	datawriter.WriteString("    push eax\n")
	datawriter.WriteString("    push ecx\n")
	datawriter.WriteString("    mov edx, dword[memp]\n")
	datawriter.WriteString("    mov eax, 0\n")
	datawriter.WriteString("    mov al, byte[edx]\n")
	datawriter.WriteString("    push eax\n")
	datawriter.WriteString("    call putchar\n")
	datawriter.WriteString("    add esp, 4\n")
	datawriter.WriteString("    pop ecx\n")
	datawriter.WriteString("    pop eax\n")
	datawriter.WriteString("    ret\n")

	// read char
	datawriter.WriteString("readChar:\n")
	datawriter.WriteString("    push eax\n")
	datawriter.WriteString("    push ecx\n")
	datawriter.WriteString("    call getchar\n")
	datawriter.WriteString("    mov ecx, dword[memp]\n")
	datawriter.WriteString("    mov byte[ecx], al\n")
	datawriter.WriteString("    pop ecx\n")
	datawriter.WriteString("    pop eax\n")
	datawriter.WriteString("    ret\n")

	datawriter.WriteString("main:\n")
	datawriter.WriteString("    push ebp\n")
	datawriter.WriteString("    mov ebp, esp\n")
	datawriter.WriteString("    mov dword[memp], mem		; initialize pointer\n")
	datawriter.WriteString("    mov ecx, dword[ebp + 8] 	; ecx - address of currently executed char\n")

	// iterate through all commands
	// TODO: pre optimization
	// TODO: make something better lol
	for ip := 0; ip < len(commands); ip++ {
		op := commands[ip]

		switch op.id {
		case COMMAND_INCR_PTR:
			datawriter.WriteString("    inc dword[memp]\n")
			break

		case COMMAND_DECR_PTR:
			datawriter.WriteString("    dec dword[memp]\n")
			break

		case COMMAND_INCR_DPTR:
			datawriter.WriteString("    mov edx, dword[memp]\n")
			datawriter.WriteString("    inc byte[edx]\n")
			break

		case COMMAND_DECR_DPTR:
			datawriter.WriteString("    mov edx, dword[memp]\n")
			datawriter.WriteString("    dec byte[edx]\n")
			break

		case COMMAND_PCHAR:
			datawriter.WriteString("    call printChar\n")
			break

		case COMMAND_GCHAR:
			datawriter.WriteString("    call readChar\n")

		case COMMAND_JMPFW:
			datawriter.WriteString(fmt.Sprintf("addr_%d:\n", ip))
			datawriter.WriteString("    mov edx, dword[memp]\n")
			datawriter.WriteString("    mov al, byte[edx]\n")
			datawriter.WriteString("    test al, al\n")
			datawriter.WriteString(fmt.Sprintf("    jz addr_%d\n", op.refIp))
			break

		case COMMAND_JMPBW:
			datawriter.WriteString(fmt.Sprintf("addr_%d:\n", ip))
			datawriter.WriteString("    mov edx, dword[memp]\n")
			datawriter.WriteString("    mov al, byte[edx]\n")
			datawriter.WriteString("    test al, al\n")
			datawriter.WriteString(fmt.Sprintf("    jnz addr_%d\n", op.refIp))
			break

		default:
			panic("Unreachable")
		}
	}

	// exit
	datawriter.WriteString("    mov esp, ebp\n")
	datawriter.WriteString("    pop ebp\n")
	datawriter.WriteString("    ret\n")

	// close everything
	datawriter.Flush()
	file.Close()
}

// Run a command and output it results
func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()
	Assert(err == nil, "ERROR:\n%s\n%", out, err.Error())
}

// Compiled using NASM and GCC
func BuildProgram(opath string) {
	runCommand("nasm", "-felf32", fmt.Sprintf("%s.asm", opath))
	runCommand("gcc", "-m32", fmt.Sprintf("%s.o", opath), "-o", opath)
}

// Run the executable
func RunProgram(opath string) {
}
