package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2023
	day     = 13
	example = `

#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	v bool
}

type grid struct {
	*Matrix[cell]
}

func newGrid(width int) *grid {
	return &grid{
		Matrix: NewMatrix[cell](Rectangle(width, 0)),
	}
}

func (g *grid) append(row string) {
	ParseVectorFunc(g.AppendRow(), row, func(b byte) cell { return cell{b == '#'} })
}

func (g *grid) count(smudge bool) int {
	if y, ok := checkRows(g.Matrix, smudge); ok {
		return 100 * y
	}
	if x, ok := checkRows(g.Matrix.TransView(), smudge); ok {
		return x
	}
	panic("no reflection")
}

func checkRows(m Grid[cell], smudge bool) (int, bool) {
	for y0 := 1; y0 < m.Size().Y; y0++ {
		if checkRow(m, y0, smudge) {
			return y0, true
		}
	}
	return 0, false
}

func checkRow(m Grid[cell], y0 int, smudge bool) bool {
	found := false
	for d := 0; d < min(y0, m.Size().Y-y0); d++ {
		y1, y2 := y0-d-1, y0+d
		for x := 0; x < m.Size().X; x++ {
			if m.AtXY(x, y1).v != m.AtXY(x, y2).v {
				if !smudge || found {
					return false
				}
				found = true
			}
		}
	}

	return !smudge || found
}

func solve(p *Problem) {
	grids := readGrids(p)

	sum1, sum2 := 0, 0
	for _, g := range grids {
		sum1 += g.count(false)
		sum2 += g.count(true)
	}
	p.PartOne(sum1)
	p.PartTwo(sum2)
}

func readGrids(p *Problem) []*grid {
	grids := []*grid(nil)
	g := (*grid)(nil)
	for p.NextLine() {
		if p.Line() == "" {
			grids = append(grids, g)
			g = nil
		} else {
			if g == nil {
				g = newGrid(len(p.Line()))
			}
			g.append(p.Line())
		}
	}
	if g != nil {
		grids = append(grids, g)
	}
	return grids
}
