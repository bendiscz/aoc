package main

import (
	"fmt"
	. "github.com/bendiscz/aoc/aoc"
	"math"
	"strings"
)

const (
	year    = 2020
	day     = 20
	example = `

Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...

`
)

func main() {
	Run(year, day, example, solve)
}

const N = 10
const E = uint16(math.MaxUint16)

type tile struct {
	id int
	e  [8]edges
}

type edges struct {
	t, r, b, l uint16
	g          Grid[cell]
}

func (e edges) String() string {
	str := strings.Builder{}
	str.WriteByte('\n')
	str.WriteByte(' ')
	for x := 0; x < N; x++ {
		str.WriteByte(bitChar(e.t, x))
	}
	str.WriteByte('\n')

	for y := 0; y < 10; y++ {
		str.WriteByte(bitChar(e.l, y))
		for x := 0; x < 10; x++ {
			str.WriteByte(' ')
		}
		str.WriteByte(bitChar(e.r, y))
		str.WriteByte('\n')
	}

	str.WriteByte(' ')
	for x := 0; x < N; x++ {
		ch := byte('.')
		if e.b&(1<<x) != 0 {
			ch = '#'
		}
		str.WriteByte(ch)
	}
	str.WriteByte('\n')

	return str.String()
}

func bitChar(x uint16, shift int) byte {
	if x&(1<<shift) == 0 {
		return '.'
	} else {
		return '#'
	}
}

func (e edges) rotate() edges {
	return edges{flip(e.l), e.t, flip(e.r), e.b, e.g.RotateR()}
}

func (e edges) flip() edges {
	return edges{flip(e.t), e.l, flip(e.b), e.r, e.g.FlipX()}
}

func flip(x uint16) (y uint16) {
	for i := 0; i < N; i++ {
		y |= ((x >> i) & 1) << (N - 1 - i)
	}
	return
}

type tileCell struct {
	t tile
	e edges
}

type tileGrid struct {
	*Matrix[tileCell]
	tiles []tile
	used  map[int]bool
}

type cell struct {
	b, m bool
}

func (c cell) String() string {
	switch {
	case c.m:
		return "O"
	case c.b:
		return "#"
	default:
		return "."
	}
}

type grid struct {
	*Matrix[cell]
}

func solve(p *Problem) {
	s1, s2 := 0, 0

	tiles := []tile(nil)
	for p.NextLine() {
		t, e := tile{}, edges{}
		e.g = grid{NewMatrix[cell](Square(N - 2))}
		p.Scanf("Tile %d:", &t.id)

		for y := 0; y < N; y++ {
			p.NextLine()
			s := p.Line()

			if y == 0 || y == N-1 {
				tb := &e.t
				if y != 0 {
					tb = &e.b
				}
				for x := 0; x < N; x++ {
					if s[x] == '#' {
						*tb |= 1 << x
					}
				}
			} else {
				for x := 1; x < N-1; x++ {
					e.g.AtXY(x-1, y-1).b = s[x] == '#'
				}
			}

			if s[0] == '#' {
				e.l |= 1 << y
			}
			if s[N-1] == '#' {
				e.r |= 1 << y
			}
		}

		t.e[0], e = e, e.rotate()
		t.e[1], e = e, e.rotate()
		t.e[2], e = e, e.rotate()
		t.e[3], e = e, e.rotate().flip()
		t.e[4], e = e, e.rotate()
		t.e[5], e = e, e.rotate()
		t.e[6], e = e, e.rotate()
		t.e[7] = e

		tiles = append(tiles, t)

		p.NextLine()
	}

	n := int(math.Sqrt(float64(len(tiles))))
	if n*n != len(tiles) {
		panic("not a square")
	}

	tg := &tileGrid{NewMatrix[tileCell](Square(n)), tiles, map[int]bool{}}
	s1, _ = fill(tg, XY{0, 0})
	p.PartOne(s1)

	g := grid{NewMatrix[cell](Square(n * (N - 2)))}
	for xy := range tg.Dim.All() {
		CopyGrid(g.SubGrid(XY{xy.X * (N - 2), xy.Y * (N - 2)}, XY{N - 2, N - 2}), tg.At(xy).e.g)
	}

	findMonsters(g)
	findMonsters(g.RotateR())
	findMonsters(g.RotateR().RotateR())
	findMonsters(g.RotateR().RotateR().RotateR())
	findMonsters(g.Trans())
	findMonsters(g.Trans().RotateR())
	findMonsters(g.Trans().RotateR().RotateR())
	findMonsters(g.Trans().RotateR().RotateR().RotateR())

	PrintGrid[cell](g)
	fmt.Println()

	for _, c := range g.All() {
		if c.b && !c.m {
			s2++
		}
	}
	p.PartTwo(s2)
}

func fill(g *tileGrid, xy XY) (int, bool) {
	if xy.Y == g.Dim.Y {
		clear(g.used)
		return g.AtXY(0, 0).t.id * g.AtXY(g.Dim.X-1, 0).t.id * g.AtXY(0, g.Dim.Y-1).t.id * g.AtXY(g.Dim.X-1, g.Dim.Y-1).t.id, true
	}

	nxy := XY{xy.X + 1, xy.Y}
	if nxy.X == g.Dim.X {
		nxy = XY{0, xy.Y + 1}
	}

	h, v := E, E
	if xy.X > 0 {
		h = g.AtXY(xy.X-1, xy.Y).e.r
	}
	if xy.Y > 0 {
		v = g.AtXY(xy.X, xy.Y-1).e.b
	}

	for i, t := range g.tiles {
		if g.used[i] {
			continue
		}

		g.used[i] = true
		g.At(xy).t = t

		for _, e := range t.e {
			if h != E && h != e.l || v != E && v != e.t {
				continue
			}
			g.At(xy).e = e
			if r, ok := fill(g, nxy); ok {
				return r, true
			}
		}

		delete(g.used, i)
	}

	return 0, false
}

func findMonsters(g Grid[cell]) {
	monster := []string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}
	ms := Rectangle(len(monster[0]), len(monster))

	for o := range g.Size().Sub(ms).All() {
		s := g.SubGrid(o, ms)
		found := true
		for xy, c := range s.All() {
			if monster[xy.Y][xy.X] == '#' && !c.b {
				found = false
				break
			}
		}
		if found {
			for xy, c := range s.All() {
				if monster[xy.Y][xy.X] == '#' {
					c.m = true
				}
			}
		}
	}
}
