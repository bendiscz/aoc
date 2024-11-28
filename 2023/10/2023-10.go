package main

import (
	"fmt"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2023
	day     = 10
	example = `

FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L

`
)

//..........
//.S------7.
//.|F----7|.
//.||OOOO||.
//.||OOOO||.
//.|L-7F-J|.
//.|II||II|.
//.L--JL--J.
//..........

//.F----7F7F7F7F-7....
//.|F--7||||||||FJ....
//.||.FJ||||||||L7....
//FJL7L7LJLJ||LJ.L-7..
//L--J.L7...LJS7F-7L7.
//....F-J..F7FJ|L7L7L7
//....L7.F7||L7|.L7L7|
//.....|FJLJ|FJ|F7|.LJ
//....FJL-7.||.||||...
//....L---J.LJ.LJLJ...

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	val  byte
	step int
}

func (c cell) String() string { return fmt.Sprintf("%c", c.val) }

type grid = *Matrix[cell]

func test(d1, d2, t1, t2 XY) bool {
	return d1 == t1 && d2 == t2 || d1 == t2 && d2 == t1
}

func solve(p *Problem) {
	g := grid(nil)
	for p.NextLine() {
		if g == nil {
			g = NewMatrix[cell](Rectangle(len(p.Line()), 0))
		}
		ParseVectorFunc(g.AppendRow(), p.Line(), func(b byte) cell { return cell{val: b} })
	}

	start := findStart(g)
	path := findLoop(g, start)
	p.PartOne(len(path) / 2)

	d1, d2 := path[1].Sub(start), path[len(path)-2].Sub(start)
	switch {
	case test(d1, d2, PosX, NegX):
		g.At(start).val = '-'
	case test(d1, d2, PosY, NegY):
		g.At(start).val = '|'
	case test(d1, d2, NegY, PosX):
		g.At(start).val = 'L'
	case test(d1, d2, NegY, NegX):
		g.At(start).val = 'J'
	case test(d1, d2, PosY, PosX):
		g.At(start).val = '7'
	case test(d1, d2, NegY, PosX):
		g.At(start).val = 'F'
	}

	border := map[XY]bool{}
	for i, xy := range path {
		g.At(xy).step = i + 1
		border[xy] = true
	}
	count := 0
	for y := 0; y < g.Dim.Y; y++ {
		inside := false
		for x := 0; x < g.Dim.X; x++ {
			xy := XY{x, y}
			if border[xy] {
				b := g.At(xy).val
				if b == '|' || b == 'L' || b == 'J' {
					inside = !inside
				}
			} else if inside {
				count++
			}
		}
	}
	p.PartTwo(count)
}

func findLoop(g grid, start XY) []XY {
	for _, dir := range HVDirs {
		if path, ok := checkLoop(g, start, dir); ok {
			return path
		}
	}
	return nil
}

func findStart(g grid) XY {
	for y := 0; y < g.Dim.Y; y++ {
		for x := 0; x < g.Dim.X; x++ {
			if g.AtXY(x, y).val == 'S' {
				return XY{x, y}
			}
		}
	}
	panic("no start found")
}

func move(g grid, xy, dir XY) (XY, XY, bool) {
	xy2 := xy.Add(dir)
	if !g.Dim.HasInside(xy2) {
		return XY{}, XY{}, false
	}

	ch := g.At(xy2).val
	if ch == 'S' {
		return xy2, dir, true
	}

	switch dir {
	case PosX:
		switch ch {
		case '-':
			dir = PosX
		case '7':
			dir = PosY
		case 'J':
			dir = NegY
		default:
			return XY{}, XY{}, false
		}
	case NegX:
		switch ch {
		case '-':
			dir = NegX
		case 'F':
			dir = PosY
		case 'L':
			dir = NegY
		default:
			return XY{}, XY{}, false
		}
	case PosY:
		switch ch {
		case '|':
			dir = PosY
		case 'L':
			dir = PosX
		case 'J':
			dir = NegX
		default:
			return XY{}, XY{}, false
		}
	case NegY:
		switch ch {
		case '|':
			dir = NegY
		case 'F':
			dir = PosX
		case '7':
			dir = NegX
		default:
			return XY{}, XY{}, false
		}
	}
	return xy2, dir, true
}

func checkLoop(g grid, start, dir XY) ([]XY, bool) {
	xy, path := start, []XY{start}
	for {
		var ok bool
		xy, dir, ok = move(g, xy, dir)
		if !ok {
			return nil, false
		}

		path = append(path, xy)

		if g.At(xy).val == 'S' {
			return path, true
		}
	}
}
