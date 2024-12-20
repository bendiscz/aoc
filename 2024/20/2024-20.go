package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2024
	day     = 20
	example = `

###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	ch byte
	l  [2]int
}

func (c cell) String() string { return string([]byte{c.ch}) }

type grid struct {
	*Matrix[cell]
}

func solve(p *Problem) {
	g := grid{NewMatrix[cell](Rectangle(len(p.PeekLine()), 0))}
	start, end := XY0, XY0
	for p.NextLine() {
		s := p.Line()
		for x := 0; x < g.Dim.X; x++ {
			if s[x] == 'S' {
				start = XY{X: x, Y: g.Dim.Y}
			}
			if s[x] == 'E' {
				end = XY{X: x, Y: g.Dim.Y}
			}
		}
		ParseVectorFunc(g.AppendRow(), s, func(b byte) cell { return cell{ch: b} })
	}

	best := findPath(g, 0, start, end)
	findPath(g, 1, end, start)

	save := 100
	if p.Example() {
		save = 64
	}
	p.PartOne(countCheats(g, best, save, 2))
	p.PartTwo(countCheats(g, best, save, 20))
}

func findPath(g grid, index int, start, end XY) int {
	g.At(start).l[index] = 1
	q := Queue[XY]{}
	q.Push(start)
	for q.Len() > 0 {
		p := q.Pop()
		c := g.At(p)

		if p == end {
			return c.l[index] - 1
		}

		for _, d := range HVDirs {
			p2 := p.Add(d)
			c2 := g.At(p2)
			if c2.ch != '#' && c2.l[index] == 0 {
				c2.l[index] = c.l[index] + 1
				q.Push(p2)
			}
		}
	}
	panic("no path found")
}

func countCheats(g grid, best, save, cheat int) int {
	s := 0
	for p, c := range g.All() {
		if c.l[0] == 0 {
			continue
		}

		for y := p.Y - cheat; y <= p.Y+cheat; y++ {
			if !g.Dim.HasInsideY(y) {
				continue
			}
			w := cheat - Abs(y-p.Y)
			for x := p.X - w; x <= p.X+w; x++ {
				if !g.Dim.HasInsideX(x) {
					continue
				}

				p2 := XY{X: x, Y: y}
				c2, d := g.At(p2), p2.Man(p)
				if d < 2 || c2.l[1] == 0 {
					continue
				}

				l := c2.l[1] + c.l[0] + d - 2
				if l+save <= best {
					s++
				}
			}
		}
	}
	return s
}
