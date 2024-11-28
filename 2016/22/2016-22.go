package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2016
	day     = 22
	example = ``
)

func main() {
	Run(year, day, example, solve)
}

type node struct {
	size    uint
	used    uint
	blocked bool
}

func (n node) String() string {
	if n.used == 0 {
		return "_"
	} else if n.blocked {
		return "#"
	} else {
		return "."
	}
}

type grid struct {
	*Matrix[node]
}

func solve(p *Problem) {
	p.NextLine()
	p.NextLine()

	const xn, yn = 35, 25
	g := grid{NewMatrix[node](Rectangle(xn, yn))}
	empty := XY{}

	for p.NextLine() {
		f := ParseUints(p.Line())
		xy := XY{int(f[0]), int(f[1])}
		*g.At(xy) = node{size: f[2], used: f[3]}

		if f[3] == 0 {
			empty = xy
		}
	}

	count1 := 0
	g.Matrix.Dim.ForEach(func(xy XY) {
		if xy == empty {
			return
		}
		if g.At(xy).used <= g.At(empty).size {
			count1++
		} else {
			g.At(xy).blocked = true
		}
	})
	p.PartOne(count1)

	//PrintGrid[node](g)

	d := findPath(g, empty, XY{g.Dim.X - 2, 0})
	p.PartTwo(d + 1 + 5*(g.Dim.X-2))
}

func findPath(g grid, start, end XY) int {
	type path struct {
		xy   XY
		dist int
	}

	q := Queue[path]{}
	q.Push(path{start, 0})

	v := map[XY]bool{}

	for {
		p := q.Pop()
		if p.xy == end {
			return p.dist
		}

		if g.At(p.xy).blocked {
			continue
		}

		if v[p.xy] {
			continue
		}
		v[p.xy] = true

		for _, d := range HVDirs {
			n := p.xy.Add(d)
			if g.Dim.HasInside(n) {
				q.Push(path{n, p.dist + 1})
			}
		}
	}
}
