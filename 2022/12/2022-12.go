package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2022
	day     = 12
	example = `

Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	XY
	m *Matrix[byte]
}

func (c cell) Key() any {
	return c.XY
}

func (c cell) Edges() []Edge {
	var e []Edge
	for _, d := range HVDirs {
		x := c.Add(d)
		if !c.m.Dim.HasInside(x) {
			continue
		}
		ch, ch2 := *c.m.At(c.XY), *c.m.At(x)
		if int(ch)-int(ch2) <= 1 {
			e = append(e, Edge{V: cell{x, c.m}, W: 1})
		}
	}
	return e
}

func solve(p *Problem) {
	var m *Matrix[byte]
	var s, e XY
	for y := 0; p.NextLine(); y++ {
		if m == nil {
			m = NewMatrix[byte](Rectangle(len(p.Line()), 0))
		}
		row := m.AppendRow()
		for x, ch := range []byte(p.Line()) {
			if ch == 'S' {
				ch = 'a'
				s = XY{x, y}
			} else if ch == 'E' {
				ch = 'z'
				e = XY{x, y}
			}
			*row.At(x) = ch
		}
	}

	paths := ShortestPaths(cell{e, m}, func(v cell) (bool, bool) {
		return *v.m.At(v.XY) == 'a', false
	})

	for _, path := range paths {
		if path.Steps[len(path.Steps)-1].XY == s {
			p.PartOne(path.Cost)
		}
	}

	p.PartTwo(paths[0].Cost)
}
