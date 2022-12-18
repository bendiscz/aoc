package main

import (
	"math"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2021
	day     = 14
	example = `

NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C

`
)

func main() {
	Run(year, day, example, solve)
}

type pair struct {
	ch1, ch2 byte
}

func solve(p *Problem) {
	p.NextLine()
	polymer := p.Line()
	p.NextLine()

	rules := map[pair]byte{}
	for p.NextLine() {
		m := p.Parse(`^(\w)(\w) -> (\w)$`)
		rules[pair{m[1][0], m[2][0]}] = m[3][0]
	}

	pairs := map[pair]int{}
	for i := 1; i < len(polymer); i++ {
		p := pair{polymer[i-1], polymer[i]}
		pairs[p] = pairs[p] + 1
	}

	for i := 0; i < 10; i++ {
		pairs = expand(rules, pairs)
	}
	p.PartOne(findMinMax(pairs))

	for i := 0; i < 30; i++ {
		pairs = expand(rules, pairs)
	}
	p.PartTwo(findMinMax(pairs))
}

func expand(rules map[pair]byte, pairs map[pair]int) map[pair]int {
	next := map[pair]int{}
	for p, x := range pairs {
		if ch, ok := rules[p]; ok {
			p1 := pair{p.ch1, ch}
			p2 := pair{ch, p.ch2}
			next[p1] = next[p1] + x
			next[p2] = next[p2] + x
		}
	}
	return next
}

func findMinMax(pairs map[pair]int) int {
	hist := map[byte]int{}
	for p, x := range pairs {
		hist[p.ch1] = hist[p.ch1] + x
		hist[p.ch2] = hist[p.ch2] + x
	}

	max, min := math.MinInt, math.MaxInt
	for _, x := range hist {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}

	return (max - min + 2) / 2
}
