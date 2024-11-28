package main

import (
	"math/bits"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year = 2016
	day  = 13
)

func main() {
	Run(year, day, "", solve)
}

var magic int
var walls map[cell]bool

type cell struct {
	XY
}

func (c cell) wall() bool {
	if w, ok := walls[c]; ok {
		return w
	}

	z := c.X*(c.X+3) + 2*c.X*c.Y + c.Y*(c.Y+1)
	z += magic
	w := bits.OnesCount(uint(z))%2 == 1

	walls[c] = w
	return w
}

func (c cell) Key() any {
	return c
}

func (c cell) Edges() []Edge {
	e := []Edge(nil)
	for _, d := range HVDirs {
		a := cell{c.Add(d)}
		if a.X >= 0 && a.Y >= 0 && !a.wall() {
			e = append(e, Edge{V: a, W: 1})
		}
	}
	return e
}

type step struct {
	c cell
	d int
}

func solve(p *Problem) {
	p.NextLine()
	magic = ParseInt(p.Line())
	walls = map[cell]bool{}

	start, finish := cell{XY{1, 1}}, cell{XY{31, 39}}

	path := ShortestPath(start, finish)
	p.PartOne(path.Cost)

	visited := map[cell]struct{}{}
	q := Queue[step]{}

	q.Push(step{start, 0})
	visited[start] = struct{}{}

	for q.Len() > 0 {
		s := q.Pop()
		if s.d == 50 {
			continue
		}

		for _, e := range s.c.Edges() {
			c := e.V.(cell)
			if _, ok := visited[c]; ok {
				continue
			}
			visited[c] = struct{}{}
			q.Push(step{c, s.d + 1})
		}
	}
	p.PartTwo(len(visited))
}
