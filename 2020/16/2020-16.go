package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"strings"
)

const (
	year    = 2020
	day     = 16
	example = `

class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12

`
)

func main() {
	Run(year, day, example, solve)
}

type field struct {
	departure bool
	a1, b1    int
	a2, b2    int
	positions map[int]bool
	marked    bool
}

func (f field) matches(x int) bool {
	return x >= f.a1 && x <= f.b1 || x >= f.a2 && x <= f.b2
}

func (f field) index() int {
	for i := range f.positions {
		return i
	}
	panic("no position")
}

func solve(p *Problem) {
	fields := []*field(nil)
	for p.NextLine() {
		if p.Line() == "" {
			break
		}
		f := ParseInts(strings.ReplaceAll(p.Line(), "-", " "))
		fields = append(fields, &field{strings.HasPrefix(p.Line(), "departure "), f[0], f[1], f[2], f[3], make(map[int]bool), false})
	}

	for _, f := range fields {
		for i := 0; i < len(fields); i++ {
			f.positions[i] = true
		}
	}

	p.NextLine()
	p.NextLine()
	ticket := ParseInts(p.Line())
	p.NextLine()

	p.NextLine()
	s1 := 0
	for p.NextLine() {
		for i, x := range ParseInts(p.Line()) {
			invalid := true
			for _, f := range fields {
				if f.matches(x) {
					invalid = false
					break
				}
			}
			if invalid {
				s1 += x
			} else {
				for _, f := range fields {
					if !f.matches(x) {
						delete(f.positions, i)
					}
				}
			}
		}
	}
	p.PartOne(s1)

	for {
		i0 := -1
		for i, f := range fields {
			if f.marked {
				continue
			}

			if len(f.positions) == 1 {
				i0 = i
				break
			}
		}

		if i0 == -1 {
			break
		}

		pos := fields[i0].index()
		fields[i0].marked = true

		for i, f := range fields {
			if i == i0 {
				continue
			}
			delete(f.positions, pos)
		}
	}

	s2 := 1
	for _, f := range fields {
		if !f.departure {
			continue
		}
		s2 *= ticket[f.index()]
	}
	p.PartTwo(s2)
}
