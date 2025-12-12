package main

import (
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2025
	day     = 5
	example = `

3-5
10-14
16-20
12-18

1
5
8
11
17
32

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	s1, s2 := 0, 0

	ints := [][]int(nil)
	for p.NextLine() {
		if p.Line() == "" {
			break
		}

		f := ParseInts(strings.ReplaceAll(p.Line(), "-", "_"))
		ints = append(ints, f)
	}

	for p.NextLine() {
		x := ParseInt(p.Line())
		for _, i := range ints {
			if x >= i[0] && x <= i[1] {
				s1++
				break
			}
		}
	}

	set := IntervalSet[int](nil)
	for _, i := range ints {
		set = append(set, Interval[int]{X1: i[0], X2: i[1] + 1})
	}

	set = MergeIntervals[int](set...)

	for _, i := range set {
		s2 += i.X2 - i.X1
	}

	p.PartOne(s1)
	p.PartTwo(s2)
}
