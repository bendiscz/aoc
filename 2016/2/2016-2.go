package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2016
	day     = 2
	example = `

ULL
RRDDD
LURDL
UUUUD

`
)

func main() {
	Run(year, day, example, solve)
}

var (
	keypad1 = keypad{
		{'1', '2', '3'},
		{'4', '5', '6'},
		{'7', '8', '9'},
	}

	keypad2 = keypad{
		{0, 0, '1', 0, 0},
		{0, '2', '3', '4', 0},
		{'5', '6', '7', '8', '9'},
		{0, 'A', 'B', 'C', 0},
		{0, 0, 'D', 0, 0},
	}

	dirs = map[byte]XY{
		'R': PosX,
		'L': NegX,
		'D': PosY,
		'U': NegY,
	}
)

type keypad [][]byte

func (k keypad) key(c XY) byte {
	return k[c.Y][c.X]
}

func (k keypad) move(c XY, d byte) XY {
	dim := XY{len(k[0]), len(k)}
	if a := c.Add(dirs[d]); dim.HasInside(a) && k.key(a) != 0 {
		c = a
	}
	return c
}

func solveKeypad(p *Problem, start XY, k keypad) string {
	c := start
	var code []byte

	for p.NextLine() {
		for _, d := range []byte(p.Line()) {
			c = k.move(c, d)
		}
		code = append(code, k.key(c))
	}

	return string(code)
}

func solve(p *Problem) {
	p.PartOne(solveKeypad(p, XY{1, 1}, keypad1))
	p.Reset()
	p.PartTwo(solveKeypad(p, XY{0, 2}, keypad2))
}
