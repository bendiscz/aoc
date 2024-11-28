package main

import (
	"math"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 3
	example = `

1
12
23
26
1024

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	for p.NextLine() {
		x := ParseInt(p.Line())
		p.PartOne(part1(x))
		p.PartTwo(part2(x))
	}
}

func dist(x int) (d, s, y int) {
	f := math.Sqrt(float64(x))
	d = int(f)
	if float64(d) < f {
		d++
	}
	if d%2 == 0 {
		d++
	}

	s = (d*d - (d-2)*(d-2)) / 4
	if s == 0 {
		s = 1
	}

	y = Abs((d*d-x)%s - s/2)
	return
}

func part1(x int) int {
	d, _, y := dist(x)
	return d/2 + y
}

func part2(x int) int {
	n := map[XY]int{}
	dirs, d, l, s := [...]XY{PosX, PosY, NegX, NegY}, 0, 2, 1
	edges := [...]XY{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	xy := XY{0, 0}

	for i := 0; ; i++ {
		sum := 0
		if xy == (XY{0, 0}) {
			sum = 1
		}
		for _, edge := range edges {
			if e, ok := n[xy.Add(edge)]; ok {
				sum += e
			}
		}

		if sum > x {
			return sum
		}

		n[xy] = sum
		sum = 0

		s--
		if s == 0 {
			s = l / 2
			d++
			l++
		}
		xy = xy.Add(dirs[d%4])
	}
}
