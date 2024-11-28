package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 16
	example = `

s1,x3/4,pe/b

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	n := 16
	if p.Example() {
		n = 5
	}

	d := parseDance(p.ReadAll(), n)

	prg := identityPerm(n)
	for i := 0; i < n; i++ {
		prg[i] = i
	}

	d.do(prg)
	p.PartOne(prg.format())

	d.pow(999_999_999)
	d.do(prg)
	p.PartTwo(prg.format())
}

type perm []int

func makePerm(n int) perm {
	return make([]int, n)
}

func identityPerm(n int) perm {
	p := makePerm(n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return p
}

func (p perm) pow(n int) []int {
	r := identityPerm(len(p))
	for i := 0; i < len(p); i++ {
		r[i] = i
	}
	for n > 0 {
		if n&1 == 1 {
			r = r.mul(p)
		}
		p = p.mul(p)
		n >>= 1
	}
	return r
}

func (p perm) mul(q []int) perm {
	c := makePerm(len(p))
	for i := 0; i < len(p); i++ {
		c[i] = p[q[i]]
	}
	return c
}

func (p perm) format() string {
	s := make([]byte, len(p))
	for i, b := range p {
		s[i] = byte(b + 'a')
	}
	return string(s)
}

type dance struct {
	n   int
	pos perm
	prg perm
}

func parseDance(s string, n int) *dance {
	d := &dance{
		n:   n,
		pos: identityPerm(n),
		prg: identityPerm(n),
	}

	shift := 0
	t := identityPerm(n)
	for _, m := range SplitFields(s) {
		switch m[0] {
		case 's':
			x := ParseInt(m[1:])
			shift = (shift + x) % n

		case 'x':
			x := ParseInts(m[1:])
			a, b := (x[0]-shift+n)%n, (x[1]-shift+n)%n
			d.pos[a], d.pos[b] = d.pos[b], d.pos[a]

		case 'p':
			a0, b0 := int(m[1]-'a'), int(m[3]-'a')
			a1, b1 := t[a0], t[b0]
			t[a0], t[b0] = b1, a1
			d.prg[a1], d.prg[b1] = d.prg[b1], d.prg[a1]
		}
	}

	for i := 0; i < n; i++ {
		t[(i+shift)%n] = d.pos[i]
	}
	copy(d.pos, t)

	return d
}

func (d *dance) do(p perm) {
	x := makePerm(d.n)
	for i := 0; i < d.n; i++ {
		x[i] = p[d.pos[i]]
	}
	for i := 0; i < d.n; i++ {
		p[i] = d.prg[x[i]]
	}
}

func (d *dance) pow(n int) {
	d.pos = d.pos.pow(n)
	d.prg = d.prg.pow(n)
}
