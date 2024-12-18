package main

import (
	"github.com/bendiscz/aoc/2019/intcode"
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2019
	day     = 11
	example = ``
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	prog := intcode.Parse(p)

	panels := map[XY]int{}
	paint(prog, panels)
	p.PartOne(len(panels))

	clear(panels)
	panels[XY0] = 1
	paint(prog, panels)

	dim := XY0
	for c, w := range panels {
		if w == 1 {
			dim = MaxXY(dim, c)
		}
	}

	g := NewMatrix[bool](dim.Add(Square(1)))
	for c, w := range panels {
		if w == 1 {
			*g.At(c) = true
		}
	}
	PrintGridBool(g)
}

func paint(prog *intcode.Program, panels map[XY]int) {
	at, d := XY{}, NegY
	c := prog.Exec()

	for {
		out := c.ReadWrite(panels[at])
		if len(out) == 0 {
			break
		}

		panels[at] = out[0]
		if out[1] == 0 {
			d = XY{d.Y, -d.X}
		} else {
			d = XY{-d.Y, d.X}
		}
		at = at.Add(d)
	}
}
