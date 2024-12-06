package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"strings"
)

const (
	year    = 2024
	day     = 6
	example = `

....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	ch byte
	d  uint
	v  int
}

var dirs = [...]XY{NegY, PosX, PosY, NegX}

type grid struct {
	*Matrix[cell]
	iter int
}

func solve(p *Problem) {
	g := &grid{NewMatrix[cell](Rectangle(len(p.PeekLine()), 0)), 1}
	start := XY{X: -1, Y: -1}
	for p.NextLine() {
		if start.X == -1 {
			if i := strings.IndexByte(p.Line(), '^'); i >= 0 {
				start = XY{X: i, Y: g.Size().Y}
			}
		}
		ParseVectorFunc(g.AppendRow(), p.Line(), func(b byte) cell { return cell{ch: b} })
	}
	g.At(start).ch = '.'

	path := make([]XY, 0, g.Dim.X*g.Dim.Y)
	trace(g, start, 0, &path)
	p.PartOne(len(path))

	s := 0
	for _, pos := range path {
		if pos == start {
			continue
		}

		g.At(pos).ch = '#'
		if trace(g, start, 0, nil) {
			s++
		}
		g.At(pos).ch = '.'
	}
	p.PartTwo(s)
}

func trace(g *grid, pos XY, dir int, path *[]XY) bool {
	g.iter++

	g.At(pos).v = g.iter
	g.At(pos).d = 1 << dir
	if path != nil {
		*path = append(*path, pos)
	}

	for {
		n := pos.Add(dirs[dir])
		if !g.Dim.HasInside(n) {
			return false
		}
		if g.At(n).ch == '#' {
			dir = (dir + 1) % 4
			continue
		}

		pos = n
		if g.At(pos).v == g.iter {
			if g.At(pos).d&(1<<dir) != 0 {
				return true
			}
			g.At(pos).d |= 1 << dir
		} else {
			g.At(pos).v = g.iter
			g.At(pos).d = 1 << dir
			if path != nil {
				*path = append(*path, pos)
			}
		}
	}
}
