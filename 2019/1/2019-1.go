package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2019
	day     = 1
	example = `

12
14
1969
100756

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	s1, s2 := 0, 0
	for p.NextLine() {
		f := ParseInt(p.Line())/3 - 2
		s1 += f
		s2 += f
		for f > 0 {
			f = max(f/3-2, 0)
			s2 += f
		}
	}
	p.PartOne(s1)
	p.PartTwo(s2)
}
