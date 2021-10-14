package main

type command struct {
	id  uint8
	row int
	col int

	// for [ ]
	refIp int
}

const (
	// start at 1 so we can use 0 to check if token is valid or not in CompileProgram()
	COMMAND_INCR_PTR uint8 = iota + 1
	COMMAND_DECR_PTR
	COMMAND_INCR_DPTR
	COMMAND_DECR_DPTR
	COMMAND_PCHAR
	COMMAND_GCHAR
	COMMAND_JMPFW
	COMMAND_JMPBW
)

var commandsTable = map[byte]uint8{
	'>': COMMAND_INCR_PTR,
	'<': COMMAND_DECR_PTR,
	'+': COMMAND_INCR_DPTR,
	'-': COMMAND_DECR_DPTR,
	'.': COMMAND_PCHAR,
	',': COMMAND_GCHAR,
	'[': COMMAND_JMPFW,
	']': COMMAND_JMPBW,
}
