package main

import (
	"github.com/bendiscz/aoc/2019/intcode"
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2019
	day     = 15
	example = ``
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	ch byte
	v  bool
	d  int
}

func (c cell) String() string { return string([]byte{c.ch}) }

func solve(p *Problem) {
	prog := intcode.Parse(p)
	g, oxy := explore(prog)
	m, o := CollectGrid(g, func(v byte) cell { return cell{ch: v} })
	o = o.Neg()
	oxy = oxy.Add(o)
	PrintGrid(m)

	q, t := Queue[XY]{}, 0
	q.Push(oxy)
	for q.Len() > 0 {
		at := q.Pop()
		c := m.At(at)
		if at == o {
			p.PartOne(c.d)
		}

		t = max(t, c.d)

		for _, d := range HVDirs {
			n := at.Add(d)
			nc := m.At(n)
			if !nc.v && nc.ch != '#' {
				nc.v = true
				nc.d = m.At(at).d + 1
				q.Push(n)
			}
		}
	}
	p.PartTwo(t)
}

var dir = [...]XY{
	1: NegY,
	2: PosY,
	3: NegX,
	4: PosX,
}

var opposite = [...]int{
	1: 2,
	2: 1,
	3: 4,
	4: 3,
}

func explore(prog *intcode.Program) (map[XY]byte, XY) {
	g, at, oxy := map[XY]byte{}, XY0, XY0
	g[at] = 'o'

	c := prog.Exec()

	type state struct {
		back int
		next int
	}

	stack := Stack[state]{}
	stack.Push(state{next: 1})
	for {
		s := stack.Pop()
		if s.next > 4 {
			if stack.Len() == 0 {
				break
			}
			c.ReadWrite(s.back)
			at = at.Add(dir[s.back])
		} else {
			d := s.next
			to := at.Add(dir[d])
			s.next++
			stack.Push(s)

			if _, ok := g[to]; !ok {
				switch c.ReadWrite(d)[0] {
				case 0:
					g[to] = '#'
				case 1:
					at = to
					g[at] = '.'
					stack.Push(state{back: opposite[d], next: 1})
				case 2:
					at = to
					g[at] = 'X'
					oxy = at
					stack.Push(state{back: opposite[d], next: 1})
				default:
					panic("invalid status!")
				}
			}
		}
	}

	return g, oxy
}
