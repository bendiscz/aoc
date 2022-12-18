package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2021
	day     = 20
	example = `

..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..##
#..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###
.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#.
.#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#.....
.#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#..
...####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.....
..##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###

`
)

func main() {
	Run(year, day, example, solve)
}

const margin = 50

type image struct {
	*Matrix[bool]
}

func (img image) get(x0, y0 int) int {
	v := 0
	for y := y0 - 1; y <= y0+1; y++ {
		for x := x0 - 1; x <= x0+1; x++ {
			var p bool
			if x < 0 || x >= img.Dim.X || y < 0 || y >= img.Dim.Y {
				p = img.Data[0][0]
			} else {
				p = img.Data[y][x]
			}

			v *= 2
			if p {
				v++
			}
		}
	}
	return v
}

func (img image) count() int {
	count := 0
	for x := 0; x < img.Dim.X; x++ {
		for y := 0; y < img.Dim.Y; y++ {
			if img.Data[y][x] {
				count++
			}
		}
	}
	return count
}

func (img image) enhance(key string) image {
	img2 := image{NewMatrix[bool](img.Dim)}
	for x := 0; x < img.Dim.X; x++ {
		for y := 0; y < img.Dim.Y; y++ {
			img2.Data[y][x] = key[img.get(x, y)] == '#'
		}
	}
	return img2
}

func solve(p *Problem) {
	key := ""
	for p.NextLine() {
		if p.Line() == "" {
			break
		}
		key += p.Line()
	}

	var img image
	for p.NextLine() {
		if img.Matrix == nil {
			img.Matrix = NewMatrix[bool](Rectangle(len(p.Line())+2*margin, margin))
		}
		row := img.AppendRow()
		for x := 0; x < len(p.Line()); x++ {
			*row.At(x + margin) = p.Line()[x] == '#'
		}
	}

	for i := 0; i < margin; i++ {
		img.AppendRow()
	}

	img = img.enhance(key)
	img = img.enhance(key)
	p.PartOne(img.count())

	for i := 0; i < 48; i++ {
		img = img.enhance(key)
	}
	p.PartTwo(img.count())
}
