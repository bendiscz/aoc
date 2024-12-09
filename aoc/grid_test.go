package aoc

import (
	"fmt"
	"testing"
)

type grid = Grid[byte]

func TestTransMatrixView(t *testing.T) {
	m := NewMatrix[byte](Rectangle(3, 4))
	displayGrid(m)
	displayGrid(m.Trans())
}

// ####
// #...
// ###.
// #...
// #...

func TestGridView(t *testing.T) {
	m := NewMatrix[byte](Rectangle(4, 5))
	CopyVector(m.Row(0), "####")
	CopyVector(m.Row(1), "#...")
	CopyVector(m.Row(2), "###.")
	CopyVector(m.Row(3), "#...")
	CopyVector(m.Row(4), "#...")

	show := func(g Grid[byte]) {
		PrintGridFunc(g, func(t byte) string { return string(t) })
		fmt.Println()
	}

	g := gridView[byte]{grid: m}

	show(g)
	show(g.FlipX())
	show(g.FlipY())
	show(g.Trans())
	show(g.Trans().FlipX())
	show(g.Trans().FlipY())
	show(g.RotateR())
	show(g.RotateR().RotateR())
	show(g.RotateR().RotateR().RotateR())
	show(g.RotateR().RotateR().RotateR().RotateR())
	show(g.RotateL())
	show(g.RotateL().RotateL())
	show(g.RotateL().RotateL().RotateL())
	show(g.RotateL().RotateL().RotateL().RotateL())

	sg := g.SubGrid(XY{1, 1}, XY{3, 2})
	show(sg)
	show(sg.RotateR())
	show(sg.RotateR().RotateR())
	show(sg.RotateR().RotateR().RotateR())
	show(sg.RotateR().RotateR().RotateR().RotateR())
}

func displayGrid(g grid) {
	for y := 0; y < g.Size().Y; y++ {
		for x := 0; x < g.Size().X; x++ {
			fmt.Printf("%c", *g.AtXY(x, y)+'0')
		}
		fmt.Println()
	}
}
