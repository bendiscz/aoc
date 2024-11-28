package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2023
	day     = 8
	example = `

LR

AAA = (BBB, AAA)
BBB = (AAA, CCC)
CCC = (ZZZ, AAA)
ZZZ = (ZZZ, ZZZ)
11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)

`
)

func main() {
	//Run(year, day, "", solve)
	Run(year, day, example, solve)
}

type edge struct {
	l, r string
}

type graph map[string]edge

func (g graph) move(c string, d byte) string {
	n := g[c]
	if d == 'L' {
		return n.l
	} else {
		return n.r
	}
}

func solve(p *Problem) {
	path := p.ReadLineBytes()
	g := graph{}

	p.NextLine()
	for p.NextLine() {
		f := SplitFieldsDelim(p.Line(), " =(,)")
		g[f[0]] = edge{f[1], f[2]}
	}

	p.PartOne(part1(path, g))
	p.PartTwo(part2(path, g))
}

func part1(path []byte, g graph) int {
	c, steps := "AAA", 0
	for c != "ZZZ" {
		c = g.move(c, path[steps%len(path)])
		steps++
	}
	return steps
}

func part2(path []byte, g graph) int {
	cs := []int(nil)
	for n, _ := range g {
		if n[2] == 'A' {
			cs = append(cs, detectCycle(g, n, path))
		}
	}

	r := 1
	for _, c := range cs {
		r = LCM(r, c)
	}
	return r
}

func detectCycle(g graph, c string, path []byte) int {
	for i := 0; ; i++ {
		c = g.move(c, path[i%len(path)])
		if c[2] == 'Z' {
			return i + 1
		}
	}
}
