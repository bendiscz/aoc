package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"slices"
	"strings"
)

const (
	year    = 2020
	day     = 23
	example = `

389125467

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	s := strings.TrimSpace(p.ReadAll())

	cups := make([]int, 10)
	curr := int(s[0] - '0')
	for i := 1; i <= 9; i++ {
		c := int(s[i-1] - '0')
		cups[c] = cups[curr]
		cups[curr] = c
		curr = c
	}
	last := curr
	curr = cups[curr]

	p.PartOne(partOne(slices.Clone(cups), curr))
	p.PartTwo(partTwo(cups, curr, last))
}

func partOne(cups []int, curr int) string {
	for i := 0; i < 100; i++ {
		curr = move(cups, curr)
	}

	b := make([]byte, 8)
	for i, c := 0, cups[1]; i < 8; i, c = i+1, cups[c] {
		b[i] = byte('0' + c)
	}
	return string(b)
}

func partTwo(cups []int, curr, last int) int {
	mega := make([]int, 1_000_001)
	copy(mega, cups)
	cups = mega
	for i := 10; i <= 1_000_000; i++ {
		cups[i] = cups[last]
		cups[last] = i
		last = i
	}

	for i := 0; i < 10_000_000; i++ {
		curr = move(cups, curr)
	}
	return cups[1] * cups[cups[1]]
}

func move(cups []int, curr int) int {
	trio := cups[curr]
	t1, t2, t3 := trio, cups[trio], cups[cups[trio]]
	cups[curr] = cups[t3]

	dest := curr - 1
	for {
		if dest == 0 {
			dest = len(cups) - 1
		}
		if dest != t1 && dest != t2 && dest != t3 {
			break
		}
		dest--
	}

	cups[t3] = cups[dest]
	cups[dest] = trio
	return cups[curr]
}
