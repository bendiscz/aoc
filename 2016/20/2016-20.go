package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2016
	day     = 20
	example = `

5-8
0-2
4-7

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	blacklist := IntervalSet[int](nil)
	for p.NextLine() {
		var x1, x2 int
		p.Scanf("%d-%d", &x1, &x2)
		blacklist = append(blacklist, Interval[int]{x1, x2 + 1})
	}

	blacklist = MergeIntervals[int](blacklist...)
	r1 := 0
	if blacklist[0].X1 == 0 {
		r1 = blacklist[0].X2
	}
	p.PartOne(r1)

	count := 1 << 32
	for _, i := range blacklist {
		count -= i.X2 - i.X1
	}
	p.PartTwo(count)
}
