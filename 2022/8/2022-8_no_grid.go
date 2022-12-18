package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

type treeOld struct {
	height int8
	vx, vy bool
}

type grid struct {
	size int
	data [][]treeOld
}

func newGrid(size int) *grid {
	g := &grid{
		size: size,
		data: make([][]treeOld, size),
	}
	for i := range g.data {
		g.data[i] = make([]treeOld, size)
	}
	return g
}

func (g *grid) score(x0, y0 int) int {
	h := g.data[x0][y0].height
	score, count := 1, 0

	for x, stop := x0-1, false; x >= 0 && !stop; x-- {
		count++
		stop = g.data[x][y0].height >= h
	}
	score, count = score*count, 0

	for x, stop := x0+1, false; x < g.size && !stop; x++ {
		count++
		stop = g.data[x][y0].height >= h
	}
	score, count = score*count, 0

	for y, stop := y0-1, false; y >= 0 && !stop; y-- {
		count++
		stop = g.data[x0][y].height >= h
	}
	score, count = score*count, 0

	for y, stop := y0+1, false; y < g.size && !stop; y++ {
		count++
		stop = g.data[x0][y].height >= h
	}
	score = score * count

	return score
}

func solveOld(p *Problem) {
	var g *grid
	for y := 0; p.NextLine(); y++ {
		if g == nil {
			g = newGrid(len(p.Line()))
		}
		for x, ch := range []byte(p.Line()) {
			g.data[x][y].height = int8(ch - '0')
		}
	}

	count := 0
	for y := 0; y < g.size; y++ {
		for x, min := 0, int8(-1); x < g.size; x++ {
			if g.data[x][y].height > min {
				min = g.data[x][y].height
				g.data[x][y].vx = true
				count++
			}
		}
		for x, min := g.size-1, int8(-1); x >= 0 && !g.data[x][y].vx; x-- {
			if g.data[x][y].height > min {
				min = g.data[x][y].height
				g.data[x][y].vx = true
				count++
			}
		}
	}

	for x := 0; x < g.size; x++ {
		for y, min := 0, int8(-1); y < g.size; y++ {
			if g.data[x][y].height > min {
				min = g.data[x][y].height
				g.data[x][y].vy = true
				if !g.data[x][y].vx {
					count++
				}
			}
		}
		for y, min := g.size-1, int8(-1); y >= 0 && !g.data[x][y].vy; y-- {
			if g.data[x][y].height > min {
				min = g.data[x][y].height
				g.data[x][y].vy = true
				if !g.data[x][y].vx {
					count++
				}
			}
		}
	}

	p.PartOne(count)

	max := 0
	for x := 0; x < g.size; x++ {
		for y := 0; y < g.size; y++ {
			score := g.score(x, y)
			if score > max {
				max = score
			}
		}
	}

	p.PartTwo(max)
}
