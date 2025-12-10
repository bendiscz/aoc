package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2025
	day     = 9
	example = `

7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	s1, s2 := 0, 0

	points := []XY(nil)
	for p.NextLine() {
		f := ParseInts(p.Line())
		points = append(points, XY{f[0], f[1]})
	}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1, p2 := points[i], points[j]
			d := p1.Sub(p2)
			area := (Abs(d.X) + 1) * (Abs(d.Y) + 1)
			s1 = max(s1, area)

			if area > s2 && checkInside(p1, p2, points) {
				s2 = area
			}
		}
	}

	p.PartOne(s1)
	p.PartTwo(s2)
}

func checkInside(p1, p2 XY, pts []XY) bool {
	xmin := min(p1.X, p2.X)
	xmax := max(p1.X, p2.X)
	ymin := min(p1.Y, p2.Y)
	ymax := max(p1.Y, p2.Y)

	for i := 0; i < len(pts); i++ {
		l1, l2 := pts[i], pts[(i+1)%len(pts)]
		if intersects(l1, l2, xmin, xmax, ymin, ymax) {
			return false
		}
	}
	return true
}

func intersects(l1, l2 XY, xmin, xmax, ymin, ymax int) bool {
	if l1.X == l2.X {
		x, y1, y2 := l1.X, l1.Y, l2.Y
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		if x <= xmin || x >= xmax {
			return false
		}
		if y2 <= ymin || y1 >= ymax {
			return false
		}
		return true
	} else {
		x1, x2, y := l1.X, l2.X, l1.Y
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		if x2 <= xmin || x1 >= xmax {
			return false
		}
		if y <= ymin || y >= ymax {
			return false
		}
		return true
	}
}
