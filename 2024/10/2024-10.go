package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2024
	day     = 10
	example = `

89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	ch byte
}

type grid struct {
	*Matrix[cell]
}

func solve(p *Problem) {
	s1, s2 := 0, 0

	g := grid{NewMatrix[cell](Rectangle(len(p.PeekLine()), 0))}
	for p.NextLine() {
		ParseVectorFunc(g.AppendRow(), p.Line(), func(b byte) cell { return cell{b} })
	}

	for start := range g.Dim.All() {
		if g.At(start).ch == '0' {
			t1, t2 := score(g, start)
			s1 += t1
			s2 += t2
		}
	}

	p.PartOne(s1)
	p.PartTwo(s2)
}

func score(g grid, start XY) (int, int) {
	s, q, r := 0, Queue[XY]{}, Set[XY]{}
	q.Push(start)
	for q.Len() > 0 {
		xy := q.Pop()
		if g.At(xy).ch == '9' {
			r[xy] = SET
			s++
			continue
		}
		for _, d := range HVDirs {
			n := xy.Add(d)
			if !g.Dim.HasInside(n) {
				continue
			}
			if g.At(n).ch-g.At(xy).ch == 1 {
				q.Push(n)
			}
		}
	}
	return len(r), s
}
