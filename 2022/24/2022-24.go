package main

import (
	"fmt"
	"math/bits"
	"strconv"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2022
	day     = 24
	example = `

#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#

`
)

func main() {
	Run(year, day, example, solve)
}

const (
	N = 1 << 0
	S = 1 << 1
	W = 1 << 2
	E = 1 << 3
)

type cell struct {
	wind byte
	next byte
}

func (c cell) String() string {
	switch c.wind {
	case N:
		return "^"
	case S:
		return "v"
	case W:
		return "<"
	case E:
		return ">"
	case 0:
		return "."
	default:
		return strconv.Itoa(bits.OnesCount(uint(c.wind)))
	}
}

type grid struct {
	*Matrix[cell]
}

func (g grid) mod(c XY) XY {
	return XY{(c.X + g.Dim.X) % g.Dim.X, (c.Y + g.Dim.Y) % g.Dim.Y}
}

func (g grid) tick() {
	for x := 0; x < g.Dim.X; x++ {
		for y := 0; y < g.Dim.Y; y++ {
			c := XY{x, y}
			if c2 := g.mod(c.Add(XY{0, 1})); g.Data[c2.Y][c2.X].wind&N != 0 {
				g.Data[c.Y][c.X].next |= N
			}
			if c2 := g.mod(c.Add(XY{0, -1})); g.Data[c2.Y][c2.X].wind&S != 0 {
				g.Data[c.Y][c.X].next |= S
			}
			if c2 := g.mod(c.Add(XY{1, 0})); g.Data[c2.Y][c2.X].wind&W != 0 {
				g.Data[c.Y][c.X].next |= W
			}
			if c2 := g.mod(c.Add(XY{-1, 0})); g.Data[c2.Y][c2.X].wind&E != 0 {
				g.Data[c.Y][c.X].next |= E
			}
		}
	}

	for x := 0; x < g.Dim.X; x++ {
		for y := 0; y < g.Dim.Y; y++ {
			e := g.Data[y][x]
			g.Data[y][x] = cell{e.next, 0}
		}
	}
}

type step struct {
	c    XY
	n    int
	s, e bool
}

func (g grid) search() (int, int) {
	start, end := XY{0, -1}, XY{g.Dim.X - 1, g.Dim.Y}
	ns, n := 0, -1
	v := map[step]bool{}
	q := Queue[step]{}
	q.Push(step{start, 0, false, false})

	for q.Len() > 0 {
		s := q.Pop()
		if v[s] {
			continue
		}
		v[s] = true

		if s.c == start && s.e {
			s.s = true
		}
		if s.c == end {
			if s.s {
				return ns, n
			}
			if ns == 0 {
				ns = n
			}
			s.e = true
		}

		for n < s.n {
			g.tick()
			//g.print()
			n++
		}

		for _, d := range [...]XY{{0, 0}, {0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
			c := s.c.Add(d)
			if c == start || c == end || g.Dim.HasInside(c) && g.Data[c.Y][c.X].wind == 0 {
				q.Push(step{c, s.n + 1, s.s, s.e})
			}
		}
	}

	return 0, 0
}

func (g grid) print() {
	PrintGrid[cell](g)
	fmt.Printf("---\n")
}

func solve(p *Problem) {
	p.NextLine()
	g := grid{NewMatrix[cell](Rectangle(len(p.Line())-2, 0))}
	for p.NextLine() && p.Line()[1] != '#' {
		line := p.Line()
		ParseVectorMap(g.AppendRow(), line[1:len(line)-1], map[byte]cell{
			'^': {wind: N},
			'v': {wind: S},
			'>': {wind: E},
			'<': {wind: W},
			'.': {},
		})
	}

	n1, n2 := g.search()
	p.PartOne(n1)

	//os.Exit(1)

	p.PartTwo(n2)
}
