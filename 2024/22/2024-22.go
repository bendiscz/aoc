package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2024
	day     = 22
	example = `

1
2
3
2024
`
)

func main() {
	Run(year, day, example, solve)
}

const M = 16777216

func next(x int) int {
	x ^= x << 6
	x %= M
	x ^= x >> 5
	x ^= x << 11
	x %= M
	return x
}

type diff [4]int

func generateSecrets(sums map[diff]int, seed, n int) int {
	x, prev, d := seed, seed%10, diff{}
	seen := Set[diff]{}

	for i := 0; i < n; i++ {
		x = next(x)
		b := x % 10
		copy(d[:], d[1:])
		d[3], prev = b-prev, b
		if i >= 4 && !seen.Contains(d) {
			seen[d] = SET
			sums[d] += b
		}
	}
	return x
}

func solve(p *Problem) {
	s1, sums := 0, map[diff]int{}
	for p.NextLine() {
		seed := ParseInt(p.Line())
		s1 += generateSecrets(sums, seed, 2000)
	}
	p.PartOne(s1)

	s2 := 0
	for _, x := range sums {
		s2 = max(s2, x)
	}
	p.PartTwo(s2)
}
