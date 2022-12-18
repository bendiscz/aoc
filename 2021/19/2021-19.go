package main

import (
	_ "embed"
	"math"
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year = 2021
	day  = 19
)

//go:embed example
var example2 string

func main() {
	Run(year, day, example2, solve)
}

type xyz [3]int

var rotations = [...]xyz{
	// facing ±X
	{1, 2, 3},
	{1, -3, 2},
	{1, -2, -3},
	{1, 3, -2},
	{-1, -2, 3},
	{-1, 3, 2},
	{-1, 2, -3},
	{-1, -3, -2},
	// facing ±Y
	{2, 3, 1},
	{2, -1, 3},
	{2, -3, -1},
	{2, 1, -3},
	{-2, -3, 1},
	{-2, 1, 3},
	{-2, 3, -1},
	{-2, -1, -3},
	// facing ±Z
	{3, 1, 2},
	{3, -2, 1},
	{3, -1, -2},
	{3, 2, -1},
	{-3, -1, 2},
	{-3, 2, 1},
	{-3, 1, -2},
	{-3, -2, -1},
}

func (p xyz) rotate(r xyz) xyz {
	var p2 xyz
	for i, a := range r {
		if a < 0 {
			p[i] = -p[i]
		}
		p2[Abs(a)-1] = p[i]
	}
	return p2
}

func (p xyz) dist(p2 xyz) xyz {
	return xyz{
		p2[0] - p[0],
		p2[1] - p[1],
		p2[2] - p[2],
	}
}

func (p xyz) add(p2 xyz) xyz {
	return xyz{
		p2[0] + p[0],
		p2[1] + p[1],
		p2[2] + p[2],
	}
}

func (p xyz) manhattan(p2 xyz) int {
	return Abs(p2[0]-p[0]) + Abs(p2[1]-p[1]) + Abs(p2[2]-p[2])
}

type box struct {
	index  int
	points []xyz
	linked bool
	offset xyz
}

func (b *box) rotate(r xyz) {
	for i := 0; i < len(b.points); i++ {
		b.points[i] = b.points[i].rotate(r)
	}
}

func match(b0, b1 *box) (offset xyz, rotation xyz, ok bool) {
	for _, rotation = range rotations {
		m := map[xyz]int{}
		for _, p0 := range b0.points {
			for _, p1 := range b1.points {
				p1 = p1.rotate(rotation)
				d := p1.dist(p0)
				c := m[d] + 1
				if c == 12 {
					return d, rotation, true
				}
				m[d] = c
			}
		}
	}
	return
}

func link(boxes []*box, index int) {
	b0 := boxes[index]
	for i := 0; i < len(boxes); i++ {
		b := boxes[i]
		if b.linked {
			continue
		}
		if offset, rotation, ok := match(b0, b); ok {
			b.rotate(rotation)
			b.offset = b0.offset.add(offset)
			b.linked = true
			link(boxes, i)
		}
	}
}

func solve(p *Problem) {
	var boxes []*box
	for p.NextLine() {
		m := p.Parse(`^--- scanner (\d+) ---$`)
		b := box{index: ParseInt(m[1])}
		for p.NextLine() {
			if len(p.Line()) == 0 {
				break
			}
			v := strings.Split(p.Line(), ",")
			b.points = append(b.points, xyz{ParseInt(v[0]), ParseInt(v[1]), ParseInt(v[2])})
		}
		boxes = append(boxes, &b)
	}

	boxes[0].linked = true
	link(boxes, 0)

	beacons := map[xyz]struct{}{}
	for _, b := range boxes {
		for _, p := range b.points {
			beacons[p.add(b.offset)] = struct{}{}
		}
	}

	max := math.MinInt
	for i := 0; i < len(boxes); i++ {
		for j := 0; j < len(boxes); j++ {
			d := boxes[i].offset.manhattan(boxes[j].offset)
			if d > max {
				max = d
			}
		}
	}

	p.PartOne(len(beacons))
	p.PartTwo(max)
}
