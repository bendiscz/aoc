package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2022
	day     = 20
	example = `

1
2
-3
3
-2
0
4

`
)

func main() {
	Run(year, day, example, solve)
}

type num struct {
	value      int
	prev, next *num
}

func move(n *num, d int) *num {
	for d < 0 {
		n = n.prev
		d++
	}
	for d > 0 {
		n = n.next
		d--
	}
	return n
}

func mix(n *num, l int) {
	t := n.prev
	n.prev.next = n.next
	n.next.prev = n.prev

	t = move(t, n.value%(l-1))
	n.prev = t
	n.next = t.next
	n.prev.next = n
	n.next.prev = n
}

func mixAll(m []*num) {
	for _, n := range m {
		mix(n, len(m))
	}
}

func reset(m []*num) {
	for i := 1; i < len(m); i++ {
		m[i].prev = m[i-1]
		m[i-1].next = m[i]
	}
	m[0].prev = m[len(m)-1]
	m[len(m)-1].next = m[0]
}

func sum(n0 *num) int {
	s := 0
	for i, n := 0, n0; i < 3; i++ {
		n = move(n, 1000)
		s += n.value
	}
	return s
}

func solve(p *Problem) {
	var m []*num
	var n0 *num
	for p.NextLine() {
		n := &num{value: ParseInt(p.Line())}
		if n.value == 0 {
			n0 = n
		}
		m = append(m, n)
	}

	reset(m)
	mixAll(m)
	p.PartOne(sum(n0))

	const key = 811589153
	for _, n := range m {
		n.value *= key
	}

	reset(m)
	for i := 0; i < 10; i++ {
		mixAll(m)
	}

	p.PartOne(sum(n0))
}
