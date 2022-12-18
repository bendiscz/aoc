package main

import (
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year = 2021
	day  = 24
)

func main() {
	Run(year, day, "", solve)
}

type block struct {
	index    int
	sub, add int
}

type link struct {
	i1, i2 int
	diff   int
}

func (l link) max() (int, int) {
	if l.diff > 0 {
		return 9 - l.diff, 9
	} else {
		return 9, 9 + l.diff
	}
}

func (l link) min() (int, int) {
	if l.diff > 0 {
		return 1, 1 + l.diff
	} else {
		return 1 - l.diff, 1
	}
}

func digit(x int) byte {
	return byte('0' + x)
}

func readBlock(p *Problem, index int) block {
	b := block{index: index}
	for i := 0; i < 18; i++ {
		p.NextLine()
		switch i {
		case 5:
			b.sub = ParseInt(strings.Fields(p.Line())[2])
		case 15:
			b.add = ParseInt(strings.Fields(p.Line())[2])
		}
	}
	return b
}

func solve(p *Problem) {
	var stack []block
	var links []link
	for i := 0; i < 14; i++ {
		b := readBlock(p, i)
		if b.sub < 0 {
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			links = append(links, link{p.index, b.index, p.add + b.sub})
		} else {
			stack = append(stack, b)
		}
	}

	var max, min [14]byte
	for _, l := range links {
		max1, max2 := l.max()
		max[l.i1], max[l.i2] = digit(max1), digit(max2)
		min1, min2 := l.min()
		min[l.i1], min[l.i2] = digit(min1), digit(min2)
	}

	p.PartOne(string(max[:]))
	p.PartTwo(string(min[:]))
}
