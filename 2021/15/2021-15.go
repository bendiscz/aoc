package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2021
	day     = 15
	example = `

1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581

`
)

func main() {
	Run(year, day, example, solve)
}

// TODO matrix???
type grid struct {
	d int
	w [][]int
}

type point struct {
	XY
	g *grid
}

func (p point) Key() any { return p.XY }

func (p point) Edges() []Edge {
	var edges []Edge
	add := func(x, y int) {
		edges = append(edges, Edge{V: point{
			XY: XY{x, y},
			g:  p.g,
		}, W: p.g.w[x][y]})
	}

	if p.X > 0 {
		add(p.X-1, p.Y)
	}
	if p.X < p.g.d-1 {
		add(p.X+1, p.Y)
	}
	if p.Y > 0 {
		add(p.X, p.Y-1)
	}
	if p.Y < p.g.d-1 {
		add(p.X, p.Y+1)
	}

	return edges
}

func newGrid(d int) *grid {
	g := grid{
		d: d,
		w: make([][]int, d),
	}
	for i := range g.w {
		g.w[i] = make([]int, d)
	}
	return &g
}

func solve(p *Problem) {
	var g *grid
	y := 0
	for p.NextLine() {
		if g == nil {
			g = newGrid(len(p.Line()))
		}
		for x, ch := range []byte(p.Line()) {
			g.w[x][y] = int(ch - '0')
		}
		y++
	}

	path1 := ShortestPath(point{XY{0, 0}, g}, point{XY{g.d - 1, g.d - 1}, g})
	p.PartOne(path1.Cost)
	p.Printf("path: %+v", path1.Steps)

	eg := newGrid(g.d * 5)
	for x0 := 0; x0 < 5; x0++ {
		for y0 := 0; y0 < 5; y0++ {
			for x := 0; x < g.d; x++ {
				for y := 0; y < g.d; y++ {
					eg.w[x0*g.d+x][y0*g.d+y] = (g.w[x][y]+x0+y0-1)%9 + 1
				}
			}
		}
	}

	path2 := ShortestPath(point{XY{0, 0}, eg}, point{XY{eg.d - 1, eg.d - 1}, g})
	p.PartTwo(path2.Cost)
}
