package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 15
	example = `

Generator A starts with 65
Generator B starts with 8921

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	const xa, xb, r = 16807, 48271, 2147483647

	var sa, sb uint64
	p.NextLine()
	p.Scanf("Generator A starts with %d", &sa)
	p.NextLine()
	p.Scanf("Generator B starts with %d", &sb)

	a, b, sum := sa, sb, 0
	for i := 0; i < 40_000_000; i++ {
		a = (a * xa) % r
		b = (b * xb) % r
		if (a^b)&0xffff == 0 {
			sum++
		}
	}
	p.PartOne(sum)

	a, b, sum = sa, sb, 0
	for i := 0; i < 5_000_000; i++ {
		for {
			a = (a * xa) % r
			if a%4 == 0 {
				break
			}
		}
		for {
			b = (b * xb) % r
			if b%8 == 0 {
				break
			}
		}
		if (a^b)&0xffff == 0 {
			sum++
		}
	}
	p.PartTwo(sum)
}
