package screen

import "golang.org/x/term"

var Colors = new(term.EscapeCodes)

var (
	ColorReset = "\033[0m"

	ColorRed   = "\033[31m"
	ColorGreen = "\033[32m"
)
