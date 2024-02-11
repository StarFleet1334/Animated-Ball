package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"time"
)

func main() {

	const (
		width  = 50
		height = 10

		cellEmpty = ' '
		ball      = '⚾'

		maxFrames = 1200
	)

	var (
		px, py int
		cell   rune
		vx, vy = 1, 1
	)

	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}

	screen.Clear()

	for i := 0; i < maxFrames; i++ {
		px += vx
		py += vy

		if px <= 0 || px >= width-1 {
			vx *= -1
		}

		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}

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
