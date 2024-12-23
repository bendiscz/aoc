package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2018
	day     = 22
	example = `

depth: 510
target: 10,10

`
)

func main() {
	Run(year, day, example, solve)
}

const (
	xf = 16807
	yf = 48271
	m  = 20183

	torch = 1
	//neither = 0
	//gear    = 2
)

type grid struct {
	depth  int
	target XY
	dim    XY
	level  map[XY]int
}

func newGrid(depth int, target XY) *grid {
	return &grid{
		depth:  depth,
		target: target,
		dim:    Rectangle(target.X+1, 1),
		level:  map[XY]int{},
	}
}

func (g *grid) index(c XY) int {
	switch {
	case c == XY0 || c == g.target:
		return 0
	case c.Y == 0:
		return c.X * xf
	case c.X == 0:
		return c.Y * yf
	default:
		return g.erosion(c.Sub(PosX)) * g.erosion(c.Sub(PosY))
	}
}

func (g *grid) erosion(c XY) int {
	if level, ok := g.level[c]; ok {
		return level
	}

	level := (g.index(c) + g.depth) % m
	g.level[c] = level
	return level
}

func (g *grid) regionType(c XY) int {
	return g.erosion(c) % 3
}

type rescuer struct {
	g    *grid
	at   XY
	tool int
}

func canHandle(tool, terrain int) bool { return tool != terrain }

func (r rescuer) Key() any { return r }

func (r rescuer) Edges() []Edge {
	e := []Edge(nil)
	for _, d := range HVDirs {
		c := r.at.Add(d)
		if c.X < 0 || c.Y < 0 {
			continue
		}
		if canHandle(r.tool, r.g.regionType(c)) {
			e = append(e, Edge{V: rescuer{r.g, c, r.tool}, W: 1})
		}
	}
	terrain := r.g.regionType(r.at)
	for tool := 0; tool < 3; tool++ {
		if tool != r.tool && canHandle(tool, terrain) {
			e = append(e, Edge{V: rescuer{r.g, r.at, tool}, W: 7})
		}
	}
	return e
}

func solve(p *Problem) {
	f := ParseInts(p.ReadAll())
	g := newGrid(f[0], XY{X: f[1], Y: f[2]})

	s1 := 0
	for xy := range g.target.Add(Square(1)).All() {
		s1 += g.regionType(xy)
	}
	p.PartTwo(s1)

	path := ShortestPath(rescuer{g, XY0, torch}, rescuer{g, g.target, torch})
	p.PartTwo(path.Cost)
}
