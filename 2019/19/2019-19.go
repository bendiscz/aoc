package main

import (
	"github.com/bendiscz/aoc/2019/intcode"
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2019
	day     = 19
	example = ``
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	ch byte
}

func (c cell) String() string { return string([]byte{c.ch}) }

type grid struct {
	*Matrix[cell]
}

type scan struct {
	y      int
	x0, x1 int
}

func solve(p *Problem) {
	prog := intcode.Parse(p)

	g := grid{NewMatrix[cell](Square(50))}
	s1 := 0
	for xy, c := range g.All() {
		if check(prog, xy) {
			c.ch = '#'
			s1++
		} else {
			c.ch = '.'
		}
	}
	PrintGrid(g)
	p.PartOne(s1)

	line := scan{y: g.Dim.Y - 2}
	for g.AtXY(line.x0, line.y).ch == '.' {
		line.x0++
	}
	line.x1 = line.x0
	for g.AtXY(line.x1, line.y).ch != '.' {
		line.x1++
	}

	lines := [100]scan{}
	for i := range lines {
		line = nextScan(prog, line)
		lines[i] = line
	}

	i := 0
	for {
		j := (i + 1) % len(lines)
		if lines[j].x1-lines[i].x0 >= 100 {
			p.PartTwo(lines[i].x0*10_000 + lines[j].y)
			break
		}
		lines[j] = nextScan(prog, lines[i])
		i = j
	}
}

func check(prog *intcode.Program, xy XY) bool {
	output := prog.Exec().ReadWrite(xy.X, xy.Y)
	return output[0] == 1
}

func nextScan(prog *intcode.Program, line scan) scan {
	next := scan{y: line.y + 1}
	c := XY{X: line.x0, Y: next.y}
	for !check(prog, c) {
		c.X++
	}
	next.x0 = c.X
	c.X = line.x1
	for check(prog, c) {
		c.X++
	}
	next.x1 = c.X
	return next
}
