package program

import (
	"bufio"
	"fmt"
	"os"
)

func Assert(condition bool, format string, args ...interface{}) {
	if condition {
		fmt.Printf(format+"\n", args...)
		os.Exit(1)
	}
}

func ParseProgram(fpath string) []token {
	var tokens []token

	// read file
	file, err := os.Open(fpath)
	Assert(err != nil, "ERROR: %s", err)

	// initialize a new scanner, and split by lines
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// iterate through all lines
	row := 0
	for scanner.Scan() {
		line := scanner.Text()

		for col, ch := range line {
			// iterate through all characters
			tokenId := intrinsics[byte(ch)]

			// if token is 0, then do nothing with it
			if tokenId == 0 {
				continue
			}

			// add token to tokens slice
			token := token{
				id: tokenId,
				pos: tokenLocation{
					row: row,
					col: col,
				},
			}

			tokens = append(tokens, token)
		}
	}

	return tokens
}
