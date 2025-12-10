package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2025
	day     = 7
	example = `

.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	ch byte
	x  int
}

func (c cell) String() string { return string(c.ch) }

type grid struct {
	*Matrix[cell]
}

func solve(p *Problem) {
	s1, s2 := 0, 0

	g := &grid{Matrix: NewMatrix[cell](Rectangle(len(p.PeekLine()), 0))}
	for p.NextLine() {
		ParseVectorFunc(g.AppendRow(), p.Line(), func(b byte) cell {
			return cell{ch: b}
		})
	}

	for x := 0; x < g.Dim.X; x++ {
		if g.AtXY(x, 0).ch == 'S' {
			g.AtXY(x, 1).ch = '|'
			g.AtXY(x, 1).x = 1
		}
	}

	for y := 2; y < g.Dim.Y; y++ {
		for x := 0; x < g.Dim.X; x++ {
			if g.AtXY(x, y-1).ch != '|' {
				continue
			}

			if g.AtXY(x, y).ch == '^' {
				s1++
				if x > 0 {
					g.AtXY(x-1, y).ch = '|'
					g.AtXY(x-1, y).x += g.AtXY(x, y-1).x
				}
				if x < g.Dim.X-1 {
					g.AtXY(x+1, y).ch = '|'
					g.AtXY(x+1, y).x += g.AtXY(x, y-1).x
				}

			} else {
				g.AtXY(x, y).ch = '|'
				g.AtXY(x, y).x += g.AtXY(x, y-1).x
			}
		}
	}
	//PrintGrid(g)

	//for y := 0; y < g.Dim.Y; y++ {
	//	for x := 0; x < g.Dim.X; x++ {
	//		if g.AtXY(x, y).x > 0 {
	//			fmt.Printf("%d ", g.AtXY(x, y).x)
	//		} else {
	//			fmt.Printf("  ")
	//		}
	//	}
	//	fmt.Printf("\n")
	//}

	for x := 0; x < g.Dim.X; x++ {
		s2 += g.AtXY(x, g.Dim.Y-1).x
	}

	p.PartOne(s1)
	p.PartTwo(s2)

	//os.Exit(1)
}
