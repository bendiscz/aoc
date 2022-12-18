package main

import (
	"sort"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2022
	day     = 1
	example = `

1000
2000
3000

4000

5000
6000

7000
8000
9000

10000

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	sum := 0
	var sums []int
	for p.NextLine() {
		if p.Line() == "" {
			sums = append(sums, sum)
			sum = 0
			continue
		}
		sum += ParseInt(p.Line())
	}

	sort.Ints(sums)
	l := len(sums)

	p.PartOne(sums[l-1])
	p.PartTwo(sums[l-1] + sums[l-2] + sums[l-3])
}
