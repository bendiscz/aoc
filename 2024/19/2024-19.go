package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"strings"
)

const (
	year    = 2024
	day     = 19
	example = `

r, wr, b, g, bwu, rb, gb, br

bwurrg
brwrr
bggr
gbbr
rrbgbr
ubwu
brgr
bbrgwb

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	patterns := SplitFields(p.ReadLine())
	p.NextLine()

	s1, s2 := 0, 0
	for p.NextLine() {
		c := count(patterns, p.Line())
		if c > 0 {
			s1++
		}
		s2 += c
	}
	p.PartOne(s1)
	p.PartTwo(s2)
}

func count(patterns []string, towel string) int {
	d := make([]int, len(towel)+1)
	d[0] = 1
	for i := 0; i < len(towel); i++ {
		for _, p := range patterns {
			if strings.HasPrefix(towel[i:], p) {
				d[i+len(p)] += d[i]
			}
		}
	}
	return d[len(d)-1]
}
