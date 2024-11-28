package main

import (
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2016
	day     = 1
	example = `

R2, L3
R2, R2, R2
R5, L5, R5, R3
R8, R4, R4, R8

`
)

func main() {
	Run(year, day, example, solve)
}

var dirs = [...]XY{PosY, PosX, NegY, NegX}

func dist(c XY) int {
	return Abs(c.X) + Abs(c.Y)
}

func solvePath(p *Problem, s string) {
	dir, c, t, visited, found := 0, XY{}, XY{}, map[XY]bool{}, false
	visited[c] = true

	for _, op := range strings.Fields(p.Line()) {
		switch op[0] {
		case 'R':
			dir = (dir + 1) % 4
		case 'L':
			dir = (dir + 3) % 4
		}

		op = strings.TrimRight(op[1:], ",")
		for i, n, d := 0, ParseInt(op), dirs[dir]; i < n; i++ {
			c = c.Add(d)
			if !found {
				if visited[c] {
					t = c
					found = true
				}
				visited[c] = true
			}
		}
	}

	p.PartOne(dist(c))
	p.PartTwo(dist(t))
}

func solve(p *Problem) {
	for i := 1; p.NextLine(); i++ {
		p.Printf("path #%d", i)
		solvePath(p, p.Line())
	}
}
