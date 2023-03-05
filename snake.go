package main

import (
	"fmt"
	"snake/screen"
)

type Part [2]int

type Snake struct {
	head   Part
	blocks []Part
	dir    string
}

func (s *Snake) Render() {
	for _, block := range s.blocks {
		x, y := block[0], block[1]
		screen.MoveCursorTo(x, y)

		fmt.Printf(string(screen.ColorGreen))
		fmt.Printf("x")
		fmt.Printf(string(screen.ColorReset))
	}

	screen.MoveCursorTo(s.head[0], s.head[1])
}

func (s *Snake) MoveRight() {
	s.head[0] += 1
	s.dir = "right"
}

func (s *Snake) MoveLeft() {
	s.head[0] -= 1
	s.dir = "left"
}

func (s *Snake) MoveUp() {
	s.head[1] -= 1
	s.dir = "up"
}

func (s *Snake) MoveDown() {
	s.head[1] += 1
	s.dir = "down"
}

func (s *Snake) IsEating(a *Apple) bool {
	if s.head[0] == a.x && s.head[1] == a.y {
		return true
	}

	return false
}

func (s *Snake) Grow() {
	s.blocks = append(s.blocks, [2]int{s.head[0], s.head[1]})
}

func (s *Snake) Move() {
	switch s.dir {
	case "up":
		s.MoveUp()
		break

	case "down":
		s.MoveDown()
		break

	case "left":
		s.MoveLeft()
		break

	case "right":
		s.MoveRight()
		break
	}
}

func (s *Snake) IsOutOfBounds(w, h int) bool {
	x, y := s.head[0], s.head[1]

	if x < 0 {
		return true
	}

	if x >= w {
		return true
	}

	if y < 0 {
		return true
	}

	if y >= h {
		return true
	}

	return false
}

func (s *Snake) IsCollidingWithItself() bool {
	x, y := s.head[0], s.head[1]

	for _, block := range s.blocks {
		if x == block[0] && y == block[1] {
			return true
		}
	}

	return false
}
