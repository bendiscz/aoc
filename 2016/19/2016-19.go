package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2016
	day     = 19
	example = `

5
10

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	for p.NextLine() {
		n := ParseInt(p.Line())
		p.PartOne(part1(n))
		p.PartTwo(part2(n))
	}

}
func part1(n int) int {
	f, s, r := 1, 1, 0
	for n > 1 {
		if r == 1 {
			f += s
		}
		s = s * 2
		r = n % 2
		n = (n + 1 - r) / 2
	}
	return f
}

func part2(n int) int {
	type elf struct {
		id   int
		prev int
		next int
	}

	elves := make([]elf, n)
	for i := 0; i < n; i++ {
		elves[i] = elf{
			id:   i + 1,
			prev: i - 1,
			next: i + 1,
		}
	}
	elves[0].prev = n - 1
	elves[n-1].next = 0

	v := n / 2
	for ; n > 1; n-- {
		e := elves[v]
		elves[e.prev].next = e.next
		elves[e.next].prev = e.prev
		v = e.next

		if n%2 == 1 {
			v = elves[v].next
		}
	}

	return elves[v].id
}
