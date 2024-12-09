package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2020
	day     = 2
	example = `

1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	s1, s2 := 0, 0
	for p.NextLine() {
		var a, b int
		var ch byte
		var pwd string
		p.Scanf("%d-%d %c: %s", &a, &b, &ch, &pwd)

		c := 0
		for i := 0; i < len(pwd); i++ {
			if pwd[i] == ch {
				c++
			}
		}

		if a <= c && c <= b {
			s1++
		}

		ch1, ch2 := pwd[a-1], pwd[b-1]
		if ch1 == ch && ch2 != ch || ch1 != ch && ch2 == ch {
			s2++
		}
	}

	p.PartOne(s1)
	p.PartTwo(s2)
}
