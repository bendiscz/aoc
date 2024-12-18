package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"strings"
)

const (
	year    = 2019
	day     = 22
	example = `

deal into new stack
cut -2
deal with increment 7
cut 8
cut -4
deal with increment 7
cut 3
deal with increment 9
deal with increment 3
cut -1

`
)

func main() {
	Run(year, day, example, solve)
}

type lin struct {
	x, y int
}

func solve(p *Problem) {
	m, a := 10007, 2019
	if p.Example() {
		m, a = 10, 3
	}

	cs := []lin(nil)
	for p.NextLine() {
		s := p.Line()
		switch {
		case s == "deal into new stack":
			a = Mod(-a-1, m)
			cs = append(cs, lin{-1, -1})
		case strings.HasPrefix(s, "cut"):
			b := ParseInts(s)[0]
			a = Mod(a-b, m)
			cs = append(cs, lin{1, -b})
		default:
			b := ParseInts(s)[0]
			a = Mod(a*b, m)
			cs = append(cs, lin{b, 0})
		}
	}
	p.PartOne(a)

	if p.Example() {
		return
	}

	m, a = 119315717514047, 2020
	x, y := 1, 0
	for _, c := range cs {
		// a1 = a0*x0+y0 mod m
		// a2 = a1*x1+y1 mod m = (a0*x0+y0)*x1+y1 mod m = a0*x0*x1 + y0*x1 + y1
		// x2 = x1*x0, y2 = y0*x1 + y1
		x, y = Mod(c.x*x, m), Mod(c.x*y+c.y, m)
	}

	// ^n
	n := 101741582076661
	x0, y0 := 1, 0
	for n != 0 {
		if n&1 == 1 {
			x0, y0 = Mod(mul(x, x0, m), m), Mod(mul(x, y0, m)+y, m)
		}
		x, y = Mod(mul(x, x, m), m), Mod(mul(x, y, m)+y, m)
		n >>= 1
	}

	// a_n = ax+y mod m
	// a = (a_n-y) / x mod m
	_, x0inv, _ := GCDExt(x0, m)
	a = mul(Mod(a-y0, m), x0inv, m)

	p.PartTwo(a)
}

func mul(a, b, m int) int {
	c := 0
	for b != 0 {
		if b&1 == 1 {
			c = Mod(c+a, m)
		}
		a = Mod(a+a, m)
		b >>= 1
	}
	return c
}
