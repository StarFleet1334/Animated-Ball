package main

import "fmt"

func main() {

	const (
		width  = 50
		height = 10

		cellEmpty = ' '
		ball      = '⚾'
	)

	var cell rune

	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}

	buf := make([]rune, 0, width*height)

	board[0][0] = true

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

	fmt.Println(string(buf))
}
