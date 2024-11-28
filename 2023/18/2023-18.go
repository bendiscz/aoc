package main

import (
	"strconv"
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2023
	day     = 18
	example = `

R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)

`
)

func main() {
	Run(year, day, example, solve)
}

var dirs = [...]XY{
	0:   PosX,
	1:   PosY,
	2:   NegX,
	3:   NegY,
	'R': PosX,
	'L': NegX,
	'D': PosY,
	'U': NegY,
}

type step struct {
	dir XY
	d   int
}

func (s step) String() string {
	str := strconv.Itoa(s.d)
	switch s.dir {
	case PosX:
		return "R " + str
	case NegX:
		return "L " + str
	case PosY:
		return "D " + str
	case NegY:
		return "U " + str
	}
	panic("dir")
}

func solve(p *Problem) {
	steps1 := []step(nil)
	steps2 := []step(nil)
	for p.NextLine() {
		steps1 = append(steps1, parseStep1(p.Line()))
		steps2 = append(steps2, parseStep2(p.Line()))
	}

	p.PartOne(compute(steps1))
	p.PartTwo(compute(steps2))
}

func parseStep1(s string) step {
	i := strings.IndexByte(s, '(')
	return step{dirs[s[0]], ParseInt(s[2 : i-1])}
}

func parseStep2(s string) step {
	i := strings.IndexByte(s, '(')
	x, _ := strconv.ParseInt(s[i+2:len(s)-1], 16, 64)
	st := step{dirs[x%16], int(x) / 16}
	return st
}

func compute(steps []step) int {
	area := 0
	for b, i := (XY{}), 0; i < len(steps); i++ {
		ps, s, ns := steps[(i-1+len(steps))%len(steps)], steps[i], steps[(i+1)%len(steps)]
		d := s.d + (ps.dir.X*s.dir.Y-ps.dir.Y*s.dir.X+s.dir.X*ns.dir.Y-s.dir.Y*ns.dir.X)/2
		b.X += d * s.dir.X
		b.Y += d * s.dir.Y
		area += s.dir.Y * b.X * d
	}
	return area
}
