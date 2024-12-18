package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"strings"
)

const (
	year    = 2019
	day     = 4
	example = ``
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	s1, s2 := 0, 0
	f := SplitFieldsDelim(strings.TrimSpace(p.ReadAll()), "-")
	a, b := ParseInt(f[0]), ParseInt(f[1])
	for x := a; x <= b; x++ {
		c1, c2 := check(x)
		if c1 {
			s1++
		}
		if c2 {
			s2++
		}
	}
	p.PartOne(s1)
	p.PartTwo(s2)
}

func check(x int) (bool, bool) {
	l, c, p1, p2 := 10, 1, false, false
	for i := 0; i < 6; i++ {
		d := x % 10
		x /= 10
		if d > l {
			return false, false
		}
		if d == l {
			p1 = true
			c++
		} else {
			if c == 2 {
				p2 = true
			}
			c = 1
			l = d
		}
	}
	return p1, p2 || c == 2
}
