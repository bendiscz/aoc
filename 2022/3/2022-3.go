package main

import (
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2022
	day     = 3
	example = `

vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	sum := 0
	for p.NextLine() {
		l := len(p.Line()) / 2
		p1, p2 := p.Line()[:l], p.Line()[l:]

		for _, ch := range []byte(p1) {
			if strings.IndexByte(p2, ch) != -1 {
				sum += prio(ch)
				break
			}
		}
	}

	p.PartOne(sum)

	sum = 0
	p.Reset()
	for p.NextLine() {
		counts := [100]int{}
		for _, ch := range []byte(p.Line()) {
			counts[prio(ch)] |= 1
		}
		p.NextLine()
		for _, ch := range []byte(p.Line()) {
			counts[prio(ch)] |= 2
		}
		p.NextLine()
		for _, ch := range []byte(p.Line()) {
			counts[prio(ch)] |= 4
		}

		for prio, count := range counts {
			if count == 7 {
				sum += prio
			}
		}
	}

	p.PartTwo(sum)
}

func prio(ch byte) int {
	if ch >= 'a' {
		return int(ch - 'a' + 1)
	} else {
		return int(ch - 'A' + 27)
	}
}
