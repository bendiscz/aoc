package main

import (
	"fmt"
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2024
	day     = 20
	example = `

###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	ch byte
	l  int
	l2 int
}

func (c cell) String() string {
	if c.l > 0 {
		return "O"
	}
	return string([]byte{c.ch})
}

type grid struct {
	*Matrix[cell]
}

type vertex struct {
	g  grid
	at XY
}

func solve(p *Problem) {
	s1, s2 := 0, 0

	g := grid{NewMatrix[cell](Rectangle(len(p.PeekLine()), 0))}
	start, end := XY0, XY0
	for p.NextLine() {
		s := p.Line()
		for x := 0; x < g.Dim.X; x++ {
			if s[x] == 'S' {
				start = XY{X: x, Y: g.Dim.Y}
			}
			if s[x] == 'E' {
				end = XY{X: x, Y: g.Dim.Y}
			}
		}
		ParseVectorFunc(g.AppendRow(), s, func(b byte) cell { return cell{ch: b} })
	}

	//p.Printf("%v %v", start, end)
	//PrintGrid(g)

	//length := findPath(g, start, end)
	//p.Printf("%d", findPath(g, start, end))

	//save := 100
	//for y := 1; y < g.Dim.Y-1; y++ {
	//	for x := 1; x < g.Dim.X-2; x++ {
	//		//p.Printf("trying %d,%d", x, y)
	//		c1, c2 := g.AtXY(x, y), g.AtXY(x, y)
	//		if c1.ch == '#' && c2.ch == '#' {
	//			continue
	//		}
	//		ch1, ch2 := c1.ch, c2.ch
	//		c1.ch, c2.ch = '.', '.'
	//
	//		//for _, c := range g.All() {
	//		//	c.l = 0
	//		//}
	//		//PrintGrid(g)
	//
	//		if l := findPath(g, start, end); l <= length-save && c1.l > 0 && c2.l > 0 /* && Abs(c2.l-c1.l) == 1*/ {
	//			fmt.Printf("!!!!!!!!!!!!!!!! X cheat found %d %d\n", x, y)
	//			PrintGrid(g)
	//			s1++
	//		}
	//		c1.ch, c2.ch = ch1, ch2
	//		//fmt.Println("------------------------")
	//	}
	//}

	//for y := 1; y < g.Dim.Y-2; y++ {
	//	for x := 1; x < g.Dim.X-1; x++ {
	//		//p.Printf("trying %d,%d", x, y)
	//		c1, c2 := g.AtXY(x, y), g.AtXY(x, y+1)
	//		if c1.ch == '#' && c2.ch == '#' {
	//			continue
	//		}
	//		ch1, ch2 := c1.ch, c2.ch
	//		c1.ch, c2.ch = '.', '.'
	//		//PrintGrid(g)
	//		if l := findPath(g, start, end); l <= length-save && c1.l > 0 && c2.l > 0/* && Abs(c2.l-c1.l) == 1*/ {
	//			fmt.Printf("!!!!!!!!!!!!!!!! Y cheat found %d %d  %d\n", x, y, l)
	//			PrintGrid(g)
	//			s1++
	//		}
	//		c1.ch, c2.ch = ch1, ch2
	//	}
	//}

	//for l := 1; l <= length; l++ {
	//	for _, d := range HVDirs {
	//		l2 := findPathCheat(g, start, end, l, d)
	//		if l2 <= length-save {
	//			s1++
	//		}
	//	}
	//}
	p.PartOne(s1)

	l1 := findPath(g, start, end)
	l2 := findPathL2(g, start, end)
	p.Printf("%d %d", l1, l2)
	p.Printf("%d %d", g.AtXY(7, 7).l, g.AtXY(5, 7).l2)

	s2 = solvePart2(g, l1, 100, 20)
	p.PartTwo(s2)

	//os.Exit(1)
}

func solvePart2(g grid, length, save, cheat int) int {
	s := 0
	for xy, c := range g.All() {
		if c.l == 0 {
			continue
		}
		for x := xy.X - cheat; x <= xy.X+cheat; x++ {
			for y := xy.Y - cheat; y <= xy.Y+cheat; y++ {
				if xy == (XY{7, 7}) && x == 5 && y == 7 {
					//fmt.Println()
				}

				if c.l == 0 || !g.Dim.HasInside(XY{x, y}) || g.AtXY(x, y).l2 == 0 {
					continue
				}

				d := Abs(x-xy.X) + Abs(y-xy.Y)
				if d < 2 || d > cheat {
					continue
				}

				l := g.AtXY(x, y).l2 + c.l + d - 2
				if l < length && l+save <= length {
					fmt.Printf("cheat %v - %v\n", xy, XY{x, y})
					s++
				}
			}
		}
	}
	return s
}

func findPath(g grid, start, end XY) int {
	for _, c := range g.All() {
		c.l = 0
	}
	q := Queue[XY]{}
	q.Push(start)
	g.At(start).l = 1
	for q.Len() > 0 {
		xy := q.Pop()
		c := g.At(xy)
		for _, d := range HVDirs {
			xy2 := xy.Add(d)
			if xy2 == end {
				g.At(xy2).l = c.l + 1
				return c.l
			}

			c2 := g.At(xy2)
			if c2.ch != '#' && c2.l == 0 {
				c2.l = c.l + 1
				q.Push(xy2)
			}
		}
	}
	panic("no path")
}

func findPathL2(g grid, start, end XY) int {
	for _, c := range g.All() {
		c.l2 = 0
	}
	q := Queue[XY]{}
	q.Push(end)
	g.At(end).l2 = 1
	for q.Len() > 0 {
		xy := q.Pop()
		c := g.At(xy)
		for _, d := range HVDirs {
			xy2 := xy.Add(d)
			if xy2 == start {
				g.At(xy2).l2 = c.l2 + 1
				return c.l2 + 1
			}

			c2 := g.At(xy2)
			if c2.ch != '#' && c2.l2 == 0 {
				c2.l2 = c.l2 + 1
				q.Push(xy2)
			}
		}
	}
	panic("no path")
}

func findPathCheat(g grid, start, end XY, cheat int, cheatDir XY) int {
	for _, c := range g.All() {
		c.l = 0
	}
	q := Queue[XY]{}
	q.Push(start)
	g.At(start).l = 1
	for q.Len() > 0 {
		xy := q.Pop()
		c := g.At(xy)

		cheating := c.l == cheat

		for _, d := range HVDirs {
			xy2 := xy.Add(d)
			if !g.Dim.HasInside(xy2) {
				continue
			}

			if xy2 == end {
				g.At(xy2).l = c.l + 1
				return c.l
			}

			c2 := g.At(xy2)
			if (c2.ch != '#' || cheating && cheatDir == d) && c2.l == 0 {
				c2.l = c.l + 1
				q.Push(xy2)
			}
		}
	}
	panic("no path")
}
