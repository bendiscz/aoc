package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 1
	example = `

1122
1111
1234
91212129
1212

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	for p.NextLine() {
		s, sum1, sum2 := p.Line(), 0, 0
		pch := s[len(s)-1]
		for i, ch := range []byte(s) {
			if ch == pch {
				sum1 += int(ch - '0')
			}
			pch = ch

			if ch == s[(i+len(s)/2)%len(s)] {
				sum2 += int(ch - '0')
			}
		}
		p.PartOne(sum1)
		p.PartTwo(sum2)
	}
}
