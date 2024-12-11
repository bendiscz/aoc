package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2024
	day     = 11
	example = `

125 17

`
)

func main() {
	Run(year, day, example, solve)
}

type state struct {
	x, n int
}

func solve(p *Problem) {
	s1, s2, cache := 0, 0, map[state]int{}
	p.NextLine()
	for _, s := range ParseInts(p.Line()) {
		s1 += count(cache, s, 25)
		s2 += count(cache, s, 75)
	}
	p.PartOne(s1)
	p.PartTwo(s2)
}

func count(cache map[state]int, x, n int) (r int) {
	if n == 0 {
		return 1
	}

	st := state{x, n}
	if c, ok := cache[st]; ok {
		return c
	}
	defer func() {
		cache[st] = r
	}()

	n--

	if x == 0 {
		return count(cache, 1, n)
	}

	d := 0
	for i := x; i > 0; i /= 10 {
		d++
	}
	if d%2 == 1 {
		return count(cache, x*2024, n)
	}

	b := 10
	for i := 1; i < d/2; i++ {
		b *= 10
	}
	return count(cache, x/b, n) + count(cache, x%b, n)
}
