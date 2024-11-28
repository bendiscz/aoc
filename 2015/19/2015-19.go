package main

import (
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2015
	day     = 19
	example = `

e => H
e => O
H => HO
H => OH
O => HH

HOHOHO


`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	rules := map[string][]string{}
	for p.NextLine() && len(p.Line()) > 0 {
		f := SplitFields(p.Line())
		rules[f[0]] = append(rules[f[0]], f[2])
	}
	medicine := p.ReadLine()

	p.PartOne(part1(medicine, rules))
	p.PartTwo(part2(medicine, rules))
}

func part1(medicine string, rules map[string][]string) int {
	variants := map[string]bool{}
	for i := 0; i < len(medicine); i++ {
		for k, v := range rules {
			if strings.HasPrefix(medicine[i:], k) {
				for _, r := range v {
					m := medicine[:i] + r + medicine[i+len(k):]
					variants[m] = true
				}
			}
		}
	}
	return len(variants)
}

func part2(medicine string, rules map[string][]string) int {
	r := map[string][]string{}
	for k, v := range rules {
		for _, x := range v {
			r[x] = append(r[x], k)
		}
	}

	type path struct {
		molecule string
		steps    int
	}

	s := Stack[path]{}
	s.Push(path{medicine, 0})

	seen := map[string]bool{}

	for {
		p := s.Pop()
		if p.molecule == "e" {
			return p.steps
		}

		for i := 0; i < len(p.molecule); i++ {
			for k, v := range r {
				if strings.HasPrefix(p.molecule[i:], k) {
					for _, r := range v {
						m := p.molecule[:i] + r + p.molecule[i+len(k):]
						if seen[m] {
							continue
						}
						seen[m] = true
						s.Push(path{m, p.steps + 1})
					}
				}
			}
		}

	}
}
