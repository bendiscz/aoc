package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"math"
)

const (
	year    = 2020
	day     = 13
	example = `

939
7,13,x,x,59,x,31,19

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	ch byte
}

type grid struct {
	*Matrix[cell]
}

type bus struct {
	id  int
	off int
}

func solve(p *Problem) {
	p.NextLine()
	ts := ParseInt(p.Line())
	p.NextLine()
	bs := []bus(nil)
	for i, f := range SplitFields(p.Line()) {
		if f == "x" {
			continue
		}
		bs = append(bs, bus{ParseInt(f), i})
	}

	minD := math.MaxInt
	minB := 0
	maxB := bus{}
	for _, b := range bs {
		d := b.id - ts%b.id
		if d < minD {
			minD = d
			minB = b.id
		}
		if b.id > maxB.id {
			maxB = b
		}
	}
	p.PartOne(minB * minD)

	b0 := bus{id: 1, off: 0}
	for _, b := range bs {
		b0 = merge(b0, b)
	}
	p.PartTwo(b0.off)
}

func merge(b1, b2 bus) bus {
	if b1.id < b2.id {
		b1, b2 = b2, b1
	}

	for x := b1.off; ; x += b1.id {
		if (x+b2.off)%b2.id == 0 {
			return bus{id: LCM(b1.id, b2.id), off: x}
		}
	}
}
