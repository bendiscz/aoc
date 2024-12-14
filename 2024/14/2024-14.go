package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2024
	day     = 14
	example = ``
)

func main() {
	Run(year, day, example, solve)
}

type robot struct {
	p, v XY
}

func solve(p *Problem) {
	X, Y := 101, 103
	q := [4]int{}
	rs := []robot(nil)
	for p.NextLine() {
		f := ParseInts(p.Line())
		rs = append(rs, robot{p: XY{f[0], f[1]}, v: XY{f[2], f[3]}})

		x := Mod(f[0]+100*f[2], X)
		y := Mod(f[1]+100*f[3], Y)

		if x == X/2 || y == Y/2 {
			continue
		}

		i := 0
		if x > X/2 {
			i += 1
		}
		if y >= Y/2 {
			i += 2
		}
		q[i]++
	}
	p.PartOne(q[0] * q[1] * q[2] * q[3])

	for s2 := 0; s2 < 100_000; s2++ {
		g := NewMatrix[bool](Rectangle(X, Y))
		for _, r := range rs {
			x := Mod(r.p.X+s2*r.v.X, X)
			y := Mod(r.p.Y+s2*r.v.Y, Y)
			*g.AtXY(x, y) = true
		}
		for y := 0; y < Y; y++ {
			if findLine(g.Row(y)) >= 14 {
				p.PartTwo(s2)
				PrintGridBool(g)
				return
			}
		}
	}
}

func findLine(v Vector[bool]) int {
	l, l0 := 0, 0
	for x := 0; x < v.Len(); x++ {
		if *v.At(x) {
			l++
			l0 = max(l0, l)
		} else {
			l = 0
		}
	}
	return l0
}
