package main

import (
	_ "embed"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year = 2022
	day  = 10
)

//go:embed example
var example string

func main() {
	Run(year, day, example, solve)
}

type display struct {
	*Matrix[bool]
}

func (d display) draw(cycle, reg int) {
	y, x := cycle/40, cycle%40
	if Abs(x-reg) <= 1 {
		*d.AtXY(x, y) = true
	}
}

func solve(p *Problem) {
	reg, cycle, next, sum := 1, 0, 20, 0
	disp := display{NewMatrix[bool](Rectangle(40, 6))}

	for p.NextLine() {
		d := 0
		if p.Line() == "noop" {
			disp.draw(cycle, reg)
			cycle += 1
		} else {
			disp.draw(cycle, reg)
			disp.draw(cycle+1, reg)
			cycle += 2
			d = ParseInt(p.Line()[5:])
		}

		if cycle >= next {
			sum += next * reg
			next += 40
		}

		reg += d
	}

	p.PartOne(sum)
	p.PartTwo("")
	PrintGridBool(disp)
}
