package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2020
	day     = 25
	example = `

5764801
17807724

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	p.NextLine()
	pk1 := ParseUint(p.Line())
	p.NextLine()
	pk2 := ParseUint(p.Line())

	l1, l2 := 0, 0
	sub := uint(1)
	for i := 1; l1 == 0 || l2 == 0; i++ {
		sub = (sub * 7) % 20201227
		if sub == pk1 {
			l1 = i
		}
		if sub == pk2 {
			l2 = i
		}
	}

	sub = 1
	for i := 0; i < l1; i++ {
		sub = (sub * pk2) % 20201227
	}
	p.PartOne(sub)
}
