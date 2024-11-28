package main

import (
	"math"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2018
	day     = 10
	example = `

position=< 9,  1> velocity=< 0,  2>
position=< 7,  0> velocity=<-1,  0>
position=< 3, -2> velocity=<-1,  1>
position=< 6, 10> velocity=<-2, -1>
position=< 2, -4> velocity=< 2,  2>
position=<-6, 10> velocity=< 2, -2>
position=< 1,  8> velocity=< 1, -1>
position=< 1,  7> velocity=< 1,  0>
position=<-3, 11> velocity=< 1, -2>
position=< 7,  6> velocity=<-1, -1>
position=<-2,  3> velocity=< 1,  0>
position=<-4,  3> velocity=< 2,  0>
position=<10, -3> velocity=<-1,  1>
position=< 5, 11> velocity=< 1, -2>
position=< 4,  7> velocity=< 0, -1>
position=< 8, -2> velocity=< 0,  1>
position=<15,  0> velocity=<-2,  0>
position=< 1,  6> velocity=< 1,  0>
position=< 8,  9> velocity=< 0, -1>
position=< 3,  3> velocity=<-1,  1>
position=< 0,  5> velocity=< 0, -1>
position=<-2,  2> velocity=< 2,  0>
position=< 5, -2> velocity=< 1,  2>
position=< 1,  4> velocity=< 2,  1>
position=<-2,  7> velocity=< 2, -2>
position=< 3,  6> velocity=<-1, -1>
position=< 5,  0> velocity=< 1,  0>
position=<-6,  0> velocity=< 2,  0>
position=< 5,  9> velocity=< 1, -2>
position=<14,  7> velocity=<-2,  0>
position=<-3,  6> velocity=< 2, -1>

`
)

func main() {
	Run(year, day, example, solve)
}

type point struct {
	pos, dir XY
}

func (p point) move() point {
	return point{
		pos: p.pos.Add(p.dir),
		dir: p.dir,
	}
}

func solve(p *Problem) {
	points := []point(nil)
	for p.NextLine() {
		f := ParseInts(p.Line())
		points = append(points, point{
			pos: XY{f[0], f[1]},
			dir: XY{f[2], f[3]},
		})
	}

	var pMin0, pMax0 XY
	steps := 0
	dy0 := math.MaxInt
	np := make([]point, len(points))
	for {
		pMin := XY{math.MaxInt, math.MaxInt}
		pMax := XY{math.MinInt, math.MinInt}
		for i, pt := range points {
			pt = pt.move()
			np[i] = pt
			pMin = XY{min(pMin.X, pt.pos.X), min(pMin.Y, pt.pos.Y)}
			pMax = XY{max(pMax.X, pt.pos.X), max(pMax.Y, pt.pos.Y)}
		}

		dy := pMax.Y - pMin.Y
		if dy > dy0 {
			break
		}

		copy(points, np)
		pMin0, pMax0 = pMin, pMax
		dy0 = dy
		steps++
	}

	g := NewMatrix[bool](pMax0.Sub(pMin0).Add(XY{1, 1}))
	for _, pt := range points {
		*g.At(pt.pos.Sub(pMin0)) = true
	}
	p.PartOne("")
	PrintGridBool(g)

	p.PartTwo(steps)
}
