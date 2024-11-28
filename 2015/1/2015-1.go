package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2015
	day     = 1
	example = `

(())
()()
(()(()(
))(((((
)())())
()())

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	for n := 1; p.NextLine(); n++ {
		p.Printf("line #%d", n)

		level := 0
		basement := 0

		for i, b := range []byte(p.Line()) {
			if b == '(' {
				level++
			} else if b == ')' {
				level--
			}

			if level < 0 && basement == 0 {
				basement = i + 1
			}
		}

		p.PartOne(level)
		p.PartTwo(basement)
	}
}
