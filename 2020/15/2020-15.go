package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2020
	day     = 15
	example = `

0,3,6
1,3,2
2,1,3
1,2,3
2,3,1
3,2,1
3,1,2

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	for p.NextLine() {
		s := p.Line()
		ts, last := map[int]int{}, 0
		for i, x := range ParseInts(s) {
			ts[x] = i + 1
			last = x
		}
		delete(ts, last)

		for i := len(ts) + 2; i <= 30000000; i++ {
			if i-1 == 2020 {
				p.PartOne(last)
			}

			t, ok := ts[last]
			ts[last] = i - 1
			if ok {
				last = i - t - 1
			} else {
				last = 0
			}
		}

		p.PartTwo(last)
	}

}
