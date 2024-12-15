package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"strings"
)

const (
	year    = 2024
	day     = 15
	example = `

##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^

`
)

//#######
//#...#.#
//#.....#
//#..OO@#
//#..O..#
//#.....#
//#######
//
//<vv<<^^<<^^

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

var dir = []XY{
	'>': PosX,
	'<': NegX,
	'v': PosY,
	'^': NegY,
}

func solve(p *Problem) {
	g1 := grid{NewMatrix[cell](Rectangle(len(p.PeekLine()), 0))}
	g2 := grid{NewMatrix[cell](Rectangle(len(p.PeekLine())*2, 0))}
	xy1, xy2 := XY0, XY0
	for p.NextLine() && p.Line() != "" {
		if x := strings.IndexByte(p.Line(), '@'); x != -1 {
			xy1 = XY{X: x, Y: g1.Dim.Y}
			xy2 = XY{X: 2 * x, Y: g1.Dim.Y}
		}

		s := p.Line()
		l2 := strings.Builder{}
		for i := 0; i < len(s); i++ {
			switch s[i] {
			case '#':
				l2.WriteString("##")
			case 'O':
				l2.WriteString("[]")
			case '@':
				l2.WriteString("@.")
			default:
				l2.WriteString("..")
			}
		}
		ParseVectorFunc(g1.AppendRow(), s, func(b byte) cell { return cell{ch: b} })
		ParseVectorFunc(g2.AppendRow(), l2.String(), func(b byte) cell { return cell{ch: b} })
	}

	for p.NextLine() {
		s := p.Line()
		for i := 0; i < len(s); i++ {
			d := dir[s[i]]
			xy1 = pushOne(g1, xy1, d)
			xy2 = pushTwo(g2, xy2, d)
		}
	}

	p.PartOne(score(g1))
	p.PartTwo(score(g2))
}

func score(g grid) int {
	s := 0
	for xy, c := range g.All() {
		if c.ch == 'O' || c.ch == '[' {
			s += xy.X + xy.Y*100
		}
	}
	return s
}

func pushOne(g grid, xy, d XY) XY {
	c := xy.Add(d)
	var ch byte
	for {
		if !g.Dim.HasInside(c) {
			ch = '#'
			break
		}

		ch = g.At(c).ch
		if ch != '[' && ch != ']' && ch != 'O' {
			break
		}
		c = c.Add(d)
	}

	if ch == '#' {
		return xy
	}

	for {
		c2 := c.Sub(d)
		g.At(c).ch = g.At(c2).ch
		if c2 == xy {
			break
		}
		c = c2
	}

	g.At(xy).ch = '.'
	xy = xy.Add(d)
	return xy
}

func pushTwo(g grid, xy, d XY) XY {
	if d.Y == 0 {
		return pushOne(g, xy, d)
	}

	c := xy.Add(d)
	if !g.Dim.HasInside(c) || g.At(c).ch == '#' {
		return xy
	}

	boxes, ok := []XY(nil), true
	switch {
	case g.At(c).ch == '[':
		boxes, ok = collect(g, c, d)
	case g.At(c).ch == ']':
		boxes, ok = collect(g, c.Add(NegX), d)
	}
	if !ok {
		return xy
	}

	for i := len(boxes) - 1; i >= 0; i-- {
		c1 := boxes[i]
		c2 := c1.Add(PosX)
		g.At(c1.Add(d)).ch = '['
		g.At(c2.Add(d)).ch = ']'
		g.At(c1).ch = '.'
		g.At(c2).ch = '.'
	}

	g.At(xy).ch = '.'
	g.At(c).ch = '@'
	return c
}

func collect(g grid, b, d XY) ([]XY, bool) {
	q := Queue[XY]{}
	q.Push(b)

	boxes := []XY(nil)
	for q.Len() > 0 {
		xy := q.Pop()

		c1 := xy.Add(d)
		if !g.Dim.HasInside(c1) {
			return nil, false
		}
		c2 := c1.Add(PosX)

		ch1, ch2 := g.At(c1).ch, g.At(c2).ch

		if ch1 == '#' || ch2 == '#' {
			return nil, false
		}

		boxes = append(boxes, xy)

		if ch1 == '[' {
			q.Push(c1)
		}
		if ch1 == ']' {
			q.Push(c1.Add(NegX))
		}
		if ch2 == '[' {
			q.Push(c2)
		}
	}

	return boxes, true
}
