package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2021
	day     = 25
	example = `

v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>

`
)

func main() {
	Run(year, day, example, solve)
}

type sea struct {
	*Matrix[byte]
}

func (s sea) move() (moved bool) {
	moved = s.moveEast() || moved
	moved = s.moveSouth() || moved
	return
}

func (s sea) moveEast() (moved bool) {
	for y := 0; y < s.Dim.Y; y++ {
		row := s.Data[y]
		x, free := 1, row[0] == '.'
		for x < s.Dim.X {
			if row[x] == '.' && row[x-1] == '>' {
				moved = true
				row[x-1], row[x] = '.', '>'
				x++
			}
			x++
		}
		if x == s.Dim.X && free && row[x-1] == '>' {
			moved = true
			row[x-1], row[0] = '.', '>'
		}
	}
	return
}

func (s sea) moveSouth() (moved bool) {
	for x := 0; x < s.Dim.X; x++ {
		y, free := 1, s.Data[0][x] == '.'
		for y < s.Dim.Y {
			if s.Data[y][x] == '.' && s.Data[y-1][x] == 'v' {
				moved = true
				s.Data[y-1][x], s.Data[y][x] = '.', 'v'
				y++
			}
			y++
		}
		if y == s.Dim.Y && free && s.Data[y-1][x] == 'v' {
			moved = true
			s.Data[y-1][x], s.Data[0][x] = '.', 'v'
		}
	}
	return
}

func solve(p *Problem) {
	s := sea{}
	for p.NextLine() {
		if s.Matrix == nil {
			s.Matrix = NewMatrix[byte](Rectangle(len(p.Line()), 0))
		}
		CopyVector(s.AppendRow(), p.Line())
	}

	count := 1
	for s.move() {
		count++
	}

	p.PartOne(count)
	p.PartTwo("Merry Christmas!")
}
