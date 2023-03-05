package main

import (
	"fmt"
	"math/rand"
	"snake/screen"
)

type Apple struct {
	x int
	y int
}

func (a *Apple) Render() {
	screen.MoveCursorTo(a.x, a.y)

	fmt.Printf(string(screen.ColorRed))
	fmt.Printf("o")
	fmt.Printf(string(screen.ColorReset))
}

func (a *Apple) Spawn(x, y int) {
	a.x = x
	a.y = y
}

func (a *Apple) Respawn(maxW, maxH int) {
	a.x = rand.Intn(maxW)
	a.y = rand.Intn(maxH)
}
