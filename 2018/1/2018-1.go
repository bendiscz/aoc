package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2018
	day     = 1
	example = `

+1
+1
-3

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	freq, steps := 0, []int(nil)
	for p.NextLine() {
		step := ParseInt(p.Line())
		steps = append(steps, step)
		freq += step
	}
	p.PartOne(freq)

	freq, freqs := 0, map[int]bool{}
	for i := 0; !freqs[freq]; i = (i + 1) % len(steps) {
		freqs[freq] = true
		freq += steps[i]
	}
	p.PartTwo(freq)
}
