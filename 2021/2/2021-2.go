package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2021
	day     = 2
	example = `

forward 5
down 5
forward 8
up 3
down 8
forward 2

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	dist, d1, d2, aim := 0, 0, 0, 0
	for p.NextLine() {
		m := p.Parse(`^(forward|up|down) (\d+)$`)
		x := ParseInt(m[2])
		switch m[1] {
		case "forward":
			dist += x
			d2 += aim * x
		case "down":
			d1 += x
			aim += x
		case "up":
			d1 -= x
			aim -= x
		}
	}

	p.PartOne(dist * d1)
	p.PartTwo(dist * d2)
}
