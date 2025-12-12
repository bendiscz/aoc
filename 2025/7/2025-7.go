package main

import (
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2025
	day     = 7
	example = `

.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	s1, s2 := 0, 0

	p.NextLine()
	s := make([]int, len(p.Line()))
	n := make([]int, len(p.Line()))
	s[strings.IndexByte(p.Line(), 'S')] = 1

	p.SkipLines(1)
	for p.NextLine() {
		for i, ch := range p.Line() {
			if ch == '^' {
				if s[i] > 0 {
					s1++
					n[i-1] += s[i]
					n[i+1] += s[i]
				}
			} else {
				n[i] += s[i]
			}
		}

		copy(s, n)
		for j := range n {
			n[j] = 0
		}
		p.SkipLines(1)
	}

	for _, x := range s {
		s2 += x
	}

	p.PartOne(s1)
	p.PartTwo(s2)
}
