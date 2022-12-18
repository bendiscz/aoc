package main

import (
	"math"
	"sort"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2022
	day     = 15
	example = `

Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3

`
)

func main() {
	Run(year, day, example, solve)
}

type bound struct {
	x, d int
}

func solve(p *Problem) {
	y0, d0 := 2000000, 4000000
	if p.Example() {
		y0, d0 = 10, 20
	}

	bounds := make([][]bound, d0)
	points := map[int]struct{}{}

	for p.NextLine() {
		m := p.Parse(`Sensor at x=(\d+), y=(\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
		sx, sy, bx, by := ParseInt(m[1]), ParseInt(m[2]), ParseInt(m[3]), ParseInt(m[4])
		d := Abs(bx-sx) + Abs(by-sy)

		for y, my := Max(0, sy-d), Min(d0, sy+d+1); y < my; y++ {
			w := 2*(d-Abs(sy-y)) + 1
			bounds[y] = append(bounds[y], bound{sx - w/2, 1}, bound{sx + w/2 + 1, -1})
		}

		if sy == y0 {
			points[sx] = struct{}{}
		}
		if by == y0 {
			points[bx] = struct{}{}
		}
	}

	for _, b := range bounds {
		sort.Slice(b, func(i, j int) bool { return b[i].x < b[j].x })
	}

	p.PartOne(count(bounds[y0]) - len(points))

	for y := 0; y < d0; y++ {
		x, ok := search(bounds[y], d0)
		if ok {
			p.PartTwo(4000000*x + y)
			break
		}
	}
}

func count(bounds []bound) int {
	sum, x, d := 0, math.MinInt, 0
	for _, b := range bounds {
		if d > 0 {
			sum += b.x - x
		}
		d += b.d
		x = b.x
	}
	return sum
}

func search(bounds []bound, d0 int) (int, bool) {
	x, d := 0, 0
	for _, b := range bounds {
		if b.x > d0 {
			break
		}
		if b.x >= 0 && d == 0 && b.x-x > 0 {
			return x, true
		}
		d += b.d
		x = b.x
	}
	return 0, false
}
