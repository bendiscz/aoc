package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"strings"
)

const (
	year    = 2018
	day     = 17
	example = `

x=495, y=2..7
y=7, x=495..501
x=501, y=3..7
x=498, y=2..4
x=506, y=1..2
x=498, y=10..13
x=504, y=10..13
y=13, x=498..504

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	ch byte
}

func (c cell) String() string { return string([]byte{c.ch}) }

type grid struct {
	*Matrix[cell]
}

type line struct {
	a      int
	b0, b1 int
}

func (l line) fillH(g Grid[cell], o XY) {
	for x := l.b0; x <= l.b1; x++ {
		g.AtXY(x-o.X, l.a-o.Y).ch = '#'
	}
}

func solve(p *Problem) {
	hs, vs := []line(nil), []line(nil)
	water := XY{X: 500, Y: 0}
	minXY, maxXY := MaxIntXY, MinIntXY
	for p.NextLine() {
		f := ParseInts(p.Line())
		if strings.HasPrefix(p.Line(), "y=") {
			hs = append(hs, line{a: f[0], b0: f[1], b1: f[2]})
			minXY = MinXY(minXY, XY{f[1], f[0]})
			maxXY = MaxXY(maxXY, XY{f[2], f[0]})
		} else {
			vs = append(vs, line{a: f[0], b0: f[1], b1: f[2]})
			minXY = MinXY(minXY, XY{f[0], f[1]})
			maxXY = MaxXY(maxXY, XY{f[0], f[2]})
		}
	}
	minXY.X--
	maxXY.X++

	g := grid{NewMatrix[cell](maxXY.Sub(minXY).Add(Square(1)))}
	for _, c := range g.All() {
		c.ch = '.'
	}
	water = water.Sub(minXY)
	water.Y = max(0, water.Y)

	for _, h := range hs {
		h.fillH(g, minXY)
	}
	for _, v := range vs {
		v.fillH(g.Trans(), minXY.Trans())
	}

	pour(g, water)
	//PrintGrid(g)

	s1, s2 := 0, 0
	for _, c := range g.All() {
		if c.ch == '|' || c.ch == '~' {
			s1++
		}
		if c.ch == '~' {
			s2++
		}
	}
	p.PartOne(s1)
	p.PartTwo(s2)
}

func pour(g grid, src XY) {
	s := &Stack[XY]{}
	s.Push(src)
	for s.Len() > 0 {
		xy := s.Top()
		below := xy.Add(PosY)

		if !g.Dim.HasInside(below) || g.At(below).ch == '|' {
			g.At(xy).ch = '|'
			s.Pop()
			continue
		}

		if g.At(below).ch == '.' {
			g.At(xy).ch = '|'
			s.Push(below)
			continue
		}

		s.Pop()
		spill(g, s, xy)
	}
}

func spill(g grid, s *Stack[XY], xy XY) {
	xl, el := spillDir(g, s, xy.X, xy.Y, -1)
	xr, er := spillDir(g, s, xy.X, xy.Y, +1)
	ch := byte('~')
	if el || er {
		ch = '|'
	}
	for x := xl + 1; x < xr; x++ {
		g.AtXY(x, xy.Y).ch = ch
	}
}

func spillDir(g grid, s *Stack[XY], x, y, d int) (int, bool) {
	for {
		x += d
		if !g.Dim.HasInsideX(x) {
			return x, true
		}
		if g.AtXY(x, y).ch == '#' {
			return x, false
		}
		below := XY{X: x, Y: y + 1}
		if !g.Dim.HasInside(below) || g.At(below).ch == '.' || g.At(below).ch == '|' {
			s.Push(XY{X: x, Y: y})
			return x, true
		}
	}
}
