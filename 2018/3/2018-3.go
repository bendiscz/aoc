package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2018
	day     = 3
	example = `

#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	v int
}

type rect struct {
	id     int
	p1, p2 XY
}

func solve(p *Problem) {
	g := NewMatrix[cell](Square(1000))
	rectangles := []rect(nil)
	for p.NextLine() {
		f := ParseInts(p.Line())
		rectangles = append(rectangles, rect{f[0], XY{f[1], f[2]}, XY{f[1] + f[3], f[2] + f[4]}})
		for x := f[1]; x < f[1]+f[3]; x++ {
			for y := f[2]; y < f[2]+f[4]; y++ {
				g.AtXY(x, y).v++
			}
		}
	}

	count := 0
	g.Dim.ForEach(func(xy XY) {
		if g.At(xy).v > 1 {
			count++
		}
	})
	p.PartOne(count)

loop:
	for _, r := range rectangles {
		for x := r.p1.X; x < r.p2.X; x++ {
			for y := r.p1.Y; y < r.p2.Y; y++ {
				if g.AtXY(x, y).v > 1 {
					continue loop
				}
			}
		}
		p.PartTwo(r.id)
		break
	}
}
