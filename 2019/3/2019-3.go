package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"math"
)

const (
	year    = 2019
	day     = 3
	example = `

R8,U5,L5,D3
U7,R6,D4,L4

`
)

func main() {
	Run(year, day, example, solve)
}

var dirs = []XY{
	'R': PosX,
	'L': NegX,
	'U': PosY,
	'D': NegY,
}

func solve(p *Problem) {
	wire := map[XY]int{}
	xy, l := XY{}, 0
	p.NextLine()
	for _, f := range SplitFields(p.Line()) {
		for d, n := dirs[f[0]], ParseInt(f[1:]); n > 0; n-- {
			xy = xy.Add(d)
			l++
			wire[xy] = l
		}
	}

	xy, l = XY{}, 0
	d1, d2 := math.MaxInt, math.MaxInt
	p.NextLine()
	for _, f := range SplitFields(p.Line()) {
		for d, n := dirs[f[0]], ParseInt(f[1:]); n > 0; n-- {
			xy = xy.Add(d)
			l++
			if l1, ok := wire[xy]; ok {
				d1 = min(d1, Abs(xy.X)+Abs(xy.Y))
				d2 = min(d2, l1+l)
			}
		}
	}

	p.PartOne(d1)
	p.PartOne(d2)
}

//type line struct {
//	u      int
//	v0, v1 int
//}
//
//func cmpLines(a, b line) int {
//	return cmp.Compare(a.u, b.u)
//}
//
//func findIntersection(lines []line, u1, u2, v int) (int, bool) {
//	i1, _ := slices.BinarySearchFunc(lines, line{u: u1}, cmpLines)
//	i2, _ := slices.BinarySearchFunc(lines, line{u: u2}, cmpLines)
//
//	u0 := math.MaxInt
//	for _, l := range lines[i1:i2] {
//		if v == 0 && l.u == 0 {
//			continue
//		}
//		if v < l.v0 || l.v1 < v {
//			continue
//		}
//		if Abs(l.u) < Abs(u0) {
//			u0 = l.u
//		}
//	}
//
//	return u0, u0 < math.MaxInt
//}
//
//func solve(p *Problem) {
//	hs, vs := []line(nil), []line(nil)
//	x, y := 0, 0
//	s := 0
//	p.NextLine()
//	for _, f := range SplitFields(p.Line()) {
//		ch, d := f[0], ParseInt(f[1:])
//		s += d
//		switch ch {
//		case 'R':
//			hs = append(hs, line{y, x, x + d})
//			x += d
//		case 'L':
//			hs = append(hs, line{y, x - d, x})
//			x -= d
//		case 'U':
//			vs = append(vs, line{x, y, y + d})
//			y += d
//		case 'D':
//			vs = append(vs, line{x, y - d, y})
//			y -= d
//		}
//	}
//
//	p.Printf("%d %d %d", x, y, s)
//	slices.SortFunc(hs, cmpLines)
//	slices.SortFunc(vs, cmpLines)
//
//	p.NextLine()
//	x, y = 0, 0
//	d0 := math.MaxInt
//	for _, f := range SplitFields(p.Line()) {
//		ch, d := f[0], ParseInt(f[1:])
//		switch ch {
//		case 'R':
//			if xa, ok := findIntersection(vs, x, x+d, y); ok {
//				d0 = min(d0, Abs(xa)+Abs(y))
//			}
//			x += d
//		case 'L':
//			if xa, ok := findIntersection(vs, x-d, x, y); ok {
//				d0 = min(d0, Abs(xa)+Abs(y))
//			}
//			x -= d
//		case 'U':
//			if ya, ok := findIntersection(hs, y, y+d, x); ok {
//				d0 = min(d0, Abs(x)+Abs(ya))
//			}
//			y += d
//		case 'D':
//			if ya, ok := findIntersection(hs, y-d, y, x); ok {
//				d0 = min(d0, Abs(x)+Abs(ya))
//			}
//			y -= d
//		}
//	}
//
//	p.PartOne(d0)
//}
