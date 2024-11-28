package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2018
	day     = 0
	example = `

example

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	for p.NextLine() {
		p.Printf("%s", p.Line())
	}

	p.PartOne("TODO")

	p.PartTwo("TODO")
}
