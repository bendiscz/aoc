package main

import (
	"fmt"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2016
	day     = 8
	example = `

rect 3x2
rotate column x=1 by 1
rotate row y=0 by 4
rotate column x=1 by 1

`
)

func main() {
	Run(year, day, example, solve)
}

const mx, my = 50, 6

type display struct {
	v [mx][my]bool
}

func (d *display) rect(w, h int) {
	if w > mx {
		w = mx
	} else if w < 0 {
		w = 0
	}
	if h > my {
		h = my
	} else if h < 0 {
		h = 0
	}

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			d.v[x][y] = true
		}
	}
}

func (d *display) rotateRow(y, l int) {
	l = mod(l, mx)
	var row [mx]bool
	for x := 0; x < mx; x++ {
		row[(x+l)%mx] = d.v[x][y]
	}
	for x := 0; x < mx; x++ {
		d.v[x][y] = row[x]
	}
}

func (d *display) rotateColumn(x, l int) {
	l = mod(l, my)
	var column [my]bool
	for y := 0; y < my; y++ {
		column[(y+l)%my] = d.v[x][y]
	}
	for y := 0; y < my; y++ {
		d.v[x][y] = column[y]
	}
}

func (d *display) countLit() int {
	count := 0
	for x := 0; x < mx; x++ {
		for y := 0; y < my; y++ {
			if d.v[x][y] {
				count++
			}
		}
	}
	return count
}

func (d *display) print() {
	for y := 0; y < my; y++ {
		for x := 0; x < mx; x++ {
			if d.v[x][y] {
				fmt.Print(SymbolFull)
			} else {
				fmt.Print(SymbolEmpty)
			}
		}
		fmt.Println()
	}
}

func mod(a, b int) int {
	a %= b
	if a < 0 {
		a += b
	}
	return a
}

func solve(p *Problem) {
	d := display{}
	for p.NextLine() {
		m := p.Parse(`^rect (\d+)x(\d+)$`)
		if m != nil {
			d.rect(ParseInt(m[1]), ParseInt(m[2]))
			continue
		}

		m = p.Parse(`^rotate (row|column) [xy]=(\d+) by (\d+)$`)
		if m == nil {
			continue
		}

		switch m[1] {
		case "row":
			d.rotateRow(ParseInt(m[2]), ParseInt(m[3]))
		case "column":
			d.rotateColumn(ParseInt(m[2]), ParseInt(m[3]))
		}
	}

	p.PartOne(d.countLit())
	p.PartTwo("")
	d.print()
}
