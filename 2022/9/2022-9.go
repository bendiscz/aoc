package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2022
	day     = 9
	example = `

R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2

`
)

func main() {
	Run(year, day, example, solve)
}

var dir = map[string]XY{
	"R": PosX,
	"L": NegX,
	"U": PosY,
	"D": NegY,
}

type rope []XY

func (r rope) move(d XY) XY {
	r[0] = r[0].Add(d)
	for i := 1; i < len(r); i++ {
		r[i] = follow(r[i-1], r[i])
	}
	return r[len(r)-1]
}

func follow(h, t XY) XY {
	dx, dy := h.X-t.X, h.Y-t.Y
	if Abs(dx) <= 1 && Abs(dy) <= 1 {
		return t
	}

	t.X += Sign(dx)
	t.Y += Sign(dy)
	return t
}

func solve(p *Problem) {
	visited2, visited10 := map[XY]struct{}{}, map[XY]struct{}{}
	rope2, rope10 := rope(make([]XY, 2)), rope(make([]XY, 10))

	for p.NextLine() {
		m := p.Parse(`^(.) (\d+)$`)
		for d, n := dir[m[1]], ParseInt(m[2]); n > 0; n-- {
			visited2[rope2.move(d)] = struct{}{}
			visited10[rope10.move(d)] = struct{}{}
		}
	}

	p.PartOne(len(visited2))
	p.PartTwo(len(visited10))
}
