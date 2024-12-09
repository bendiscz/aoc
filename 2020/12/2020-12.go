package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2020
	day     = 12
	example = `

F10
N3
F7
R90
F11

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	pos1, pos2, dir1, dir2 := XY{0, 0}, XY{0, 0}, PosX, XY{10, -1}
	for p.NextLine() {
		var ch byte
		var n int
		p.Scanf("%c%d", &ch, &n)

		switch ch {
		case 'F':
			pos1 = move(pos1, dir1, n)
			pos2 = move(pos2, dir2, n)
		case 'N':
			pos1 = move(pos1, NegY, n)
			dir2 = move(dir2, NegY, n)
		case 'S':
			pos1 = move(pos1, PosY, n)
			dir2 = move(dir2, PosY, n)
		case 'W':
			pos1 = move(pos1, NegX, n)
			dir2 = move(dir2, NegX, n)
		case 'E':
			pos1 = move(pos1, PosX, n)
			dir2 = move(dir2, PosX, n)
		case 'R':
			dir1 = rotate(dir1, n)
			dir2 = rotate(dir2, n)
		case 'L':
			dir1 = rotate(dir1, 360-n)
			dir2 = rotate(dir2, 360-n)
		}
	}

	p.PartOne(Abs(pos1.X) + Abs(pos1.Y))
	p.PartTwo(Abs(pos2.X) + Abs(pos2.Y))
}

func move(xy, dir XY, n int) XY {
	return xy.Add(XY{dir.X * n, dir.Y * n})
}

func rotate(dir XY, d int) XY {
	for d > 0 {
		dir = XY{-dir.Y, dir.X}
		d -= 90
	}
	return dir
}
