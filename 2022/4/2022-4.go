package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2022
	day     = 4
	example = `

2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	count1, count2 := 0, 0
	for p.NextLine() {
		m := p.Parse(`^(\d+)-(\d+),(\d+)-(\d+)$`)
		i1, j1, i2, j2 := ParseInt(m[1]), ParseInt(m[2]), ParseInt(m[3]), ParseInt(m[4])
		if i1 >= i2 && j1 <= j2 || i2 >= i1 && j2 <= j1 {
			count1++
		}
		if !(j1 < i2 || i1 > j2) {
			count2++
		}
	}

	p.PartOne(count1)
	p.PartTwo(count2)
}
