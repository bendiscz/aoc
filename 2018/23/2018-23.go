package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"math"
	"slices"
	"sort"
)

const (
	year    = 2018
	day     = 23
	example = `

pos=<10,12,12>, r=2
pos=<12,14,12>, r=2
pos=<16,12,12>, r=4
pos=<14,14,14>, r=6
pos=<50,50,50>, r=200
pos=<10,10,10>, r=5
`
)

func main() {
	Run(year, day, example, solve)
}

type xyz struct{ x, y, z int }

func dist(c1, c2 xyz) int {
	return Abs(c1.x-c2.x) + Abs(c1.y-c2.y) + Abs(c1.z-c2.z)
}

type bot struct {
	pos xyz
	rad int
}

func (b bot) intersects(b2 bot) bool {
	return dist(b.pos, b2.pos) <= b.rad+b2.rad
}

func (b bot) boundary() octahedron {
	o := octahedron{}
	for i, n := range normals {
		o.d[i] = b.rad + n.x*b.pos.x + n.y*b.pos.y + n.z*b.pos.z
	}
	return o
}

// octahedron normals
// [i] and [i+4] are opposite
var normals = [8]xyz{
	{+1, +1, +1},
	{-1, +1, +1},
	{+1, -1, +1},
	{-1, -1, +1},
	{-1, -1, -1},
	{+1, -1, -1},
	{-1, +1, -1},
	{+1, +1, -1},
}

type octahedron struct {
	// together with corresponding normals [a,b,c] defines eight sides as planes
	// ax + by + cz = d
	d [8]int
}

func (o octahedron) intersects(o2 octahedron) bool {
	for i := 0; i < 4; i++ {
		if o.d[i] < -o2.d[i+4] || -o.d[i+4] > o2.d[i] {
			return false
		}
	}
	return true
}

func solve(p *Problem) {
	bs, b0 := []bot(nil), bot{}
	for p.NextLine() {
		b := bot{}
		p.Scanf("pos=<%d,%d,%d>, r=%d", &b.pos.x, &b.pos.y, &b.pos.z, &b.rad)
		bs = append(bs, b)

		if b.rad > b0.rad {
			b0 = b
		}
	}

	s1 := 0
	for _, b := range bs {
		if dist(b.pos, b0.pos) <= b0.rad {
			s1++
		}
	}
	p.PartOne(s1)
	p.PartTwo(solvePart2(bs))
}

func solvePart2(bots []bot) int {
	// 1. use Bron–Kerbosch algorithm to find maximum clique of intersecting bots
	// 2. express bots as octahedra defined as an intersection of eight half-spaces
	// 3. find the ultimate intersection of these octahedra
	// 4. use binary search to find the smallest octahedra at origin that intersects
	//    the ultimate intersection of bots
	bots = maximumClique(bots)
	boundaries := make([]octahedron, len(bots))
	for i, b := range bots {
		boundaries[i] = b.boundary()
	}

	intersection := octahedron{}
	for i := range normals {
		intersection.d[i] = math.MaxInt
		for _, o := range boundaries {
			intersection.d[i] = min(intersection.d[i], o.d[i])
		}
	}

	for r := 1 << 10; ; r <<= 1 {
		if i := sort.Search(r, func(i int) bool {
			return intersection.intersects(bot{rad: i}.boundary())
		}); i < r {
			return i
		}
	}
}

func maximumClique(bots []bot) []bot {
	p := make([]int, len(bots))
	for i := 0; i < len(bots); i++ {
		p[i] = i
	}
	r := bronKerbosch(bots, nil, p, nil, 0)
	clique := []bot(nil)
	for _, i := range r {
		clique = append(clique, bots[i])
	}
	return clique
}

func bronKerbosch(bots []bot, clique, p, x []int, limit int) []int {
	if len(p) == 0 && len(x) == 0 {
		return clique
	}

	if len(p) == 0 || len(p)+len(clique) < limit {
		if len(x) == 0 {
			return clique
		}
		return nil
	}

	maxClique := []int(nil)
	pivot := p[0]
	for _, v := range p {
		if v != pivot && bots[pivot].intersects(bots[v]) {
			continue
		}

		nextClique := append(slices.Clone(clique), v)

		pn := slices.DeleteFunc(slices.Clone(p), func(i int) bool {
			return i == v || !bots[i].intersects(bots[v])
		})

		xn := slices.DeleteFunc(slices.Clone(x), func(i int) bool {
			return i == v || !bots[i].intersects(bots[v])
		})

		if c := bronKerbosch(bots, nextClique, pn, xn, len(maxClique)); len(c) > len(maxClique) {
			maxClique = c
		}

		p = p[1:]
		x = append(x, v)
	}
	return maxClique
}
