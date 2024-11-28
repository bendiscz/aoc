package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 12
	example = `

0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5

`
)

func main() {
	Run(year, day, example, solve)
}

type graph map[int][]int

func solve(p *Problem) {
	g := graph{}
	for p.NextLine() {
		f := ParseInts(p.Line())
		g[f[0]] = f[1:]
	}

	p.PartOne(fill(g, 0))

	groups := 1
	for len(g) > 0 {
		for start, _ := range g {
			fill(g, start)
			groups++
			break
		}
	}

	p.PartTwo(groups)
}

func fill(g graph, start int) int {
	s := Stack[int]{}
	count := 0

	s.Push(start)
	for s.Len() > 0 {
		x := s.Pop()
		if _, ok := g[x]; !ok {
			continue
		}

		count++
		for _, y := range g[x] {
			if _, ok := g[y]; ok {
				s.Push(y)
			}
		}

		delete(g, x)
	}

	return count
}
