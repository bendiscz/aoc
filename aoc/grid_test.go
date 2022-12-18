package aoc

import (
	"fmt"
	"testing"
)

type grid = Grid[byte]

func TestTransMatrixView(t *testing.T) {
	m := NewMatrix[byte](Rectangle(3, 4))
	displayGrid(m)
	displayGrid(m.TransView())
}

func displayGrid(g grid) {
	for y := 0; y < g.Size().Y; y++ {
		for x := 0; x < g.Size().X; x++ {
			fmt.Printf("%c", *g.AtXY(x, y)+'0')
		}
		fmt.Println()
	}
}
