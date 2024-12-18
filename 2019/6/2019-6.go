package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2019
	day     = 6
	example = `

COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
I)SAN
K)YOU

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	children := map[string][]string{}
	parent := map[string]string{}
	for p.NextLine() {
		f := SplitFieldsDelim(p.Line(), ")")
		children[f[0]] = append(children[f[0]], f[1])
		parent[f[1]] = f[0]
	}
	p.PartOne(count(children, "COM", 0))

	path1 := findPath(parent, "YOU")
	path2 := findPath(parent, "SAN")

	l1, l2, d := len(path1), len(path2), 1
	for path1[l1-d] == path2[l2-d] {
		d++
	}
	p.PartTwo(l1 + l2 - 2*d)
}

func count(orbit map[string][]string, at string, c int) int {
	s := c
	for _, e := range orbit[at] {
		s += count(orbit, e, c+1)
	}
	return s
}

func findPath(parent map[string]string, to string) []string {
	path := []string(nil)
	for at := to; at != "COM"; at = parent[at] {
		path = append(path, at)
	}
	return path
}
