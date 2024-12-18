package main

import (
	"fmt"
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2024
	day     = 18
	example = `

5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	v bool
}

func (c cell) String() string {
	if c.v {
		return "#"
	} else {
		return "."
	}
}

type grid struct {
	*Matrix[cell]
}

type vertex struct {
	g  grid
	at XY
}

func (v vertex) Key() any {
	return v.at
}

func (v vertex) Edges() []Edge {
	e := []Edge(nil)
	for _, d := range HVDirs {
		xy := v.at.Add(d)
		if !v.g.Dim.HasInside(xy) || v.g.At(xy).v {
			continue
		}
		e = append(e, Edge{
			V: vertex{v.g, xy},
			W: 1,
		})
	}
	return e
}

func solve(p *Problem) {
	N, B := 71, 1024
	if p.Example() {
		N, B = 7, 12
	}

	g := grid{NewMatrix[cell](Square(N))}
	bytes := []XY(nil)
	for p.NextLine() {
		f := ParseInts(p.Line())
		bytes = append(bytes, XY{X: f[0], Y: f[1]})
	}

	for _, b := range bytes[:B] {
		g.At(b).v = true
	}
	start, end := vertex{g, XY0}, vertex{g, Square(N - 1)}
	path := ShortestPath(start, end)
	p.PartOne(len(path.Steps) - 1)

	for _, b := range bytes[B:] {
		g.At(b).v = true
	}
	for i := len(bytes) - 1; i >= B; i-- {
		b := bytes[i]
		g.At(b).v = false
		path = ShortestPath(start, end)
		if len(path.Steps) > 0 {
			p.PartTwo(fmt.Sprintf("%d,%d", b.X, b.Y))
			break
		}
	}
}
