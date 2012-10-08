package main

import (
	"fmt"
)

const (
	Black = iota
	White
	Red
	Yellow
	Green
)

func PaintFill(screen [][]int, x, y, ocolor, ncolor int) bool {
	if x < 0 || x > len(screen[0]) - 1 || y < 0 || y > len(screen) - 1 {
		return false
	}
	if (screen[x][y] == ocolor) {
		screen[x][y] = ncolor
		PaintFill(screen, x - 1, y, ocolor, ncolor)
		PaintFill(screen, x + 1, y, ocolor, ncolor)
		PaintFill(screen, x, y - 1, ocolor, ncolor)
		PaintFill(screen, x, y + 1, ocolor, ncolor)
	}
	return true
}

func Paint(screen [][]int, x, y, color int) bool {
	return PaintFill(screen, x, y, screen[x][y], color)
}

func main() {
	screen := [][]int{{0, 1, 2}, {0, 1, 2}, {0, 1, 2}}
	fmt.Println(Paint(screen, 0, 1, 3))
	for _, v := range screen {
		fmt.Println(v)
	}
}
