package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2020
	day     = 1
	example = `

1721
979
366
299
675
1456

`
)

func main() {
	Run(year, day, example, solve)
}

type set = map[int]struct{}

func solve(p *Problem) {
	s1, s2 := 0, 0

	v := set{}
	for p.NextLine() {
		v[ParseInt(p.Line())] = struct{}{}
	}

	s1, _ = findPair(v, 2020)
	p.PartOne(s1)

	for x := range v {
		if y, ok := findPair(v, 2020-x); ok {
			s2 = x * y
			break
		}
	}
	p.PartTwo(s2)
}

func findPair(v set, sum int) (int, bool) {
	for x := range v {
		y := sum - x
		if _, ok := v[y]; ok {
			return x * y, true
		}
	}
	return 0, false
}
