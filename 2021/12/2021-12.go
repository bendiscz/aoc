package main

import (
	"strings"
	"unicode"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2021
	day     = 12
	example = `

dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc

`
)

func main() {
	Run(year, day, example, solve)
}

type path struct {
	part    int
	visited map[string]bool
	second  string
}

func newPath(part int) *path {
	return &path{
		part:    part,
		visited: map[string]bool{},
	}
}

func (p *path) enter(n string) bool {
	if !small(n) {
		return true
	}

	if p.visited[n] {
		if p.part == 1 {
			return false
		}
		if n == "start" || n == "end" || len(p.second) > 0 {
			return false
		}
		p.second = n
	} else {
		p.visited[n] = true
	}
	return true
}

func (p *path) leave(n string) {
	if n == p.second {
		p.second = ""
	} else {
		delete(p.visited, n)
	}
}

func small(s string) bool {
	for _, r := range s {
		return unicode.IsLower(r)
	}
	return false
}

func addEdge(edges map[string][]string, m, n string) {
	e := edges[m]
	e = append(e, n)
	edges[m] = e
}

func travel(edges map[string][]string, p *path, n string) int {
	if n == "end" {
		return 1
	}

	if !p.enter(n) {
		return 0
	}
	defer p.leave(n)

	sum := 0
	for _, m := range edges[n] {
		sum += travel(edges, p, m)
	}
	return sum
}

func solve(p *Problem) {
	edges := map[string][]string{}
	for p.NextLine() {
		line := p.Line()
		p := strings.IndexByte(line, '-')
		n1, n2 := line[:p], line[p+1:]
		addEdge(edges, n1, n2)
		addEdge(edges, n2, n1)
	}

	p.PartOne(travel(edges, newPath(1), "start"))
	p.PartTwo(travel(edges, newPath(2), "start"))
}
