package program

const (
	// start at 1 so we can use 0 to check if token is valid or not in CompileProgram()
	INTRINSIC_INCR_PTR uint8 = iota + 1
	INTRINSIC_DECR_PTR
	INTRINSIC_INCR_DPTR
	INTRINSIC_DECR_DPTR
	INTRINSIC_PCHAR
	INTRINSIC_GCHAR
	INTRINSIC_WHILE
	INTRINSIC_END
)

var intrinsics = map[byte]uint8{
	'>': INTRINSIC_INCR_PTR,
	'<': INTRINSIC_DECR_PTR,
	'+': INTRINSIC_INCR_DPTR,
	'-': INTRINSIC_DECR_DPTR,
	'.': INTRINSIC_PCHAR,
	',': INTRINSIC_GCHAR,
	'[': INTRINSIC_WHILE,
	']': INTRINSIC_END,
}
