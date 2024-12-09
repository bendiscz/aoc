package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2020
	day     = 6
	example = `

abc

a
b
c

ab
ac

a
a
a
a

b

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	s1, s2 := 0, 0
	c, l := map[byte]int{}, 0
	for {
		p.NextLine()
		s := p.Line()

		if s == "" {
			if len(c) == 0 {
				break
			}

			s1 += len(c)
			for _, x := range c {
				if l == x {
					s2++
				}
			}

			clear(c)
			l = 0
			continue
		}

		for i := 0; i < len(s); i++ {
			c[s[i]]++
		}
		l++
	}

	p.PartOne(s1)
	p.PartTwo(s2)
}
