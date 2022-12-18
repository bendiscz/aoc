package main

import (
	"math"
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2022
	day     = 14
	example = `

498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9

`
)

func main() {
	Run(year, day, example, solve)
}

const (
	Empty = iota
	Full
	Sand
)

type cell struct {
	block byte
}

func (c cell) String() string {
	switch c.block {
	case Sand:
		return "o"
	case Full:
		return SymbolFull
	default:
		return SymbolEmpty
	}
}

func draw(m *Matrix[cell], c1, c2 XY) {
	if c1.X == c2.X {
		d := Sign(c2.Y - c1.Y)
		for y := c1.Y; y != c2.Y+d; y += d {
			m.AtXY(c1.X, y).block = Full
		}
	} else {
		d := Sign(c2.X - c1.X)
		for x := c1.X; x != c2.X+d; x += d {
			m.AtXY(x, c1.Y).block = Full
		}
	}
}

func fill(m *Matrix[cell], c XY) (int, bool) {
	if !m.Dim.HasInside(c) {
		return 0, false
	}
	if m.At(c).block != Empty {
		return 0, true
	}

	count := 0
	for _, d := range []int{0, -1, 1} {
		n, s := fill(m, XY{c.X + d, c.Y + 1})
		count += n
		if !s {
			return count, false
		}
	}

	m.At(c).block = Sand
	return count + 1, true
}

func solve(p *Problem) {
	min, max := XY{math.MaxInt, math.MaxInt}, XY{math.MinInt, math.MinInt}
	var lines [][]XY
	for p.NextLine() {
		var line []XY
		for _, s := range strings.Split(p.Line(), " -> ") {
			coords := ParseInts(s)
			c := XY{coords[0], coords[1]}
			line = append(line, c)
			min = XY{Min(c.X, min.X), Min(c.Y, min.Y)}
			max = XY{Max(c.X, max.X), Max(c.Y, max.Y)}

		}
		lines = append(lines, line)
	}

	orig := XY{500, 0}
	max.Y += 2
	min.X = Min(orig.X-max.Y, min.X)
	max.X = Max(orig.X+max.Y, max.X)

	c0 := XY{min.X, 0}
	d := max.Sub(c0).Add(XY{1, 1})
	orig = orig.Sub(c0)

	m := NewMatrix[cell](d)
	for _, line := range lines {
		for i := 1; i < len(line); i++ {
			draw(m, line[i-1].Sub(c0), line[i].Sub(c0))
		}
	}

	count, _ := fill(m, orig)
	PrintGrid[cell](m)
	p.PartOne(count)

	for x := 0; x < m.Dim.X; x++ {
		for y := 0; y < m.Dim.Y-1; y++ {
			if m.AtXY(x, y).block == Sand {
				m.AtXY(x, y).block = Empty
			}
		}
		m.AtXY(x, m.Dim.Y-1).block = Full
	}

	count, _ = fill(m, orig)
	//PrintGrid[cell](m)
	p.PartTwo(count)
}
