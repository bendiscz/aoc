package main

import (
	"cmp"
	"math"
	"slices"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2015
	day     = 24
	example = `

1
2
3
4
5
7
8
9
10
11

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	weights := ParseInts(p.ReadAll())
	sum := 0
	for _, w := range weights {
		sum += w
	}

	slices.SortFunc(weights, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	r1 := bestResult{n: math.MaxInt, qe: math.MaxInt}
	search1(&r1, weights, sum/3, 3, 1, selected{qe: 1})
	p.PartOne(r1.qe)

	r2 := bestResult{n: math.MaxInt, qe: math.MaxInt}
	search1(&r2, weights, sum/4, 4, 1, selected{qe: 1})
	p.PartTwo(r2.qe)
}

type bestResult struct {
	n, qe int
}

type selected struct {
	m   uint64
	i   int
	sum int
	n   int
	qe  int
}

func search1(r *bestResult, weights []int, targetSum, targetGroups, group int, s selected) {
	for i := s.i; i < len(weights); i++ {
		if s.m&(1<<i) != 0 {
			continue
		}

		ns := selected{
			m:   s.m | (1 << i),
			i:   i + 1,
			sum: s.sum + weights[i],
			n:   s.n,
			qe:  s.qe,
		}
		if group == 1 {
			ns.n++
			ns.qe *= weights[i]
		}

		if ns.n > r.n || ns.n == r.n && ns.qe >= r.qe || ns.sum > targetSum {
			continue
		}

		if ns.sum == targetSum {
			if group == targetGroups-1 {
				if ns.n < r.n || ns.n == r.n && ns.qe < r.qe {
					r.n = ns.n
					r.qe = ns.qe
				}
			} else {
				ns.i = 0
				ns.sum = 0
				search1(r, weights, targetSum, targetGroups, group+1, ns)
			}
		} else {
			search1(r, weights, targetSum, targetGroups, group, ns)
		}
	}
}
