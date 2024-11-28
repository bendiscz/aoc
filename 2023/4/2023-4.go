package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2023
	day     = 4
	example = `

Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	sum1, sum2 := 0, 0
	c, i := []int(nil), 0
	for p.NextLine() {
		n := parseCard(p.Line())

		c = add(c, i, 1)
		sum1 += 1 << n >> 1
		sum2 += c[i]

		for j := i + 1; j <= i+n; j++ {
			c = add(c, j, c[i])
		}

		i++
	}
	p.PartOne(sum1)
	p.PartTwo(sum2)
}

func add(c []int, i, d int) []int {
	for i >= len(c) {
		c = append(c, 0)
	}
	c[i] += d
	return c
}

func parseCard(s string) int {
	u := map[int]bool{}
	f := SplitFields(s)[2:]
	for ; f[0] != "|"; f = f[1:] {
		u[ParseInt(f[0])] = true

	}

	n := 0
	for f = f[1:]; len(f) != 0; f = f[1:] {
		if u[ParseInt(f[0])] {
			n++
		}
	}
	return n
}
