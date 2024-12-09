package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2020
	day     = 3
	example = `

..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	ch byte
}

type grid struct {
	*Matrix[cell]
}

func solve(p *Problem) {
	g := grid{NewMatrix[cell](Rectangle(len(p.PeekLine()), 0))}
	for p.NextLine() {
		ParseVectorFunc(g.AppendRow(), p.Line(), func(b byte) cell { return cell{b} })
	}

	s1 := countTrees(g, 3, 1)
	p.PartOne(s1)
	p.PartTwo(s1 * countTrees(g, 1, 1) * countTrees(g, 5, 1) * countTrees(g, 7, 1) * countTrees(g, 1, 2))
}

func countTrees(g grid, xd, yd int) int {
	s := 0
	for x, y := 0, 0; y < g.Dim.Y; y += yd {
		if g.AtXY(x, y).ch == '#' {
			s++
		}
		x = (x + xd) % g.Dim.X
	}
	return s
}
