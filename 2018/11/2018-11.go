package main

import (
	"fmt"
	"math"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2018
	day     = 11
	example = `

18
42

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	level int
}

type grid struct {
	*Matrix[cell]
}

func solve(p *Problem) {
	for p.NextLine() {
		serial := ParseInt(p.Line())
		g := grid{NewMatrix[cell](Square(300))}
		for x := 0; x < g.Dim.X; x++ {
			id := x + 11
			for y := 0; y < g.Dim.Y; y++ {
				level := id * (y + 1)
				level += serial
				level *= id
				level = (level/100)%10 - 5
				g.AtXY(x, y).level = level
			}
		}

		level0, xy0 := math.MinInt, XY{}
		cg := condense(g, 3)
		cg.Dim.ForEach(func(xy XY) {
			level := cg.At(xy).level
			if level > level0 {
				level0 = level
				xy0 = xy
			}
		})
		p.PartOne(fmt.Sprintf("%d,%d", xy0.X+1, xy0.Y+1))

		level0, xy0, d0 := math.MinInt, XY{}, 0
		for d := 1; d <= g.Dim.X; d++ {
			cg = condense(g, d)
			cg.Dim.ForEach(func(xy XY) {
				level := cg.At(xy).level
				if level > level0 {
					level0 = level
					xy0 = xy
					d0 = d
				}
			})
		}
		p.PartTwo(fmt.Sprintf("%d,%d,%d", xy0.X+1, xy0.Y+1, d0))
	}

}

func condense(g grid, d int) grid {
	return condenseX(condenseX(g.TransView(), d).TransView(), d)
}

func condenseX(g Grid[cell], d int) grid {
	cg := grid{NewMatrix[cell](g.Size().Sub(XY{d - 1, 0}))}
	for y := 0; y < cg.Dim.Y; y++ {
		level := 0
		for x := 0; x < d; x++ {
			level += g.AtXY(x, y).level
		}

		cg.AtXY(0, y).level = level
		for x := 1; x < cg.Dim.X; x++ {
			level += g.AtXY(x+d-1, y).level - g.AtXY(x-1, y).level
			cg.AtXY(x, y).level = level
		}
	}
	return cg
}
