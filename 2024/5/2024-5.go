package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"slices"
)

const (
	year    = 2024
	day     = 5
	example = `

47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	pairs := map[int][]int{}
	for p.NextLine() {
		if p.Line() == "" {
			break
		}
		var a, b int
		p.Scanf("%d|%d", &a, &b)
		pairs[a] = append(pairs[a], b)
	}

	cmp := func(a, b int) int {
		switch {
		case slices.Contains(pairs[a], b):
			return -1
		case slices.Contains(pairs[b], a):
			return 1
		default:
			return 0
		}
	}

	s1, s2 := 0, 0
	for p.NextLine() {
		f := ParseInts(p.Line())
		if slices.IsSortedFunc(f, cmp) {
			s1 += f[len(f)/2]
		} else {
			slices.SortStableFunc(f, cmp)
			s2 += f[len(f)/2]
		}
	}

	p.PartOne(s1)
	p.PartTwo(s2)
}
