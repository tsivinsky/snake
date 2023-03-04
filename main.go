package main

import (
	"bufio"
	"log"
	"os"
	"snake/screen"
	"time"

	"golang.org/x/term"
)

const FPS = 40

var FD = int(os.Stdin.Fd())

const (
	KeyQ = 0x71
	KeyH = 0x68
	KeyJ = 0x6a
	KeyK = 0x6b
	KeyL = 0x6c
)

func main() {
	oldState, err := term.MakeRaw(FD)
	if err != nil {
		log.Fatal(err)
	}
	defer term.Restore(FD, oldState)

	w, h, err := term.GetSize(FD)
	if err != nil {
		log.Fatal(err)
	}

	snake := &Snake{head: [2]int{w / 2, h / 2}, dir: "right"}
	apple := &Apple{x: 12, y: 12}

	screen.MoveCursorTo(w/2, h/2)

	go func() {
		r := bufio.NewReaderSize(os.Stdin, 1)

		for {
			b, err := r.ReadByte()
			if err != nil {
				continue
			}

			switch b {
			case KeyQ:
				os.Exit(0)

			case KeyH:
				snake.MoveLeft()
				break

			case KeyJ:
				snake.MoveDown()
				break

			case KeyK:
				snake.MoveUp()
				break

			case KeyL:
				snake.MoveRight()
				break
			}
		}
	}()

	for {
		screen.Clear()

		if snake.IsEating(apple) {
			apple.Respawn(w, h)
			snake.Grow()
		}

		snake.Move()

		if snake.IsOutOfBounds(w, h) {
			os.Exit(0)
		}

		apple.Render()
		snake.Render()

		time.Sleep(time.Second / FPS)
	}
}
