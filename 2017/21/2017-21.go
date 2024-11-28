package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 21
	example = `

../.# => ##./#../...
.#./..#/### => #..#/..../..../#..#

`
)

type cell struct {
	v bool
}

func (c cell) String() string {
	if c.v {
		return SymbolFull
	} else {
		return SymbolEmpty
	}
}

type grid = *Matrix[cell]

type mapping struct {
	s, d grid
}

func parseMapping2(line []byte) mapping {
	s, d := NewMatrix[cell](Square(2)), NewMatrix[cell](Square(3))
	parseGrid(s, line[0:])
	parseGrid(d, line[9:])
	return mapping{s, d}
}

func parseMapping3(line []byte) mapping {
	s, d := NewMatrix[cell](Square(3)), NewMatrix[cell](Square(4))
	parseGrid(s, line[0:])
	parseGrid(d, line[15:])
	return mapping{s, d}
}

func parseGrid(g grid, line []byte) {
	g.Dim.ForEach(func(xy XY) {
		g.At(xy).v = line[(g.Dim.X+1)*xy.Y+xy.X] == '#'
	})
}

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	var m2, m3 []mapping
	for p.NextLine() {
		if len(p.Line()) == 20 {
			m2 = append(m2, parseMapping2(p.LineBytes()))
		} else {
			m3 = append(m3, parseMapping3(p.LineBytes()))
		}
	}

	g := NewMatrix[cell](Square(3))
	parseGrid(g, []byte(".#./..#/###"))

	for i := 0; i < 5; i++ {
		g = iterate(g, m2, m3)
	}
	p.PartOne(count(g))

	for i := 0; i < 13; i++ {
		g = iterate(g, m2, m3)
	}
	p.PartTwo(count(g))
}

func count(g grid) int {
	c := 0
	g.Dim.ForEach(func(xy XY) {
		if g.At(xy).v {
			c++
		}
	})
	return c
}

func iterate(g grid, m2, m3 []mapping) grid {
	n, ns, nd, mappings := g.Dim.X, 3, 4, m3
	if n%2 == 0 {
		ns, nd, mappings = 2, 3, m2
	}

	n0 := n / ns
	r := NewMatrix[cell](Square(n0 * nd))
	for x0 := 0; x0 < n0; x0++ {
		for y0 := 0; y0 < n0; y0++ {

			m, ok := find(g, mappings, x0*ns, y0*ns)
			if !ok {
				return g
			}
			fill(r, m.d, x0*nd, y0*nd)
		}
	}

	return r
}

func find(g grid, mappings []mapping, x0, y0 int) (mapping, bool) {
	for _, m := range mappings {
		if matchesTransformed(g, m.s, x0, y0) {
			return m, true
		}
	}
	return mapping{}, false
}

type transformer func(xy, dim XY) XY

func rot(xy, dim XY) XY  { return XY{dim.Y - xy.Y - 1, xy.X} }
func flip(xy, dim XY) XY { return XY{dim.X - xy.X - 1, xy.Y} }

func matchesTransformed(g, p grid, x0, y0 int) bool {
	return matches(g, p, x0, y0) ||
		matches(g, p, x0, y0, rot) ||
		matches(g, p, x0, y0, rot, rot) ||
		matches(g, p, x0, y0, rot, rot, rot) ||
		matches(g, p, x0, y0, flip) ||
		matches(g, p, x0, y0, flip, rot) ||
		matches(g, p, x0, y0, flip, rot, rot) ||
		matches(g, p, x0, y0, flip, rot, rot, rot)
}

func matches(g, p grid, x0, y0 int, tr ...transformer) bool {
	m := true
	p.Dim.ForEach(func(xy XY) {
		pc := xy
		for _, t := range tr {
			pc = t(pc, p.Dim)
		}
		if g.AtXY(x0+xy.X, y0+xy.Y).v != p.At(pc).v {
			m = false
		}
	})
	return m
}

func fill(g, p grid, x0, y0 int) {
	p.Dim.ForEach(func(xy XY) {
		g.AtXY(x0+xy.X, y0+xy.Y).v = p.At(xy).v
	})
}
