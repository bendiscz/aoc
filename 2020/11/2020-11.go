package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2020
	day     = 11
	example = `

L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	seat bool
	full bool
	next bool
}

func (c cell) String() string {
	switch {
	case c.full:
		return "#"
	case c.seat:
		return "L"
	default:
		return "."
	}
}

type grid struct {
	*Matrix[cell]
}

func solve(p *Problem) {
	g := &grid{NewMatrix[cell](Rectangle(len(p.PeekLine()), 0))}
	for p.NextLine() {
		ParseVectorFunc(g.AppendRow(), p.Line(), func(b byte) cell { return cell{seat: b == 'L'} })
	}

	for {
		if s, stable := g.step(); stable {
			p.PartOne(s)
			break
		}
	}

	g.reset()
	for {
		if s, stable := g.step2(); stable {
			p.PartOne(s)
			break
		}
	}
}

func (g *grid) reset() {
	g.Dim.ForEach(func(xy XY) {
		g.At(xy).full = false
	})
}

func (g *grid) count(p XY) int {
	s := 0
	for _, d := range AllDirs {
		p2 := p.Add(d)
		if g.Dim.HasInside(p2) && g.At(p2).full {
			s++
		}
	}
	return s
}

func (g *grid) step() (int, bool) {
	g.Dim.ForEach(func(xy XY) {
		if !g.At(xy).seat {
			return
		}

		s := g.count(xy)
		switch {
		case s == 0:
			g.At(xy).next = true
		case s >= 4:
			g.At(xy).next = false
		}
	})

	s, stable := 0, true
	g.Dim.ForEach(func(xy XY) {
		c := g.At(xy)
		if c.full {
			s++
		}
		if c.next != c.full {
			c.full = c.next
			stable = false
		}
	})
	return s, stable
}

func (g *grid) count2(p XY) int {
	s := 0
	for _, d := range AllDirs {
		p2 := p
		for {
			p2 = p2.Add(d)
			if !g.Dim.HasInside(p2) {
				break
			}
			if !g.At(p2).seat {
				continue
			}
			if g.At(p2).full {
				s++
			}
			break
		}
	}
	return s
}

func (g *grid) step2() (int, bool) {
	g.Dim.ForEach(func(xy XY) {
		if !g.At(xy).seat {
			return
		}

		s := g.count2(xy)
		switch {
		case s == 0:
			g.At(xy).next = true
		case s >= 5:
			g.At(xy).next = false
		default:
			g.At(xy).next = g.At(xy).full
		}
	})

	s, stable := 0, true
	g.Dim.ForEach(func(xy XY) {
		c := g.At(xy)
		if c.full {
			s++
		}
		if c.next != c.full {
			c.full = c.next
			stable = false
		}
	})
	return s, stable
}
