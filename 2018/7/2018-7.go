package main

import (
	"cmp"
	"math"

	. "github.com/bendiscz/aoc/aoc"
	"golang.org/x/exp/slices"
)

const (
	year    = 2018
	day     = 7
	example = `

Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.

`
)

func main() {
	Run(year, day, example, solve)
}

type step struct {
	id   rune
	next []*step
	inc1 int
	inc2 int
}

func (s *step) duration(base int) int {
	return int(s.id) - 'A' + 1 + base
}

type stepList [128]*step

func (l *stepList) get(id rune) *step {
	s := l[id]
	if s == nil {
		s = &step{id: id}
		l[id] = s
	}
	return s
}

type work struct {
	s     *step
	until int
}

func solve(p *Problem) {
	steps := stepList{}
	for p.NextLine() {
		var id1, id2 rune
		p.Scanf("Step %c must be finished before step %c can begin.", &id1, &id2)
		s1, s2 := steps.get(id1), steps.get(id2)
		s1.next = append(s1.next, s2)
		s2.inc1++
		s2.inc2++
	}

	h := NewHeap[*step](func(s1 *step, s2 *step) bool {
		return s1.id < s2.id
	})
	for _, s := range steps {
		if s != nil && s.inc1 == 0 {
			h.Push(s)
		}
	}

	r1 := []rune(nil)
	for h.Len() > 0 {
		s := h.Pop()
		r1 = append(r1, s.id)

		for _, n := range s.next {
			n.inc1--
			if n.inc1 == 0 {
				h.Push(n)
			}
		}
	}
	p.PartOne(string(r1))

	workers, base := 5, 60
	if p.Example() {
		workers, base = 2, 0
	}

	q := []work(nil)
	for _, s := range steps {
		if s != nil && s.inc2 == 0 {
			w := work{s: s, until: math.MaxInt}
			if len(q) < workers {
				w.until = w.s.duration(base)
			}
			q = append(q, w)
		}
	}

	t := 0
	for len(q) > 0 {
		w := q[0]
		copy(q, q[1:])
		q = q[:len(q)-1]

		t = w.until

		for _, n := range w.s.next {
			n.inc2--
			if n.inc2 == 0 {
				q = append(q, work{s: n, until: math.MaxInt})
			}
		}

		wq, needsSort := q[:min(workers, len(q))], false
		for i := range wq {
			if wq[i].until == math.MaxInt {
				wq[i].until = t + wq[i].s.duration(base)
				needsSort = true
			}
		}

		if needsSort {
			slices.SortFunc(wq, func(a, b work) int {
				return cmp.Compare(a.until, b.until)
			})
		}
	}
	p.PartTwo(t)
}
