package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2016
	day     = 18
	example = `

.^^.^.^^^^

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	row, n := p.ReadLine(), 40
	if p.Example() {
		n = 10
	}

	safe := 0
	for i := 0; i < 400_000; i++ {
		if i == n {
			p.PartOne(safe)
		}

		next := make([]byte, len(row))
		for j := 0; j < len(row); j++ {
			if row[j] == '.' {
				safe++
			}

			var l, c, r bool
			if j > 0 {
				l = row[j-1] == '^'
			}
			if j < len(row)-1 {
				r = row[j+1] == '^'
			}
			c = row[j] == '^'

			if l && c && !r || !l && c && r || l && !c && !r || !l && !c && r {
				next[j] = '^'
			} else {
				next[j] = '.'
			}
		}
		row = string(next)
	}
	p.PartTwo(safe)
}
