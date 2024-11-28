package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2015
	day     = 25
	example = `

To continue, please consult the code grid in the manual.  Enter the code at row 4, column 2.

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	r, c := 0, 0
	p.NextLine()
	p.Scanf("To continue, please consult the code grid in the manual. Enter the code at row %d, column %d.", &r, &c)

	d := r + c - 1
	n := d*(d-1)/2 + c - 1

	const a, b = 252533, 33554393
	code := 20151125

	for q := a; n != 0; n >>= 1 {
		if n&1 == 1 {
			code = (code * q) % b
		}
		q = (q * q) % b
	}
	p.PartOne(code)
}
