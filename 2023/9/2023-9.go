package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2023
	day     = 9
	example = `

0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	sum1, sum2 := 0, 0
	for p.NextLine() {
		seq := ParseInts(p.Line())

		a := make([]int, len(seq))
		copy(a, seq)
		sum1 += extrapolate(a)

		for i := 0; i < len(seq); i++ {
			a[i] = seq[len(seq)-i-1]
		}
		sum2 += extrapolate(a)
	}
	p.PartOne(sum1)
	p.PartTwo(sum2)
}

func extrapolate(a []int) int {
	n := len(a) - 1
	for z := false; !z; n-- {
		z = true
		for i := 0; i < n; i++ {
			a[i] = a[i+1] - a[i]
			if a[i] != 0 {
				z = false
			}
		}
	}

	x := 0
	for i := n; i < len(a); i++ {
		x += a[i]
	}
	return x
}
