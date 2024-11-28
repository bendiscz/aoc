package main

import (
	"fmt"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2022
	day     = 17
	example = `

>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>

`
)

func main() {
	Run(year, day, example, solve)
}

type rock struct {
	width int
	shape [][4]bool
}

const (
	WC = 7
	WR = 4
	T  = true
	F  = false
)

var rocks = []rock{
	{
		4,
		[][WR]bool{
			{T, T, T, T},
		},
	},
	{
		3,
		[][WR]bool{
			{F, T, F},
			{T, T, T},
			{F, T, F},
		},
	},
	{
		3,
		[][WR]bool{
			{T, T, T},
			{F, F, T},
			{F, F, T},
		},
	},
	{
		1,
		[][WR]bool{
			{T},
			{T},
			{T},
			{T},
		},
	},
	{
		2,
		[][WR]bool{
			{T, T},
			{T, T},
		},
	},
}

type cave struct {
	filled [][WC]bool
	wind   wind
	period period
}

func (c *cave) height() int {
	return len(c.filled)
}

func (c *cave) collides(r *rock, x0, y0 int) bool {
	if x0 < 0 || x0+r.width > WC {
		return true
	}
	for yr := 0; yr < len(r.shape); yr++ {
		y := y0 + yr
		if y < 0 {
			return true
		}
		if y >= len(c.filled) {
			return false
		}
		for xr := 0; xr < r.width; xr++ {
			if r.shape[yr][xr] && c.filled[y][x0+xr] {
				return true
			}
		}
	}
	return false
}

func (c *cave) save(r *rock, x0, y0 int) {
	for yr := 0; yr < len(r.shape); yr++ {
		y := y0 + yr
		for y >= len(c.filled) {
			c.filled = append(c.filled, [WC]bool{})
		}
		for xr := 0; xr < r.width; xr++ {
			if r.shape[yr][xr] {
				c.filled[y][x0+xr] = true
			}
		}
	}
}

func (c *cave) print() {
	for y := len(c.filled) - 1; y >= 0; y-- {
		fmt.Print("|")
		for x := 0; x < WC; x++ {
			if c.filled[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("|")
	}
	fmt.Println("+-------+")
}

func (c *cave) printTop() {
	for y := len(c.filled) - 1; y >= max(0, len(c.filled)-9); y-- {
		fmt.Print("|")
		for x := 0; x < WC; x++ {
			if c.filled[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("|")
	}
	fmt.Println("+.......+")
}

type wind struct {
	d string
	i int
}

func (w *wind) next() (int, bool) {
	x := 1
	if w.d[w.i] == '<' {
		x = -1
	}

	w.i = (w.i + 1) % len(w.d)
	return x, w.i == 0
}

type interval struct {
	rocks  int
	height int
}

type state struct {
	r, y int
}

type period struct {
	states map[state]interval
	found  bool
	s, d   interval
}

func fall(c *cave, n int) {
	i := n % len(rocks)
	r := &rocks[i]
	x, y := 2, c.height()+3

	for {
		wd, wr := c.wind.next()
		if wr && !c.period.found {
			s := state{i, y - c.height()}
			if i, ok := c.period.states[s]; ok {
				c.period.s = i
				c.period.d = interval{n - i.rocks, c.height() - i.height}
				c.period.found = true
			} else {
				c.period.states[s] = interval{n, c.height()}
			}
		}

		if xn := x + wd; !c.collides(r, xn, y) {
			x = xn
		}

		if c.collides(r, x, y-1) {
			c.save(r, x, y)
			return
		}
		y--
	}
}

func solve(p *Problem) {
	p.NextLine()
	c := cave{
		period: period{states: map[state]interval{}},
		wind:   wind{d: p.Line()},
	}

	n := 0
	for n < 2022 {
		fall(&c, n)
		n++
	}
	p.PartOne(c.height())

	for !c.period.found || (n-c.period.s.rocks-1)%c.period.d.rocks != 0 {
		fall(&c, n)
		n++
	}

	const n2 = 1_000_000_000_000
	skip := (n2 - n) / c.period.d.rocks
	n += skip * c.period.d.rocks

	for n < n2 {
		fall(&c, n)
		n++
	}

	p.PartTwo(c.height() + skip*c.period.d.height)
}
