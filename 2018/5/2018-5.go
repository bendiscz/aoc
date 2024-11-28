package main

import (
	"math"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2018
	day     = 5
	example = `

dabAcCaCBAcCcaDA

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	for p.NextLine() {
		p.PartOne(reduce(p.LineBytes(), 0))

		p2 := math.MaxInt
		for ch := byte('A'); ch <= 'Z'; ch++ {
			p2 = min(p2, reduce(p.LineBytes(), ch))
		}
		p.PartTwo(p2)
	}
}

func reduce(s []byte, skip byte) int {
	count, length := 0, len(s)
	prev := make([]int, length)

	for i1, i2 := -1, 0; i2 < len(s); {
		if i1 >= 0 && Abs(int(s[i1])-int(s[i2])) == 'a'-'A' {
			count++
			i1 = prev[i1]
		} else {
			i1, prev[i2] = i2, i1
		}

		for {
			i2++
			if i2 == len(s) || s[i2] != skip && s[i2] != skip+('a'-'A') {
				break
			}
			length--
		}
	}

	return length - 2*count
}
