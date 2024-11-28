package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 5
	example = `

0
3
0
1
-3

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	input := p.ReadAll()
	p.PartOne(solvePart(input, false))
	p.PartTwo(solvePart(input, true))
}

func solvePart(s string, part2 bool) int {
	v := ParseInts(s)
	i, count := 0, 0
	for i >= 0 && i < len(v) {
		n := v[i]
		if part2 && v[i] >= 3 {
			v[i]--
		} else {
			v[i]++
		}
		count++
		i += n
	}
	return count
}
