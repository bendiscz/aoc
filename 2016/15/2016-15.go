package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2016
	day     = 15
	example = `

Disc #1 has 5 positions; at time=0, it is at position 4.
Disc #2 has 2 positions; at time=0, it is at position 1.

`
)

func main() {
	Run(year, day, example, solve)
}

type disc struct {
	length, offset int
}

func solve(p *Problem) {
	discs := []disc(nil)
	for p.NextLine() {
		var i, l, o int
		p.Scanf("Disc #%d has %d positions; at time=0, it is at position %d.", &i, &l, &o)
		discs = append(discs, disc{l, o})
	}

	p.PartOne(solvePart(discs))

	discs = append(discs, disc{11, 0})
	p.PartTwo(solvePart(discs))
}

func solvePart(discs []disc) int {
loop:
	for t := 0; ; t++ {
		for i, d := range discs {
			if (t+d.offset+i+1)%d.length != 0 {
				continue loop
			}
		}
		return t
	}
}
