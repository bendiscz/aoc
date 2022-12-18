package main

import (
	"sort"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2021
	day     = 9
	example = `

2199943210
3987894921
9856789892
8767896789
9899965678

`
)

func main() {
	Run(year, day, example, solve)
}

type elem struct {
	height byte
	mask   bool
}

func solve(p *Problem) {
	var m *Matrix[elem]
	for p.NextLine() {
		if m == nil {
			m = NewMatrix[elem](Rectangle(len(p.Line()), 0))
		}
		row := m.AppendRow()
		ParseVectorFunc(row, p.Line(), func(ch byte) elem { return elem{height: ch - '0'} })
	}

	lows, risk := find(m)

	p.PartOne(risk)

	var basins []int
	for _, low := range lows {
		basins = append(basins, fill(m, low))
	}

	sort.Ints(basins)
	if len(basins) >= 3 {
		basins = basins[len(basins)-3:]
		p.PartTwo(basins[0] * basins[1] * basins[2])
	}
}

func find(m *Matrix[elem]) (lows []XY, risk int) {
	for x := 0; x < m.Dim.X; x++ {
		for y := 0; y < m.Dim.Y; y++ {
			c, l := XY{x, y}, m.AtXY(x, y).height
			if testLow(m, l, c.Add(PosX)) && testLow(m, l, c.Add(NegX)) && testLow(m, l, c.Add(PosY)) && testLow(m, l, c.Add(NegY)) {
				risk += int(l + 1)
				lows = append(lows, c)
			}
		}
	}
	return
}

func testLow(m *Matrix[elem], l byte, c XY) bool {
	return !m.Dim.HasInside(c) || m.At(c).height > l
}

func fill(m *Matrix[elem], low XY) (count int) {
	stack := Stack[XY]{}
	stack.Push(low)
	for stack.Len() > 0 {
		c := stack.Pop()
		if m.At(c).mask || m.At(c).height == 9 {
			continue
		}

		count++
		m.At(c).mask = true

		for _, d := range HVDirs {
			if a := c.Add(d); m.Dim.HasInside(a) {
				stack.Push(a)
			}
		}
	}
	return
}
