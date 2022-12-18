package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2022
	day     = 18
	example = `

2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5

`
)

func main() {
	Run(year, day, example, solve)
}

const size = 20

const (
	air byte = iota
	lava
	water
)

type xyz struct {
	x, y, z int
}

type grid [size][size][size]byte

func (g *grid) neigh(x, y, z int, fn func(x, y, z int, inside bool)) {
	fn(x-1, y, z, x > 0)
	fn(x+1, y, z, x < size-1)
	fn(x, y-1, z, y > 0)
	fn(x, y+1, z, y < size-1)
	fn(x, y, z-1, z > 0)
	fn(x, y, z+1, z < size-1)
}

func (g *grid) faces(x, y, z int, inside, outside byte) int {
	if g[x][y][z] != inside {
		return 0
	}

	faces := 0
	g.neigh(x, y, z, func(x, y, z int, inside bool) {
		if !inside || g[x][y][z] == outside {
			faces++
		}
	})
	return faces
}

func (g *grid) count(inside, outside byte) int {
	sum := 0
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			for z := 0; z < size; z++ {
				sum += g.faces(x, y, z, inside, outside)
			}
		}
	}
	return sum
}

func (g *grid) fill(x, y, z int) {
	if g[x][y][z] != air {
		return
	}

	q := Queue[xyz]{}
	q.Push(xyz{x, y, z})
	for q.Len() > 0 {
		c := q.Pop()
		if g[c.x][c.y][c.z] != air {
			continue
		}
		g[c.x][c.y][c.z] = water
		g.neigh(c.x, c.y, c.z, func(x, y, z int, inside bool) {
			if inside {
				q.Push(xyz{x, y, z})
			}
		})
	}
}

func (g *grid) fillAll() {
	for p := 0; p < size; p++ {
		for q := 0; q < size; q++ {
			g.fill(p, q, 0)
			g.fill(p, q, size-1)
			g.fill(p, 0, q)
			g.fill(p, size-1, q)
			g.fill(0, p, q)
			g.fill(size-1, p, q)
		}
	}
}

func solve(p *Problem) {
	g := grid{}
	for p.NextLine() {
		x, y, z := 0, 0, 0
		p.Scanf("%d,%d,%d", &x, &y, &z)
		g[x][y][z] = lava
	}

	p.PartOne(g.count(lava, air))
	g.fillAll()
	p.PartTwo(g.count(lava, water))
}
