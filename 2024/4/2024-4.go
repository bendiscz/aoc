package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2024
	day     = 4
	example = `

MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`
)

func main() {
	Run(year, day, example, solve)
}

type grid struct {
	*Matrix[ByteCell]
}

func solve(p *Problem) {
	g := grid{NewMatrix[ByteCell](Rectangle(len(p.PeekLine()), 0))}
	for p.NextLine() {
		ParseVectorFunc(g.AppendRow(), p.Line(), func(b byte) ByteCell { return ByteCell{b} })
	}

	p.PartOne(countXmas(g))
	p.PartTwo(countCross(g))

}

func countXmas(g grid) int {
	s := 0
	g.Dim.ForEach(func(xy XY) {
		for _, d := range AllDirs {
			if x, ok := extractWord(g, xy, d); ok && x == "XMAS" {
				s++
			}
		}
	})
	return s
}

func extractWord(g grid, xy XY, d XY) (string, bool) {
	s := make([]byte, 4)
	for i := 0; i < 4; i++ {
		if !g.Dim.HasInside(xy) {
			return "", false
		}
		s[i] = g.At(xy).V
		xy = xy.Add(d)
	}
	return string(s), true
}

func countCross(g grid) int {
	s := 0
	for x := 1; x < g.Size().X-1; x++ {
		for y := 1; y < g.Size().Y-1; y++ {
			if checkCross(g, XY{x, y}) {
				s++
			}
		}
	}
	return s
}

func checkCross(g grid, xy XY) bool {
	return g.At(xy).V == 'A' &&
		checkMS(g.At(xy.Add(Q1)).V, g.At(xy.Add(Q3)).V) &&
		checkMS(g.At(xy.Add(Q2)).V, g.At(xy.Add(Q4)).V)
}

func checkMS(b1, b2 byte) bool {
	return b1 == 'M' && b2 == 'S' || b1 == 'S' && b2 == 'M'
}
