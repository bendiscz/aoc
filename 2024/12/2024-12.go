package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2024
	day     = 12
	example = `

AAAAAAAA
AACBBDDA
AACBBAAA
ABBAAAAA
ABBADDDA
AAAADADA
AAAAAAAA

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	ch byte
	v  bool
}

type grid struct {
	*Matrix[cell]
}

func solve(p *Problem) {
	g0 := &grid{NewMatrix[cell](Rectangle(len(p.PeekLine()), 0))}
	for p.NextLine() {
		ParseVectorFunc(g0.AppendRow(), p.Line(), func(b byte) cell { return cell{ch: b} })
	}

	s1, s2 := 0, 0
	g := &grid{NewMatrix[cell](g0.Dim.Add(Square(2)))}
	CopyGrid(g.SubGrid(XY{1, 1}, g0.Dim), g0)

	for c := range g0.Dim.All() {
		p1, p2 := count(g, c.Add(Square(1)))
		s1 += p1
		s2 += p2
	}
	p.PartOne(s1)
	p.PartTwo(s2)
}

type side struct {
	a, b XY
	v    bool
}

func (s side) dir() XY {
	return s.b.Sub(s.a)
}

type perimeter struct {
	sides    []side
	adjacent map[XY][2]int
}

func newPerimeter() *perimeter {
	return &perimeter{
		adjacent: map[XY][2]int{},
	}
}

func (p *perimeter) add(a, b XY) {
	i := len(p.sides)
	p.sides = append(p.sides, side{a: a, b: b})
	if s, ok := p.adjacent[a]; ok {
		s[1] = i
		p.adjacent[a] = s
	} else {
		p.adjacent[a] = [2]int{i, -1}
	}
}

func (p *perimeter) findNext(b, d XY) int {
	a := p.adjacent[b]
	if a[1] == -1 {
		return a[0]
	}
	d = XY{-d.Y, d.X}
	if d == p.sides[a[0]].dir() {
		return a[0]
	} else {
		return a[1]
	}
}

func (p *perimeter) count() int {
	sum, r := 0, len(p.sides)
	for i := 0; r > 0; i = (i + 1) % len(p.sides) {
		if p.sides[i].v {
			r--
			continue
		}

		p.sides[i].v = true
		d := p.sides[i].dir()
		j := p.findNext(p.sides[i].b, d)
		for {
			p.sides[j].v = true
			d2 := p.sides[j].dir()
			if d2 != d {
				sum++
				d = d2
			}

			if j == i {
				break
			}
			j = p.findNext(p.sides[j].b, d)
		}
	}
	return sum
}

func count(g *grid, c XY) (int, int) {
	if g.At(c).v {
		return 0, 0
	}

	ch, area := g.At(c).ch, 0
	p := newPerimeter()
	q := Queue[XY]{}
	q.Push(c)

	for q.Len() > 0 {
		c = q.Pop()
		if g.At(c).v {
			continue
		}

		area++
		g.At(c).v = true

		for _, d := range HVDirs {
			c2 := c.Add(d)
			nc := g.At(c2)

			if nc.ch == ch && !nc.v {
				q.Push(c2)
			}
			if nc.ch != ch {
				switch d {
				case NegY:
					p.add(c, c.Add(XY{1, 0}))
				case PosX:
					p.add(c.Add(XY{1, 0}), c.Add(XY{1, 1}))
				case PosY:
					p.add(c.Add(XY{1, 1}), c.Add(XY{0, 1}))
				case NegX:
					p.add(c.Add(XY{0, 1}), c)
				}
			}
		}
	}

	return area * len(p.sides), area * p.count()
}
