package main

import (
	"fmt"
	"unicode"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2022
	day     = 22
	example = `

        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5

`
)

// .##
// .#.
// ##.
// #..

// ..#.
// ###.
// ..##

// .#.
// ###
// .#.
// .#.

// ####
// ####
// ####

func main() {
	//Run(year, day, example, solve)
	Run(year, day, "", solve)
}

type dir int

const (
	right dir = iota
	down
	left
	up
)

func (d dir) rotateL() dir {
	return (d + 3) % 4
}

func (d dir) rotateR() dir {
	return (d + 1) % 4
}

type coord int

const (
	posX coord = 1
	negX coord = -1
	posY coord = 2
	negY coord = -2
	posZ coord = 3
	negZ coord = -3
)

type face struct {
	side  coord
	right coord
	down  coord
}

func (f face) move(d dir) face {
	switch d {
	case right:
		return face{side: f.right, right: -f.side, down: f.down}
	case down:
		return face{side: f.down, right: f.right, down: -f.side}
	case left:
		return face{side: -f.right, right: f.side, down: f.down}
	case up:
		return face{side: -f.down, right: f.right, down: f.side}
	default:
		panic(fmt.Sprintf("invalid dir: %v", d))
	}
}

type tile byte

const (
	void tile = iota
	space
	wall
)

type cell struct {
	tile tile
}

type surface struct {
	*Matrix[cell]
	size int
}

type board struct {
	*Matrix[cell]
}

var dirs = [...]XY{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func (b board) findStart() XY {
	for x := 0; x < b.Dim.X; x++ {
		if b.AtXY(x, 0).tile == space {
			return XY{x, 0}
		}
	}
	panic("no start tile found")
}

func (b board) move(c XY, d, n int) XY {
	dir := dirs[d]
	c0 := c
	for n > 0 {
		c = c.Add(dir)
		c = XY{(c.X + b.Dim.X) % b.Dim.X, (c.Y + b.Dim.Y) % b.Dim.Y}
		t := b.At(c).tile
		if t == void {
			continue
		}
		if t == wall {
			c = c0
			break
		}
		c0 = c
		n--
	}
	return c
}

func (b board) moveCube(c XY, d dir, n int) (XY, dir) {
	c0, d0 := c, d
	for n > 0 {
		c = c.Add(dirs[d])
		//c = XY{(c.X + b.Dim.X) % b.Dim.X, (c.Y + b.Dim.Y) % b.Dim.Y}
		if !b.Dim.HasInside(c) || b.At(c).tile == void {
			if d == 1 {
				if c.X < 50 {
					c.X = 100 + c.X
					c.Y = 0
				} else if c.X < 100 {
					c.Y = 150 + c.X - 50
					c.X = 49
					d = d.rotateR()
				} else {
					c.Y = 50 + c.X - 100
					c.X = 99
					d = d.rotateR()
				}
			} else if d == 0 {
				if c.Y < 50 {
					c.Y = 100 + 49 - c.Y
					c.X = 99
					d = d.rotateR().rotateR()
				} else if c.Y < 100 {
					c.X = 100 + c.Y - 50
					c.Y = 49
					d = d.rotateL()
				} else if c.Y < 150 {
					c.Y = 0 + 149 - c.Y
					c.X = 149
					d = d.rotateR().rotateR()
				} else {
					c.X = 50 + c.Y - 150
					c.Y = 149
					d = d.rotateL()
				}
			} else if d == 3 {
				if c.X < 50 {
					c.Y = 50 + c.X
					c.X = 50
					d = d.rotateR()

				} else if c.X < 100 {
					c.Y = 150 + c.X - 50
					c.X = 0
					d = d.rotateR()
				} else {
					c.X -= 2 * 50
					c.Y = 199
				}
			} else if d == 2 {
				if c.Y < 50 {
					c.Y = 100 + 49 - c.Y
					c.X = 0
					d = d.rotateR().rotateR()
				} else if c.Y < 100 {
					c.X = 0 + c.Y - 50
					c.Y = 100
					d = d.rotateL()
				} else if c.Y < 150 {
					c.Y = 0 + 149 - c.Y
					c.X = 50
					d = d.rotateR().rotateR()
				} else {
					c.X = 50 + c.Y - 150
					c.Y = 0
					d = d.rotateL()
				}
			}
		}

		if b.At(c).tile == wall {
			c, d = c0, d0
			break
		}
		c0, d0 = c, d
		n--
	}
	return c, d
}

func solve(p *Problem) {
	b := board{NewMatrix[cell](Rectangle(150, 0))}
	for p.NextLine() && p.Line() != "" {
		row := b.AppendRow()
		ParseVectorMap(row, p.Line(), map[byte]cell{
			' ': cell{void},
			'.': cell{space},
			'#': cell{wall},
		})
	}

	p.NextLine()
	line, i, d := p.Line(), 0, dir(0)
	c := b.findStart()
	for i < len(line) {
		s := i
		for i < len(line) && unicode.IsDigit(rune(line[i])) {
			i++
		}
		n := ParseInt(line[s:i])

		//p.Printf("moving from: %v, dir: %v, by: %d", c, dirs[d], n)
		c = b.move(c, int(d), n)
		//p.Printf("c: %v", c)

		if i == len(line) {
			break
		}

		switch line[i] {
		case 'L':
			d = d.rotateL()
		case 'R':
			d = d.rotateR()
		}
		i++
	}

	pw := 1000*(c.Y+1) + 4*(c.X+1) + int(d)
	p.PartOne(pw)

	i, d = 0, dir(0)
	c = b.findStart()
	for i < len(line) {
		s := i
		for i < len(line) && unicode.IsDigit(rune(line[i])) {
			i++
		}
		n := ParseInt(line[s:i])

		//p.Printf("moving from: %v, dir: %v, by: %d", c, dirs[d], n)
		c, d = b.moveCube(c, d, n)
		//p.Printf("c: %v", c)

		if i == len(line) {
			break
		}

		switch line[i] {
		case 'L':
			d = d.rotateL()
		case 'R':
			d = d.rotateR()
		}
		i++
	}

	pw = 1000*(c.Y+1) + 4*(c.X+1) + int(d)
	p.PartTwo(pw)
}
