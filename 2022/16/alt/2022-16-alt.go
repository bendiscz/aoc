package main

import (
	"sort"
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
}

func (g *graph) node(name string) *node {
	if n, ok := g.nodes[name]; ok {
		return n
	}
	n := &node{id: len(g.nodes)}
	g.nodes[name] = n
	return n
}

type path struct {
	n *node
	d int
}

type node struct {
	id    int
	rate  int
	next  []*node
	paths []path
}

func (n *node) computePaths() {
	v := map[int]bool{}
	q := Queue[path]{}
	q.Push(path{n, 0})
	for q.Len() > 0 {
		p := q.Pop()
		if v[p.n.id] {
			continue
		}
		v[p.n.id] = true

		if p.d > 0 && p.n.rate > 0 {
			n.paths = append(n.paths, path{p.n, p.d + 1})
		}

		for _, n2 := range p.n.next {
			q.Push(path{n2, p.d + 1})
		}
	}
}

type state struct {
	n        *node
	open     uint64
	count    int
	rate     int
	left     int
	released int
}

func (s *state) hasOpen(n *node) bool {
	return s.open&(uint64(1)<<n.id) != 0
}

func (s *state) next(p path) (state, bool) {
	if p.d > s.left || s.hasOpen(p.n) {
		return state{}, false
	}
	return state{
		n:        p.n,
		open:     s.open | (uint64(1) << p.n.id),
		count:    s.count + 1,
		rate:     s.rate + p.n.rate,
		left:     s.left - p.d,
		released: s.released + p.d*s.rate,
	}, true
}

func (s *state) total() int {
	return s.released + s.rate*s.left
}

func (s *state) collides(s2 *state) bool {
	return s.open&s2.open != 0
}

func solve(p *Problem) {
	// "optimized" graph
	g := &graph{nodes: map[string]*node{}}
	for p.NextLine() {
		m := p.Parse(`^Valve ([A-Z]+) has flow rate=(\d+); tunnels? leads? to valves? (.*)$`)
		n := g.node(m[1])
		n.rate = ParseInt(m[2])
		for _, n2 := range strings.Split(m[3], ", ") {
			n.next = append(n.next, g.node(n2))
		}
	}

	start := g.nodes["AA"]
	valves := 0
	for _, n := range g.nodes {
		if n.rate > 0 {
			valves++
			n.computePaths()
		}
	}
	if start.rate == 0 {
		start.computePaths()
	}

	p.PartOne(search1(start))
	p.PartTwo(search2(start, valves))
}

func search1(start *node) int {
	q := Queue[state]{}
	q.Push(state{n: start, left: 30})

	best := state{}
	for q.Len() > 0 {
		s := q.Pop()

		terminal := true
		for _, p := range s.n.paths {
			if ns, ok := s.next(p); ok {
				terminal = false
				q.Push(ns)
			}
		}

		if terminal && s.total() > best.total() {
			best = s
		}
	}
	return best.total()
}

func search2(start *node, valves int) int {
	q := Queue[state]{}
	q.Push(state{n: start, left: 26})

	states := []state(nil)
	for q.Len() > 0 {
		s := q.Pop()
		states = append(states, s)
		for _, p := range s.n.paths {
			if ns, ok := s.next(p); ok {
				q.Push(ns)
			}
		}
	}

	sort.Slice(states, func(i, j int) bool { return states[i].count > states[j].count })
	best := 0
	for i := 0; i < len(states); i++ {
		s1 := &states[i]
		n := sort.Search(len(states), func(i int) bool { return s1.count+states[i].count <= valves })
		n = Max(n, i+1)
		if n == len(states) {
			continue
		}

		for j, l := n, states[n].count; j < len(states); j++ {
			s2 := &states[j]
			if s2.collides(s1) {
				continue
			}
			if s2.count < l {
				break
			}
			if s1.total()+s2.total() > best {
				best = s1.total() + s2.total()
			}
		}
	}

	return best
}
