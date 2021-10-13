package program

type tokenLocation struct {
	row int
	col int
}

type token struct {
	id  uint8
	pos tokenLocation
}
