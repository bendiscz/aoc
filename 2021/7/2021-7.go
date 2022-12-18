package main

import (
	"math"
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2021
	day     = 7
	example = `

16,1,2,0,4,2,7,1,2,14

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	var positions []int
	max := 0
	for _, str := range strings.Split(strings.TrimSpace(p.ReadAll()), ",") {
		p := ParseInt(str)
		positions = append(positions, ParseInt(str))
		if p > max {
			max = p
		}
	}

	p.PartOne(solvePart(positions, max, func(d int) int { return d }))
	p.PartTwo(solvePart(positions, max, func(d int) int { return (d + 1) * d / 2 }))
}

func solvePart(positions []int, max int, dist func(int) int) int {
	min := math.MaxInt
	for x := 0; x <= max; x++ {
		fuel := 0
		for _, p := range positions {
			d := Abs(x - p)
			fuel += dist(d)
		}
		if fuel < min {
			min = fuel
		}
	}

	return min
}
