package main

import (
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2022
	day     = 16
	example = `

Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II

`
)

func main() {
	Run(year, day, example, solve)
}

type graph struct {
	nodes map[string]*node
	names map[int]string
}

type node struct {
	id   int
	rate int
	next []*node
}

func (g *graph) node(name string) *node {
	if n, ok := g.nodes[name]; ok {
		return n
	}
	n := &node{id: len(g.nodes)}
	g.nodes[name] = n
	g.names[n.id] = name
	return n
}

func solve(p *Problem) {
	g := &graph{nodes: map[string]*node{}, names: map[int]string{}}
	for p.NextLine() {
		m := p.Parse(`^Valve ([A-Z]+) has flow rate=(\d+); tunnels? leads? to valves? (.*)$`)
		n := g.node(m[1])
		n.rate = ParseInt(m[2])
		for _, n2 := range strings.Split(m[3], ", ") {
			n.next = append(n.next, g.node(n2))
		}
	}

	start := g.nodes["AA"]

	p.PartOne(search1(map[state1]int{}, state1{start, 0, 0, 30}))
	p.PartTwo(search2a(map[state2]int{}, start, start, 0, 0, 26))
}

type state1 struct {
	n    *node
	open uint64
	rate int
	left int
}

func search1(cache map[state1]int, s state1) int {
	if s.left == 0 {
		return 0
	}

	if r, ok := cache[s]; ok {
		return r
	}

	best := s.rate * s.left
	mask := uint64(1) << s.n.id
	if s.open&mask == 0 && s.n.rate > 0 {
		r := s.rate + search1(cache, state1{s.n, s.open | mask, s.rate + s.n.rate, s.left - 1})
		best = Max(r, best)
	}
	for _, nn := range s.n.next {
		r := s.rate + search1(cache, state1{nn, s.open, s.rate, s.left - 1})
		best = Max(r, best)
	}

	cache[s] = best
	return best
}

type state2 struct {
	left   byte
	n1, n2 byte
	open   uint64
}

func search2a(cache map[state2]int, n1, n2 *node, open uint64, rate int, left byte) int {
	if left == 0 {
		return 0
	}

	if n1.id > n2.id {
		n1, n2 = n2, n1
	}

	s := state2{
		left: left,
		n1:   byte(n1.id),
		n2:   byte(n2.id),
		open: open,
	}
	if r, ok := cache[s]; ok {
		return r
	}

	best := rate * int(s.left)
	mask := uint64(1) << n1.id

	if open&mask == 0 && n1.rate > 0 {
		r := rate + search2b(cache, n1, n2, open|mask, rate+n1.rate, left)
		best = Max(r, best)
	}

	for _, nn := range n1.next {
		r := rate + search2b(cache, nn, n2, open, rate, left)
		best = Max(r, best)
	}

	r := rate + search2b(cache, n1, n2, open, rate, left)
	best = Max(r, best)

	cache[s] = best
	return best
}

func search2b(cache map[state2]int, n1, n2 *node, open uint64, rate int, left byte) int {
	best := 0
	mask := uint64(1) << n2.id

	if open&mask == 0 && n2.rate > 0 {
		r := search2a(cache, n1, n2, open|mask, rate+n2.rate, left-1)
		best = Max(r, best)
	}

	for _, nn := range n2.next {
		r := search2a(cache, n1, nn, open, rate, left-1)
		best = Max(r, best)
	}

	r := search2a(cache, n1, n2, open, rate, left-1)
	best = Max(r, best)

	return best
}
