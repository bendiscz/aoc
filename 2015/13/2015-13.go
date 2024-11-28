package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2015
	day     = 13
	example = `

Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.

`
)

func main() {
	Run(year, day, example, solve)
}

type pair struct {
	p1, p2 int
}

func pairOf(p1, p2 int) pair {
	if p1 <= p2 {
		return pair{p1, p2}
	} else {
		return pair{p2, p1}
	}
}

type graph struct {
	names map[string]int
	score map[pair]int
}

func (g *graph) index(s string) int {
	if i, ok := g.names[s]; ok {
		return i
	}
	i := len(g.names)
	g.names[s] = i
	return i
}

func (g *graph) add(n1, n2 string, d int) {
	k := pairOf(g.index(n1), g.index(n2))
	g.score[k] += d
}

func solve(p *Problem) {
	g := graph{map[string]int{}, map[pair]int{}}
	for p.NextLine() {
		m := p.Parse(`^(\w+) would (gain|lose) (\d+) happiness units by sitting next to (\w+).$`)
		d := ParseInt(m[3])
		if m[2] == "lose" {
			d = -d
		}
		g.add(m[1], m[4], d)
	}

	v := make([]int, len(g.names))
	for i := range v {
		v[i] = i
	}

	max1, max2 := 0, 0
	Permutations(v, func(_ []int) {
		score := 0
		for i := 0; i < len(v)-1; i++ {
			score += g.score[pairOf(v[i], v[i+1])]
		}
		max2 = max(max2, score)
		if v[0] == 0 {
			score += g.score[pairOf(v[0], v[len(v)-1])]
			max1 = max(max1, score)
		}
	})

	p.PartOne(max1)
	p.PartTwo(max2)
}
