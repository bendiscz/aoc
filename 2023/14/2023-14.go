package main

import (
	"fmt"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2023
	day     = 14
	example = `

O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	v byte
}

func (c cell) String() string { return fmt.Sprintf("%c", c.v) }

type grid struct {
	*Matrix[cell]
}

func solve(p *Problem) {
	g := grid{}
	for p.NextLine() {
		if g.Matrix == nil {
			g.Matrix = NewMatrix[cell](Rectangle(len(p.Line()), 0))
		}
		ParseVectorFunc(g.AppendRow(), p.Line(), func(b byte) cell { return cell{b} })
	}

	rollUp(g)
	p.PartOne(load(g))
	rollUp(g.Trans())
	rollDown(g)
	rollDown(g.Trans())

	grids, cycles, period := []grid(nil), 1, 0
	for period == 0 {
		grids = append(grids, grid{CloneGrid[cell](g)})
		cycle(g)
		cycles++
		for i := 1; i <= len(grids); i++ {
			if compare(g, grids[len(grids)-i]) {
				period = i
			}
		}
	}
	p.PartTwo(load(grids[len(grids)-period+(1_000_000_000-cycles)%period]))
}

func load(g grid) int {
	total := 0
	for x := 0; x < g.Dim.X; x++ {
		for y := 0; y < g.Dim.Y; y++ {
			if g.AtXY(x, y).v == 'O' {
				total += g.Dim.Y - y
			}
		}
	}
	return total
}

func cycle(g grid) {
	rollUp(g)
	rollUp(g.Trans())
	rollDown(g)
	rollDown(g.Trans())
}

func rollUp(g Grid[cell]) {
	dim := g.Size()
	for x := 0; x < dim.X; x++ {
		y0 := 0
		for y := 0; y < dim.Y; y++ {
			ch := g.AtXY(x, y).v
			if ch == '#' {
				y0 = y + 1
			}
			if ch == 'O' {
				g.AtXY(x, y).v = '.'
				g.AtXY(x, y0).v = 'O'
				y0++
			}
		}
	}
}

func rollDown(g Grid[cell]) {
	dim := g.Size()
	for x := 0; x < dim.X; x++ {
		y0 := dim.Y - 1
		for y := dim.Y - 1; y >= 0; y-- {
			ch := g.AtXY(x, y).v
			if ch == '#' {
				y0 = y - 1
			}
			if ch == 'O' {
				g.AtXY(x, y).v = '.'
				g.AtXY(x, y0).v = 'O'
				y0--
			}
		}
	}
}

func compare(g1, g2 grid) bool {
	for x := 0; x < g1.Dim.X; x++ {
		for y := 0; y < g1.Dim.Y; y++ {
			if g1.AtXY(x, y).v != g2.AtXY(x, y).v {
				return false
			}
		}
	}
	return true
}
