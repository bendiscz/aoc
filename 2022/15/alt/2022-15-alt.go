package main

import (
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

type sensor struct {
	c XY
	d int
}

func (s sensor) vertices() [4]XY {
	d := s.d + 1
	return [...]XY{
		{s.c.X + d, s.c.Y},
		{s.c.X, s.c.Y + d},
		{s.c.X - d, s.c.Y},
		{s.c.X, s.c.Y - d},
	}
}

func (s sensor) contains(c XY) bool {
	d := s.c.Sub(c)
	return Abs(d.X)+Abs(d.Y) <= s.d
}

func intersect(s1, s2 sensor) []XY {
	s := make([]XY, 0, 2)
	v1, v2 := s1.vertices(), s2.vertices()
	for _, i := range [...][4]int{
		{1, 0, 2, 1},
		{1, 0, 3, 0},
		{2, 3, 2, 1},
		{2, 3, 3, 0},
	} {
		if c, ok := intersectLines(v1[i[0]], v1[i[1]], v2[i[2]], v2[i[3]]); ok {
			s = append(s, c)
		}
		if c, ok := intersectLines(v2[i[0]], v2[i[1]], v1[i[2]], v1[i[3]]); ok {
			s = append(s, c)
		}
	}
	return s
}

func intersectLines(a1, a2, b1, b2 XY) (XY, bool) {
	// coordinates: x points right, y points up
	// a1...a2 is from top left to bottom right
	// b1...b2 is from bottom left to top right
	// ---
	// a.x + a.y = c1
	// b.x - b.y = c2
	// ---
	// x = (c1 + c2) / 2
	// y = c1 - x
	// ---
	// c1 = a.x1 + a.y1
	// c2 = b.x1 - b.y1
	x := (a1.X + a1.Y + b1.X - b1.Y) / 2
	if x < a1.X || x > a2.X || x < b1.X || x > b2.X {
		return XY{}, false
	}

	y := a1.X + a1.Y - x
	return XY{x, y}, true
}

func solve(p *Problem) {
	d0 := 4000000
	if p.Example() {
		d0 = 20
	}

	sensors := []sensor(nil)

	for p.NextLine() {
		m := p.Parse(`Sensor at x=(\d+), y=(\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
		sx, sy, bx, by := ParseInt(m[1]), ParseInt(m[2]), ParseInt(m[3]), ParseInt(m[4])
		d := Abs(bx-sx) + Abs(by-sy)
		sensors = append(sensors, sensor{XY{sx, sy}, d})
	}

	intersections := map[XY]int{}
	for i := 0; i < len(sensors); i++ {
		for j := i + 1; j < len(sensors); j++ {
		loop:
			for _, s := range intersect(sensors[i], sensors[j]) {
				if s.X < 0 || s.X > d0 || s.Y < 0 || s.Y > d0 {
					continue
				}

				for k, s2 := range sensors {
					if k != i && k != j && s2.contains(s) {
						continue loop
					}
				}

				intersections[s]++
				if intersections[s] == 4 {
					p.PartTwo(4000000*s.X + s.Y)
					return
				}
			}
		}
	}
}
