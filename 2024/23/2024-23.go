package main

import (
	"cmp"
	. "github.com/bendiscz/aoc/aoc"
	"maps"
	"slices"
	"strings"
)

const (
	year    = 2024
	day     = 23
	example = `

kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn

`
)

func main() {
	Run(year, day, example, solve)
}

type comp struct {
	id   string
	conn Set[string]
}

func addComp(comps map[string]*comp, c1, c2 string) {
	c, ok := comps[c1]
	if !ok {
		c = &comp{c1, Set[string]{}}
		comps[c1] = c
	}
	c.conn[c2] = SET
}

func solve(p *Problem) {
	comps := map[string]*comp{}
	for p.NextLine() {
		f := SplitFieldsDelim(p.Line(), "-")
		addComp(comps, f[0], f[1])
		addComp(comps, f[1], f[0])
	}

	s1 := 0
	for _, c := range comps {
		for c1 := range c.conn {
			if cmp.Compare(c1, c.id) < 0 {
				continue
			}
			for c2 := range c.conn {
				if cmp.Compare(c2, c1) < 0 || !comps[c1].conn.Contains(c2) {
					continue
				}
				if c.id[0] == 't' || c1[0] == 't' || c2[0] == 't' {
					s1++
				}
			}
		}
	}
	p.PartOne(s1)

	r := bronKerbosch(comps, nil, slices.Collect(maps.Keys(comps)), nil, 0)
	slices.Sort(r)
	p.PartTwo(strings.Join(r, ","))
}

func bronKerbosch(comps map[string]*comp, clique, p, x []string, limit int) []string {
	if len(p) == 0 && len(x) == 0 {
		return clique
	}

	if len(p) == 0 || len(p)+len(clique) < limit {
		if len(x) == 0 {
			return clique
		}
		return nil
	}

	maxClique := []string(nil)
	for _, v := range p {
		nextClique := append(slices.Clone(clique), v)
		pn := slices.DeleteFunc(slices.Clone(p), func(i string) bool {
			return !comps[v].conn.Contains(i)
		})
		xn := slices.DeleteFunc(slices.Clone(x), func(i string) bool {
			return !comps[v].conn.Contains(i)
		})
		if c := bronKerbosch(comps, nextClique, pn, xn, len(maxClique)); len(c) > len(maxClique) {
			maxClique = c
		}
		p = p[1:]
		x = append(x, v)
	}
	return maxClique
}
