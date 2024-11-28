package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 22
	example = `

..#
#..
...

`
)

func main() {
	Run(year, day, example, solve)
}

type state int

const (
	clean state = iota
	weakened
	infected
	flagged
)

func solve(p *Problem) {
	g1 := map[XY]struct{}{}
	g2 := map[XY]state{}
	y, w := 0, 0
	for p.NextLine() {
		s := p.Line()
		w = max(w, len(s))
		for x := 0; x < len(s); x++ {
			if s[x] == '#' {
				g1[XY{x, y}] = struct{}{}
				g2[XY{x, y}] = infected
			}
		}
		y++
	}

	start := XY{w / 2, y / 2}

	sum1, xy, d := 0, start, NegY
	for i := 0; i < 10_000; i++ {
		if _, ok := g1[xy]; ok {
			delete(g1, xy)
			d = right(d)
		} else {
			g1[xy] = struct{}{}
			d = left(d)
			sum1++
		}
		xy = xy.Add(d)
	}
	p.PartOne(sum1)

	sum2, xy, d := 0, start, NegY
	for i := 0; i < 10_000_000; i++ {
		switch g2[xy] {
		case clean:
			d = left(d)
			g2[xy] = weakened

		case weakened:
			g2[xy] = infected
			sum2++

		case infected:
			d = right(d)
			g2[xy] = flagged

		case flagged:
			d = d.Neg()
			delete(g2, xy)
		}
		xy = xy.Add(d)
	}
	p.PartTwo(sum2)
}

func left(d XY) XY  { return XY{d.Y, -d.X} }
func right(d XY) XY { return XY{-d.Y, d.X} }
