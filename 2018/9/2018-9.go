package main

import (
	"math"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2018
	day     = 9
	example = `

9 players; last marble is worth 25 points
10 players; last marble is worth 1618 points
13 players; last marble is worth 7999 points
17 players; last marble is worth 1104 points
21 players; last marble is worth 6111 points
30 players; last marble is worth 5807 points

`
)

func main() {
	Run(year, day, example, solve)
}

type marble struct {
	value      int
	prev, next *marble
}

func (m *marble) insert(value int) *marble {
	n := &marble{
		value: value,
		prev:  m,
		next:  m.next,
	}
	n.prev.next = n
	n.next.prev = n
	return n
}

func (m *marble) remove() (*marble, int) {
	m.prev.next = m.next
	m.next.prev = m.prev
	return m.next, m.value
}

func solve(p *Problem) {
	for p.NextLine() {
		f := ParseInts(p.Line())
		p.PartOne(part(f[0], f[1]))
		p.PartTwo(part(f[0], 100*f[1]))
	}
}

func part(elves, marbles int) int {
	head := &marble{}
	head.prev = head
	head.next = head

	score := make([]int, elves)

	for value := 1; value <= marbles; value++ {
		if value%23 != 0 {
			head = head.next.insert(value)
		} else {
			elf := (value - 1) % len(score)
			m, s := head.prev.prev.prev.prev.prev.prev.prev.remove()
			head = m
			score[elf] += value + s
		}
	}

	high := math.MinInt
	for _, x := range score {
		high = max(high, x)
	}
	return high
}
