package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"time"
)

func main() {

	const (
		width, height = 50, 10
		cellEmpty     = ' '
		ball          = '⚾'

		maxFrames = 1200
	)

	var (
		px, py   int
		cell     rune
		vx, vy   = 1, 1
		ppx, ppy int
	)

	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}

	screen.Clear()

	for i := 0; i < maxFrames; i++ {
		px += vx
		py += vy

		if px <= 0 || px >= width-vx {
			vx *= -1
		}

		if py <= 0 || py >= height-vy {
			vy *= -1
		}

		board[px][py], board[ppx][ppy] = true, false

		ppx = px
		ppy = py

		board[px][py] = true

		buf := make([]rune, 0, (width*2+1)*height)

		buf = buf[:0]

		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = ball
				}
				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}
		screen.MoveTopLeft()
		fmt.Println(string(buf))
		fmt.Printf("Cap: %d - Height: %d", cap(buf), (width*2+1)*height)
		time.Sleep(time.Second / 20)
	}
}
