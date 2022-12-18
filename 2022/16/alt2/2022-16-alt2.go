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
	nodes  []*node
	named  map[string]*node
	start  *node
	valves int
}

func (g *graph) node(name string) *node {
	if n, ok := g.named[name]; ok {
		return n
	}
	n := &node{id: len(g.nodes)}
	g.nodes = append(g.nodes, n)
	g.named[name] = n
	return n
}

func (g *graph) prepare() {
	g.start = g.named["AA"]
	valves := 0
	for _, n := range g.nodes {
		if n.rate > 0 {
			n.mask = 1 << valves
			valves++
		}
	}
	g.valves = valves
}

type node struct {
	id   int
	rate int
	next []*node
	mask int
}

const (
	N  = 60
	M  = 1 << 15
	T1 = 30
	T2 = 26
)

type states [N][M]int

func (s *states) update(n, m, v int) {
	if v > s[n][m] {
		s[n][m] = v
	}
}

func solve(p *Problem) {
	// dynamic programming
	g := graph{named: map[string]*node{}}
	for p.NextLine() {
		m := p.Parse(`^Valve ([A-Z]+) has flow rate=(\d+); tunnels? leads? to valves? (.*)$`)
		n := g.node(m[1])
		n.rate = ParseInt(m[2])
		for _, name2 := range strings.Split(m[3], ", ") {
			n2 := g.node(name2)
			n.next = append(n.next, n2)
		}
	}

	g.prepare()

	s0, s1, s2 := &states{}, &states{}, &states{}

	rates, mm := [M]int{}, 1<<g.valves
	for m := 0; m < mm; m++ {
		rate := 0
		for _, n := range g.nodes {
			if n.mask&m != 0 {
				rate += n.rate
			}
		}
		rates[m] = rate
	}

	for t := 0; t < T1; t++ {
		for n := 0; n < len(g.nodes); n++ {
			cn := g.nodes[n]
			for m := 0; m < mm; m++ {
				if cn.rate > 0 && (m&cn.mask) != 0 {
					s1.update(n, m&^cn.mask, s0[n][m]+rates[m&^cn.mask])
				}
				for _, nn := range cn.next {
					s1.update(nn.id, m, s0[n][m]+rates[m])
				}
			}
		}

		s0, s1 = s1, s0

		if t == T2-1 {
			copy(s2[:], s0[:])
		}
	}

	aa := g.start.id
	p.PartOne(s0[aa][0])

	max := 0
	for m1, v := 0, (1<<g.valves)-1; m1 < mm; m1++ {
		m2 := v &^ m1
		max = Max(s2[aa][m1]+s2[aa][m2]-T2*(rates[m1]+rates[m2]), max)
	}
	p.PartTwo(max)
}
