package main

import (
	"math"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2018
	day     = 6
	example = `

1, 1
1, 6
8, 3
3, 4
5, 5
8, 9

`
)

func main() {
	Run(year, day, example, solve)
}

type point struct {
	XY
	area int
	inf  bool
	ch   byte
}

func solve(p *Problem) {
	points := []*point(nil)
	v0, v1 := XY{math.MaxInt, math.MaxInt}, XY{math.MinInt, math.MinInt}
	for i := 0; p.NextLine(); i++ {
		f := ParseInts(p.Line())
		pt := point{XY: XY{f[0], f[1]}, area: 0, ch: byte('A' + i)}
		points = append(points, &pt)

		v0 = XY{min(v0.X, pt.X), min(v0.Y, pt.Y)}
		v1 = XY{max(v1.X, pt.X), max(v1.Y, pt.Y)}
	}

	v1 = v1.Sub(v0).Add(XY{1, 1})

	limit2 := 10_000
	if p.Example() {
		limit2 = 32
	}

	margin := Square(limit2/len(points) + 1)
	v0 = v0.Sub(margin)
	v1 = v1.Add(margin).Add(margin)

	count2 := 0
	v1.ForEach(func(xy XY) {
		xy = xy.Add(v0)

		sum, pt0, d0 := 0, (*point)(nil), math.MaxInt
		for _, pt := range points {
			d := dist(pt.XY, xy)
			if d < d0 {
				d0 = d
				pt0 = pt
			} else if d == d0 {
				pt0 = nil
			}

			sum += d
		}

		if pt0 != nil {
			pt0.area++

			if xy.X == v0.X || xy.Y == v0.Y || xy.X == v0.X+v1.X-1 || xy.Y == v0.Y+v1.Y-1 {
				pt0.inf = true
			}
		}

		if sum < limit2 {
			count2++
		}
	})

	r1 := &point{area: 0}
	for _, pt := range points {
		if !pt.inf && pt.area > r1.area {
			r1 = pt
		}
	}

	p.PartOne(r1.area)
	p.PartTwo(count2)
}

func dist(a, b XY) int {
	d := a.Sub(b)
	return Abs(d.X) + Abs(d.Y)
}
