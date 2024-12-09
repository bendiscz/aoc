package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2020
	day     = 5
	example = `

FBFBBFFRLR

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	s1, s2 := 0, 0
	seats := map[int]bool{}
	for p.NextLine() {
		x, s := 0, p.Line()
		for i, b := len(s)-1, 1; i >= 0; i, b = i-1, b*2 {
			if s[i] == 'B' || s[i] == 'R' {
				x += b
			}
		}
		seats[x] = true
		s1 = max(s1, x)
	}

	p.PartOne(s1)

	for s := 1; s < 1023; s++ {
		if !seats[s] && seats[s-1] && seats[s+1] {
			s2 = s
			break
		}
	}

	p.PartTwo(s2)
}
