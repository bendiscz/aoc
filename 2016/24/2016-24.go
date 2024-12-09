package main

import (
	"math"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2016
	day     = 24
	example = `

###########
#0.1.....2#
#.#######.#
#4.......3#
###########

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	b byte
}

type grid struct {
	*Matrix[cell]
}

func solve(p *Problem) {
	g, points := grid{}, map[int]XY{}
	for p.NextLine() {
		if g.Matrix == nil {
			g.Matrix = NewMatrix[cell](Rectangle(len(p.Line()), 0))
		}
		for x, b := range p.LineBytes() {
			if b >= '0' && b <= '9' {
				points[int(b-'0')] = XY{x, g.Dim.Y}
			}
		}
		ParseVectorFunc(g.AppendRow(), p.Line(), func(b byte) cell { return cell{b} })
	}

	n := len(points)
	dist := NewMatrix[int](Square(n))
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d := shortestPath(g, points[i], points[j])
			*dist.AtXY(i, j) = d
			*dist.AtXY(j, i) = d
		}
	}

	path := make([]int, len(points)-1)
	for i := range path {
		path[i] = i + 1
	}
	best1, best2 := math.MaxInt, math.MaxInt
	for pp := range Permutations(path) {
		d := *dist.AtXY(0, pp[0])
		for i := 1; i < len(pp); i++ {
			d += *dist.AtXY(pp[i-1], pp[i])
		}

		best1 = min(best1, d)
		best2 = min(best2, d+*dist.AtXY(0, pp[len(path)-1]))
	}
	p.PartOne(best1)
	p.PartTwo(best2)
}

func shortestPath(g grid, from, to XY) int {
	type path struct {
		xy XY
		d  int
	}

	q, v := Queue[path]{}, map[XY]bool{}
	q.Push(path{from, 0})

	for {
		p := q.Pop()
		if p.xy == to {
			return p.d
		}

		if g.At(p.xy).b == '#' || v[p.xy] {
			continue
		}
		v[p.xy] = true

		for _, d := range HVDirs {
			q.Push(path{p.xy.Add(d), p.d + 1})
		}
	}
}
