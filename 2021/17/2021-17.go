package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2021
	day     = 17
	example = `

target area: x=20..30, y=-10..-5

`
)

func main() {
	Run(year, day, example, solve)
}

func simulate(x1, x2, y1, y2, dx, dy int) bool {
	x, y := 0, 0
	for x <= x2 && y >= y1 {
		x += dx
		y += dy

		if x >= x1 && x <= x2 && y <= y2 && y >= y1 {
			return true
		}

		if dx > 0 {
			dx--
		}
		dy--
	}
	return false
}

func solve(p *Problem) {
	p.NextLine()
	m := p.Parse(`^target area: x=(-?\d+)\.\.(-?\d+), y=(-?\d+)\.\.(-?\d+)$`)
	x1, x2, y1, y2 := ParseInt(m[1]), ParseInt(m[2]), ParseInt(m[3]), ParseInt(m[4])

	mdy := -y1 - 1
	y := (1 + mdy) * mdy / 2
	p.PartOne(y)

	count := 0
	for dx := x2; dx > 0; dx-- {
		for dy := mdy; dy >= y1; dy-- {
			if simulate(x1, x2, y1, y2, dx, dy) {
				count++
			}
		}
	}

	p.PartTwo(count)
}
