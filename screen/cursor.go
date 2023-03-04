package screen

import "fmt"

func MoveCursorTo(x, y int) {
	fmt.Printf("\033[%d;%dH", y, x)
}
