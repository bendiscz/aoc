package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 11
	example = `

ne,ne,ne
ne,ne,sw,sw
ne,ne,s,s
se,sw,se,sw,sw

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	for p.NextLine() {
		xy, f := XY{}, 0
		for _, dir := range SplitFields(p.Line()) {
			xy = move(xy, dir)
			if d := dist(xy); d > f {
				f = d
			}
		}

		p.PartOne(dist(xy))
		p.PartTwo(f)
	}
}

func move(xy XY, dir string) XY {
	switch dir {
	case "n":
		xy.Y++

	case "s":
		xy.Y--

	case "ne":
		if xy.X%2 != 0 {
			xy.Y++
		}
		xy.X++

	case "se":
		if xy.X%2 == 0 {
			xy.Y--
		}
		xy.X++

	case "nw":
		if xy.X%2 != 0 {
			xy.Y++
		}
		xy.X--

	case "sw":
		if xy.X%2 == 0 {
			xy.Y--
		}
		xy.X--
	}
	return xy
}

func dist(xy XY) int {
	d := Abs(xy.X)
	if xy.Y >= 0 {
		return d + max(xy.Y-d/2, 0)
	} else {
		return d + max(-xy.Y-(d+1)/2, 0)
	}
}
