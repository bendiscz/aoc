package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2022
	day     = 8
	example = `

30373
25512
65332
33549
35390

`
)

func main() {
	Run(year, day, example, solve)
	//Run(year, day, example, solveOld)
}

type tree struct {
	height int8
	//vx, vy  bool
	visible byte
}

func solve(p *Problem) {
	var m *Matrix[tree]
	for p.NextLine() {
		if m == nil {
			m = NewMatrix[tree](Rectangle(len(p.Line()), 0))
		}
		row := m.AppendRow()
		ParseVectorFunc(row, p.Line(), func(ch byte) tree { return tree{height: int8(ch - '0')} })
	}

	p.PartOne(countRows(m, 1) + countRows(m.Trans(), 2))

	max := 0
	for x := 0; x < m.Dim.X; x++ {
		for y := 0; y < m.Dim.Y; y++ {
			c := XY{x, y}
			score := countTrees(m, c, PosX) * countTrees(m, c, NegX) * countTrees(m, c, PosY) * countTrees(m, c, NegY)
			if score > max {
				max = score
			}
		}
	}
	p.PartTwo(max)
}

func countRows(g Grid[tree], mask byte) (count int) {
	for y := 0; y < g.Size().Y; y++ {
		row := g.Row(y)
		for x, min := 0, int8(-1); x < row.Len(); x++ {
			if row.At(x).height > min {
				min = row.At(x).height
				row.At(x).visible |= mask
				if row.At(x).visible & ^mask == 0 {
					count++
				}
			}
		}
		for x, min := row.Len()-1, int8(-1); x >= 0 && row.At(x).visible&mask == 0; x-- {
			if row.At(x).height > min {
				min = row.At(x).height
				row.At(x).visible |= mask
				if row.At(x).visible & ^mask == 0 {
					count++
				}
			}
		}
	}
	return
}

func countTrees(g Grid[tree], c, d XY) (count int) {
	h := g.At(c).height
	for {
		c = c.Add(d)
		if !g.Size().HasInside(c) {
			return
		}

		count++

		if g.At(c).height >= h {
			return
		}
	}
}
