package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2021
	day     = 11
	example = `

5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526

`
)

func main() {
	Run(year, day, example, solve)
}

type bottom [10][10]int

func (b *bottom) glow() {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			b[y][x]++
		}
	}
}

func (b *bottom) iterate() bool {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if b[y][x] > 9 {
				b.flash(x, y)
				return true
			}
		}
	}
	return false
}

func (b *bottom) flash(x0, y0 int) {
	b[y0][x0] = -1
	for x := x0 - 1; x <= x0+1; x++ {
		for y := y0 - 1; y <= y0+1; y++ {
			if x < 0 || x >= 10 || y < 0 || y >= 10 {
				continue
			}
			if b[y][x] >= 0 {
				b[y][x]++
			}
		}
	}
}

func (b *bottom) countFlashes() int {
	count := 0
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if b[y][x] == -1 {
				count++
				b[y][x] = 0
			}
		}
	}
	return count
}

func solve(p *Problem) {
	b := &bottom{}
	for y := 0; p.NextLine(); y++ {
		ParseVectorDigits[int](Slice[int](b[y][:]), p.Line())
	}

	count := 0
	all := -1
	for i := 0; ; i++ {
		b.glow()
		for b.iterate() {
		}
		x := b.countFlashes()
		if i < 100 {
			count += x
		}
		if x == 100 {
			all = i + 1
			break
		}
	}

	p.PartOne(count)
	p.PartTwo(all)
}
