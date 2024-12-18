package main

import (
	"github.com/bendiscz/aoc/2019/intcode"
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2019
	day     = 13
	example = ``
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	s1, s2 := 0, 0
	prog := intcode.Parse(p)

	screen := map[XY]int{}
	c := prog.Exec()

	for {
		out := c.ReadInts(3)
		if len(out) == 0 {
			break
		}
		x, y, id := out[0], out[1], out[2]
		screen[XY{x, y}] = id
		if id == 2 {
			s1++
		}
	}
	p.PartOne(s1)

	g, _ := CollectGrid(screen, func(v int) ByteCell {
		switch v {
		case 1:
			return ByteCell{V: '+'}
		case 2:
			return ByteCell{V: '#'}
		case 3:
			return ByteCell{V: '_'}
		case 4:
			return ByteCell{V: 'o'}
		default:
			return ByteCell{V: '.'}
		}
	})
	PrintGrid(g)

	c = prog.Exec()
	c.Poke(0, 2)

	ball, paddle := XY{}, XY{}
	for {
		out := c.ReadInts(3)
		if len(out) == 0 {
			break
		}
		x, y, id := out[0], out[1], out[2]

		switch {
		case x == -1 && y == 0:
			s2 = id
		case id == 3:
			paddle = XY{x, y}
		case id == 4:
			ball = XY{x, y}
		}

		if ball != (XY{}) {
			switch {
			case ball.X < paddle.X:
				c.WriteInt(-1)
			case ball.X > paddle.X:
				c.WriteInt(1)
			default:
				c.WriteInt(0)
			}
			ball = XY{}
		}
	}
	p.PartTwo(s2)
}
