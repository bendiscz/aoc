package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2021
	day     = 5
	example = `

0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	p.PartOne(solvePart(p, false))
	p.Reset()
	p.PartTwo(solvePart(p, true))
}

func solvePart(p *Problem, diag bool) int {
	screen, count := [1000][1000]uint8{}, 0
	fill := func(x, y int) {
		n := screen[x][y]
		if n < 2 {
			if n == 1 {
				count++
			}
			screen[x][y] = n + 1
		}
	}

	for p.NextLine() {
		m := p.Parse(`^(\d+),(\d+) -> (\d+),(\d+)$`)
		x1, y1, x2, y2 := ParseInt(m[1]), ParseInt(m[2]), ParseInt(m[3]), ParseInt(m[4])

		if x1 == x2 {
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for y := y1; y <= y2; y++ {
				fill(x1, y)
			}
		} else if y1 == y2 {
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			for x := x1; x <= x2; x++ {
				fill(x, y1)
			}
		} else if diag && Abs(x2-x1) == Abs(y2-y1) {
			// part two
			dx := dir(x1, x2)
			dy := dir(y1, y2)
			for i := 0; i <= Abs(x2-x1); i++ {
				fill(x1+dx*i, y1+dy*i)
			}
		}

	}

	return count
}

func dir(a1, a2 int) int {
	if a1 < a2 {
		return 1
	} else {
		return -1
	}
}
