package main

import (
	"fmt"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 24
	example = `

0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	m := parseMaterial(p)
	s := newState(m)
	s.search(0, 0, 0)
	p.PartOne(s.best1)
	p.PartOne(s.best2.str)
}

type port struct {
	t1, t2 int
	used   bool
}

func (p *port) String() string { return fmt.Sprintf("%d/%d", p.t1, p.t2) }

func (p *port) strength() int {
	return p.t1 + p.t2
}

func (p *port) exit(t int) int {
	if p.t1 == t {
		return p.t2
	} else {
		return p.t1
	}
}

type material struct {
	ports []*port
	pn    map[int][]*port
}

func parseMaterial(p *Problem) *material {
	m := &material{pn: map[int][]*port{}}
	for p.NextLine() {
		pt := &port{}
		p.Scanf("%d/%d", &pt.t1, &pt.t2)

		m.ports = append(m.ports, pt)
		m.pn[pt.t1] = append(m.pn[pt.t1], pt)
		if pt.t1 != pt.t2 {
			m.pn[pt.t2] = append(m.pn[pt.t2], pt)
		}
	}
	return m
}

func (m *material) state() uint64 {
	s := uint64(0)
	for i, pt := range m.ports {
		if pt.used {
			s |= 1 << i
		}
	}
	return s
}

type state struct {
	m     *material
	best1 int
	best2 struct {
		l, str int
	}
}

func newState(m *material) *state {
	return &state{
		m: m,
	}
}

func (s *state) search(t, l, str int) {
	s.best1 = max(s.best1, str)

	if s.best2.l < l || s.best2.l == l && s.best2.str < str {
		s.best2.l = l
		s.best2.str = str
	}

	pn := s.m.pn[t]
	for _, pt := range pn {
		if !pt.used {
			pt.used = true
			s.search(pt.exit(t), l+1, str+pt.strength())
			pt.used = false
		}
	}
}
