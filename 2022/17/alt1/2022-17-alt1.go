package main

import (
	"fmt"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year = 2022
	day  = 17
)

func main() {
	Run(year, day, "", solve)
}

type rock struct {
	width int
	shape []uint8
}

const (
	WC = 7
)

var rocks = []rock{
	{
		4,
		[]uint8{
			0b1111,
		},
	},
	{
		3,
		[]uint8{
			0b010,
			0b111,
			0b010,
		},
	},
	{
		3,
		[]uint8{
			0b111,
			0b100,
			0b100,
		},
	},
	{
		1,
		[]uint8{
			0b1,
			0b1,
			0b1,
			0b1,
		},
	},
	{
		2,
		[]uint8{
			0b11,
			0b11,
		},
	},
}

type cave struct {
	filled []uint8
	bottom int
	wind   wind
}

func (c *cave) height() int {
	return len(c.filled) + c.bottom
}

func (c *cave) collides(r *rock, x0, y0 int) bool {
	if x0 < 0 || x0+r.width > WC {
		return true
	}
	for yr := 0; yr < len(r.shape); yr++ {
		y := y0 + yr
		if y >= len(c.filled) {
			return false
		}
		if y < 0 {
			return true
		}
		if r.shape[yr]<<x0&c.filled[y] != 0 {
			return true
		}
	}
	return false
}

func (c *cave) save(r *rock, x0, y0 int) {
	for yr := 0; yr < len(r.shape); yr++ {
		y := y0 + yr
		for y >= len(c.filled) {
			c.filled = append(c.filled, 0)
		}
		c.filled[y] |= r.shape[yr] << x0
	}
}

func (c *cave) cut() {
	const (
		limit  = 1_000_000
		remove = limit - 1_000
	)

	if len(c.filled) < limit {
		return
	}

	copy(c.filled, c.filled[remove:])
	c.filled = c.filled[:len(c.filled)-remove]
	c.bottom += remove
}

func (c *cave) print() {
	for y := len(c.filled) - 1; y >= 0; y-- {
		fmt.Print("|")
		for x := 0; x < WC; x++ {
			if c.filled[y]>>x&1 != 0 {
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
			if c.filled[y]>>x&1 != 0 {
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

func (w *wind) next() int {
	x := 1
	if w.d[w.i] == '<' {
		x = -1
	}

	w.i = (w.i + 1) % len(w.d)
	return x
}

func fall(c *cave, n int) {
	i := n % len(rocks)
	r := &rocks[i]
	x, y := 2, len(c.filled)+3

	for {
		if xn := x + c.wind.next(); !c.collides(r, xn, y) {
			x = xn
		}

		if c.collides(r, x, y-1) {
			c.save(r, x, y)
			c.cut()
			return
		}
		y--
	}
}

func solve(p *Problem) {
	p.NextLine()
	c := cave{
		wind: wind{d: p.Line()},
	}

	n := 0
	for n < 2022 {
		fall(&c, n)
		n++
	}
	p.PartOne(c.height())

	const n2 = 1_000_000_000_000
	const r = 100_000_000
	done := 0
	for n < n2 {
		fall(&c, n)
		n++
		if n%r == 0 {
			if d := n / (n2 / 100); d > done {
				done = d
				p.Printf("%d%% done", done)
			}
		}
	}
	p.PartTwo(c.height())
}
