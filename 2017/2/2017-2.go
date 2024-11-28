package main

import (
	"math"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 2
	example = `

5 9 2 8
9 4 7 3
3 8 6 5

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	sum1, sum2 := 0, 0
	for p.NextLine() {
		x0, x1 := math.MaxInt, math.MinInt
		xs := ParseInts(p.Line())
		for i, x := range xs {
			x0, x1 = min(x0, x), max(x1, x)
			for _, y := range xs[i+1:] {
				if x <= y && y%x == 0 {
					sum2 += y / x
				} else if x > y && x%y == 0 {
					sum2 += x / y
				}
			}
		}
		sum1 += Abs(x0 - x1)
	}
	p.PartOne(sum1)
	p.PartTwo(sum2)
}
