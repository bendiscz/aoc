package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"slices"
)

const (
	year    = 2018
	day     = 25
	example = `

 0,0,0,0
 3,0,0,0
 0,3,0,0
 0,0,3,0
 0,0,0,3
 0,0,0,6
 9,0,0,0
12,0,0,0

`
)

func main() {
	Run(year, day, example, solve)
}

type point struct {
	x, y, z, t int
	marked     bool
}

func connected(p1, p2 *point) bool {
	return Abs(p1.x-p2.x)+Abs(p1.y-p2.y)+Abs(p1.z-p2.z)+Abs(p1.t-p2.t) <= 3
}

func solve(p *Problem) {
	points := []*point(nil)
	for p.NextLine() {
		f := ParseInts(p.Line())
		points = append(points, &point{x: f[0], y: f[1], z: f[2], t: f[3]})
	}

	s1 := 0
	for len(points) > 0 {
		findConstellation(points)
		points = slices.DeleteFunc(points, func(p *point) bool { return p.marked })
		s1++
	}
	p.PartOne(s1)

	p.PartTwo("Merry Christmas!")
}

func findConstellation(points []*point) {
	q := Queue[*point]{}
	q.Push(points[0])
	for q.Len() > 0 {
		p := q.Pop()
		for _, p2 := range points {
			if !p2.marked && connected(p, p2) {
				p2.marked = true
				q.Push(p2)
			}
		}
	}
}
