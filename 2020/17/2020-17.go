package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"maps"
	"slices"
)

const (
	year    = 2020
	day     = 17
	example = `

.#.
..#
###

`
)

func main() {
	Run(year, day, example, solve)
}

type xyz struct {
	x, y, z, w int
}

type cell struct {
	a bool
	n int
}

type grid struct {
	cells map[xyz]*cell
}

func newGrid() grid {
	return grid{make(map[xyz]*cell)}
}

func solve(p *Problem) {
	g1, g2 := newGrid(), newGrid()
	for y := 0; p.NextLine(); y++ {
		s := p.Line()
		for x := 0; x < len(s); x++ {
			if s[x] == '#' {
				g1.cells[xyz{x: x, y: y}] = &cell{a: true}
				g2.cells[xyz{x: x, y: y}] = &cell{a: true}
			}
		}
	}

	for i := 0; i < 6; i++ {
		g1 = g1.step(false)
		g2 = g2.step(true)
	}
	p.PartOne(g1.countActive())
	p.PartTwo(g2.countActive())
}

func (g grid) step(hyper bool) grid {
	for _, v := range slices.Collect(maps.Keys(g.cells)) {
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				for z := -1; z <= 1; z++ {
					for w := -1; w <= 1; w++ {
						if !hyper && w != 0 {
							continue
						}
						if x == 0 && y == 0 && z == 0 && w == 0 {
							continue
						}
						g.mark(xyz{v.x + x, v.y + y, v.z + z, v.w + w})
					}
				}
			}
		}
	}

	ng := newGrid()
	for v, c := range g.cells {
		if c.a && (c.n == 2 || c.n == 3) || !c.a && c.n == 3 {
			ng.cells[v] = &cell{a: true}
		}
	}

	return ng
}

func (g grid) mark(v xyz) {
	c := g.cells[v]
	if c == nil {
		c = &cell{}
		g.cells[v] = c
	}
	c.n++
}

func (g grid) countActive() int {
	s := 0
	for _, c := range g.cells {
		if c.a {
			s++
		}
	}
	return s
}
