package main

import (
	"math"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2021
	day     = 1
	example = `

199
200
208
210
200
207
240
269
260
263

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	r1, prev := 0, math.MaxInt
	for p.NextLine() {
		d := ParseInt(p.Line())
		if d > prev {
			r1++
		}
		prev = d
	}

	p.PartOne(r1)

	p.Reset()
	win, i, r2 := [3]int{}, 0, 0
	for i < 3 && p.NextLine() {
		win[i] = ParseInt(p.Line())
		i++
	}
	i = 0
	for p.NextLine() {
		prev = win[0] + win[1] + win[2]
		win[i] = ParseInt(p.Line())
		i = (i + 1) % 3
		if win[0]+win[1]+win[2] > prev {
			r2++
		}
	}

	p.PartTwo(r2)
}
