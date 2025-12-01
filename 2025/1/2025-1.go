package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2025
	day     = 1
	example = `

L68
L30
R48
L5
R60
L55
L1
L99
R14
L82

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	ch byte
}

func (c cell) String() string { return string(c.ch) }

type grid struct {
	*Matrix[cell]
}

func solve(p *Problem) {
	d, s1, s2 := 50, 0, 0
	for p.NextLine() {
		l := p.Line()
		n := ParseInt(l[1:])

		//for n > 0 {
		//	if l[0] == 'L' {
		//		d--
		//	} else {
		//		d++
		//	}
		//
		//	if d == 100 {
		//		d = 0
		//	} else if d == -1 {
		//		d = 99
		//	}
		//
		//	if d == 0 {
		//		s2++
		//	}
		//
		//	n--
		//}
		//if d == 0 {
		//	s1++
		//}

		//-------------------------

		s2 += n / 100
		n = n % 100

		if n > 0 {
			if l[0] == 'L' {
				d -= n
				if d < 0 {
					if d+n != 0 {
						s2++
					}
					d += 100
				}
			} else {
				d += n
				if d > 99 {
					d -= 100
					if d > 0 {
						s2++
					}
				}
			}
		}

		if d == 0 {
			s1++
			s2++
		}
	}

	p.PartOne(s1)
	p.PartTwo(s2)
}
