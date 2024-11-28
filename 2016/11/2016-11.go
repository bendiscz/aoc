package main

import (
	"fmt"
	"regexp"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2016
	day     = 11
	example = `

The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.
The second floor contains a hydrogen generator.
The third floor contains a lithium generator.
The fourth floor contains nothing relevant.

`
)

func main() {
	Run(year, day, example, solve)
}

const max = 7

type state struct {
	el byte
	g  [max]byte
	m  [max]byte
}

func parseInitState(p *Problem) (state, int) {
	elements := map[string]int{}
	ei := func(e string) int {
		if i, ok := elements[e]; ok {
			return i
		}
		i := len(elements)
		elements[e] = i
		return i
	}

	genRE := regexp.MustCompile(`([A-Za-z]+) generator`)
	mcuRE := regexp.MustCompile(`([A-Za-z]+)-compatible microchip`)

	s := state{el: 1}
	for f := byte(1); f <= 4 && p.NextLine(); f++ {
		for _, m := range genRE.FindAllStringSubmatch(p.Line(), -1) {
			s.g[ei(m[1])] = f
		}
		for _, m := range mcuRE.FindAllStringSubmatch(p.Line(), -1) {
			s.m[ei(m[1])] = f
		}
	}
	return s, len(elements)
}

func (s state) print(elems int) {
	for f := byte(4); f > 0; f-- {
		fmt.Printf("F%d ", f)
		if s.el == byte(f) {
			fmt.Printf("E  ")
		} else {
			fmt.Printf(".  ")
		}
		for i := 0; i < elems; i++ {
			if s.g[i] == f {
				fmt.Printf("%dG ", i)
			} else {
				fmt.Printf(".  ")
			}
			if s.m[i] == f {
				fmt.Printf("%dM ", i)
			} else {
				fmt.Printf(".  ")
			}
		}
		fmt.Println()
	}
	fmt.Println("---")
}

func (s state) final() bool {
	if s.el != 4 {
		return false
	}

	for i := 0; i < max; i++ {
		if g := s.g[i]; g != 4 && g != 0 {
			return false
		}
		if m := s.m[i]; m != 4 && m != 0 {
			return false
		}
	}

	return true
}

func (s state) possible() bool {
	for i := 0; i < max; i++ {
		if s.m[i] == s.g[i] {
			continue
		}

		for j := 0; j < max; j++ {
			if i != j && s.m[i] == s.g[j] {
				return false
			}
		}
	}

	return true
}

type item struct {
	chip  bool
	index int
}

func (i item) get(s *state) byte {
	if i.chip {
		return s.m[i.index]
	} else {
		return s.g[i.index]
	}
}

func (i item) set(s *state, f byte) {
	if i.chip {
		s.m[i.index] = f
	} else {
		s.g[i.index] = f
	}
}

func (s state) move(from, to byte, i1, i2 item) (state, bool) {
	n := s
	n.el = to
	moved := false

	if i1.get(&s) == from {
		i1.set(&n, to)
		moved = true
	}
	if i2.get(&s) == from {
		i2.set(&n, to)
		moved = true
	}

	if !moved || !n.possible() {
		return state{}, false
	}

	return n, true
}

func (s state) moveAll(from, to byte, elems int) []state {
	ns := []state(nil)
	for i := 0; i < elems; i++ {
		for j := i; j < elems; j++ {
			if s2, ok := s.move(from, to, item{false, i}, item{false, j}); ok {
				ns = append(ns, s2)
			}
			if s2, ok := s.move(from, to, item{false, i}, item{true, j}); ok {
				ns = append(ns, s2)
			}
			if s2, ok := s.move(from, to, item{true, i}, item{false, j}); ok {
				ns = append(ns, s2)
			}
			if s2, ok := s.move(from, to, item{true, i}, item{true, j}); ok {
				ns = append(ns, s2)
			}
		}
	}
	return ns
}

func (s state) nextStates(elems int) []state {
	ns := []state(nil)
	if s.el < 4 {
		ns = append(ns, s.moveAll(s.el, s.el+1, elems)...)
	}
	if s.el > 1 {
		ns = append(ns, s.moveAll(s.el, s.el-1, elems)...)
	}
	return ns
}

type step struct {
	count int
	state state
	prev  *step
}

func (s *step) print(elems int) {
	if s.prev != nil {
		s.prev.print(elems)
	}
	s.state.print(elems)
}

func solve(p *Problem) {
	s0, elems := parseInitState(p)
	p.PartOne(search(s0, elems))

	s0.g[elems] = 1
	s0.m[elems] = 1
	s0.g[elems+1] = 1
	s0.m[elems+1] = 1
	if s0.possible() {
		p.PartTwo(search(s0, elems+2))
	}
}

func search(s0 state, elems int) int {
	v := map[state]struct{}{}
	q := Queue[*step]{}
	q.Push(&step{0, s0, nil})

	for q.Len() > 0 {
		s := q.Pop()
		if s.state.final() {
			//s.print(elems)
			return s.count
		}

		for _, ns := range s.state.nextStates(elems) {
			if _, ok := v[ns]; ok {
				continue
			}

			v[ns] = struct{}{}
			q.Push(&step{s.count + 1, ns, s})
		}
	}
	panic("no path found")
}
