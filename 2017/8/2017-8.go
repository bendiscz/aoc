package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 8
	example = `

b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10

`
)

func main() {
	Run(year, day, example, solve)
}

type registers map[string]int

func solve(p *Problem) {
	regs := registers{}
	part2 := 0
	for p.NextLine() {
		f := SplitFields(p.Line())
		r0, op, x := f[4], f[5], ParseInt(f[6])
		if eval(regs[r0], x, op) {
			r1, d := f[0], ParseInt(f[2])
			if f[1] == "dec" {
				d *= -1
			}

			regs[r1] = regs[r1] + d
			part2 = max(part2, regs[r1])
		}
	}

	part1 := 0
	for _, x := range regs {
		part1 = max(x, part1)
	}
	p.PartOne(part1)
	p.PartTwo(part2)
}

func eval(a, b int, op string) bool {
	switch op {
	case "==":
		return a == b
	case "!=":
		return a != b
	case "<":
		return a < b
	case "<=":
		return a <= b
	case ">":
		return a > b
	case ">=":
		return a >= b
	default:
		return false
	}
}
