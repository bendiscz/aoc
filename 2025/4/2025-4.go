package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2025
	day     = 4
	example = `

..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	ch byte
	m  bool
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

	for xy := range g.Dim.All() {
		if g.At(xy).ch != '@' {
			continue
		}

		n := 0
		for _, d := range AllDirs {
			c := xy.Add(d)
			if g.Dim.HasInside(c) && g.At(c).ch == '@' {
				n++
			}
		}

		if n < 4 {
			g.At(xy).m = true
			s1++
			s2++
		}
	}

	done := false
	for !done {
		for xy := range g.Dim.All() {
			if g.At(xy).m {
				g.At(xy).ch = '.'
				g.At(xy).m = false
			}
		}

		done = true
		for xy := range g.Dim.All() {
			if g.At(xy).ch != '@' {
				continue
			}

			n := 0
			for _, d := range AllDirs {
				c := xy.Add(d)
				if g.Dim.HasInside(c) && g.At(c).ch == '@' {
					n++
				}
			}

			if n < 4 {
				g.At(xy).m = true
				s2++
				done = false
			}
		}
	}

	p.PartOne(s1)
	p.PartTwo(s2)
}
