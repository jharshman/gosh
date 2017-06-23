package cmd

import "strings"

const EXIT string = "exit"

func Execute(in string) int {
	if strings.Compare(in, EXIT) == 0 {
		return 1
	}
	// run lexers and parsers
	return 0
}
