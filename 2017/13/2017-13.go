package main

import (
	"sort"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 13
	example = `

0: 3
1: 2
4: 4
6: 4

`
)

func main() {
	Run(year, day, example, solve)
}

type wall struct {
	t, l, c int
}

func solve(p *Problem) {
	walls := []wall(nil)
	for p.NextLine() {
		x := ParseInts(p.Line())
		w := wall{x[0], x[1], (x[1] - 1) * 2}
		walls = append(walls, w)
	}
	sort.Slice(walls, func(i, j int) bool { return walls[i].l < walls[j].l })

	p.PartOne(send(walls))

	for t0 := 0; ; t0++ {
		if check(walls, t0) {
			p.PartTwo(t0)
			break
		}
	}
}

func send(walls []wall) int {
	severity := 0
	for _, w := range walls {
		if w.t%w.c == 0 {
			severity += w.t * w.l
		}
	}
	return severity
}

func check(walls []wall, t0 int) bool {
	for _, w := range walls {
		if (w.t+t0)%w.c == 0 {
			return false
		}
	}
	return true
}
