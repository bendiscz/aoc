package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"slices"
)

const (
	year    = 2024
	day     = 1
	example = `

3   4
4   3
2   5
1   3
3   9
3   3

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	var l1, l2 []int
	f2 := make(map[int]int)
	for p.NextLine() {
		n := ParseInts(p.Line())
		l1 = append(l1, n[0])
		l2 = append(l2, n[1])
		f2[n[1]]++
	}

	slices.Sort(l1)
	slices.Sort(l2)

	d, s := 0, 0
	for i := 0; i < len(l1); i++ {
		d += Abs(l1[i] - l2[i])
		s += f2[l1[i]] * l1[i]
	}

	p.PartOne(d)
	p.PartTwo(s)
}
