package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"math"
)

const (
	year    = 2019
	day     = 24
	example = `

....#
#..#.
#..##
..#..
#....

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

type multigrid struct {
	g map[int]*grid
}

func newMultiGrid() *multigrid {
	return &multigrid{g: map[int]*grid{}}
}

type grid struct {
	*Matrix[cell]
	m *multigrid
	z int
}

func newGrid(m *multigrid, z int) *grid {
	g := &grid{Matrix: NewMatrix[cell](Square(5)), m: m, z: z}
	if m != nil {
		m.g[z] = g
	}
	return g
}

func (g *grid) cell(x, y int) bool {
	return g != nil && g.AtXY(x, y).v
}

func (g *grid) code() int {
	if g == nil {
		return 0
	}

	code, bit := 0, 1
	for y := 0; y < g.Dim.Y; y++ {
		for x := 0; x < g.Dim.X; x++ {
			if g.AtXY(x, y).v {
				code |= bit
			}
			bit <<= 1
		}
	}
	return code
}

func (g *grid) count(xy XY) int {
	s := 0
	for _, d := range HVDirs {
		c := xy.Add(d)
		if g.Dim.HasInside(c) && g.At(c).v {
			s++
		}
	}

	if g.m == nil {
		return s
	}

	if xy.X == 0 {
		if g.m.g[g.z-1].cell(1, 2) {
			s++
		}
	}
	if xy.X == 4 {
		if g.m.g[g.z-1].cell(3, 2) {
			s++
		}
	}
	if xy.Y == 0 {
		if g.m.g[g.z-1].cell(2, 1) {
			s++
		}
	}
	if xy.Y == 4 {
		if g.m.g[g.z-1].cell(2, 3) {
			s++
		}
	}

	if xy.X == 1 && xy.Y == 2 {
		for y := 0; y < 5; y++ {
			if g.m.g[g.z+1].cell(0, y) {
				s++
			}
		}
	}
	if xy.X == 3 && xy.Y == 2 {
		for y := 0; y < 5; y++ {
			if g.m.g[g.z+1].cell(4, y) {
				s++
			}
		}
	}
	if xy.X == 2 && xy.Y == 1 {
		for x := 0; x < 5; x++ {
			if g.m.g[g.z+1].cell(x, 0) {
				s++
			}
		}
	}
	if xy.X == 2 && xy.Y == 3 {
		for x := 0; x < 5; x++ {
			if g.m.g[g.z+1].cell(x, 4) {
				s++
			}
		}
	}

	return s
}

func solve(p *Problem) {
	g0 := newGrid(nil, 0)
	for y := 0; p.NextLine(); y++ {
		ParseVectorFunc(g0.Row(y), p.Line(), func(b byte) cell { return cell{v: b == '#'} })
	}

	g, codes := &grid{Matrix: CloneGrid(g0)}, Set[int]{}
	for {
		code := g.code()
		if codes.Contains(code) {
			p.PartOne(code)
			break
		}
		codes[code] = SET

		ng := newGrid(nil, 0)
		evolve(g, ng)
		g = ng
	}

	m := newMultiGrid()
	CopyGrid(newGrid(m, 0), g0)

	for i := 0; i < 200; i++ {
		if i == 10 && p.Example() {
			break
		}

		nm := newMultiGrid()
		minZ, maxZ := math.MaxInt, math.MinInt
		for z, g := range m.g {
			minZ, maxZ = min(minZ, z), max(maxZ, z)
			ng := newGrid(nm, z)
			evolve(g, ng)
		}

		evolve(newGrid(m, minZ-1), newGrid(nm, minZ-1))
		evolve(newGrid(m, maxZ+1), newGrid(nm, maxZ+1))

		for z := minZ - 1; z < 0 && nm.g[z].code() == 0; z++ {
			delete(nm.g, z)
		}
		for z := maxZ + 1; z > 0 && nm.g[z].code() == 0; z-- {
			delete(nm.g, z)
		}
		m = nm
	}

	s2 := 0
	for _, g := range m.g {
		for _, c := range g.All() {
			if c.v {
				s2++
			}
		}
	}
	p.PartTwo(s2)
}

func evolve(g, ng *grid) {
	for xy, c := range ng.All() {
		if ng.m != nil && xy.X == 2 && xy.Y == 2 {
			continue
		}

		s := g.count(xy)
		v := g.cell(xy.X, xy.Y)
		if v {
			if s != 1 {
				v = false
			}
		} else {
			if s == 1 || s == 2 {
				v = true
			}
		}
		c.v = v
	}
}
