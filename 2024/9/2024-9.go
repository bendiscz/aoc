package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"slices"
)

const (
	year    = 2024
	day     = 9
	example = `

2333133121414131402

`
)

func main() {
	Run(year, day, example, solve)
}

type block struct {
	free bool
	id   int
}

type seq struct {
	b    block
	size int
}

func solve(p *Problem) {
	p.NextLine()
	s := p.Line()

	bs := []block(nil)
	ss := []seq(nil)
	for i, id, b := 0, 0, (block{}); i < len(s); i, b.free = i+1, !b.free {
		x := int(s[i] - '0')
		if x == 0 {
			continue
		}

		if b.free {
			b.id = 0
		} else {
			b.id = id
			id++
		}

		ss = append(ss, seq{b, x})
		for ; x > 0; x-- {
			bs = append(bs, b)
		}
	}

	for i0, i1 := 0, len(bs)-1; ; {
		if i0 == i1 {
			break
		}
		if !bs[i0].free {
			i0++
			continue
		}
		if bs[i1].free {
			i1--
			continue
		}
		bs[i0], bs[i1] = bs[i1], bs[i0]
	}

	s1 := 0
	for i, b := range bs {
		if b.free {
			break
		}
		s1 += i * bs[i].id
	}
	p.PartOne(s1)

	for fp := len(ss) - 1; fp > 0; fp-- {
		if ss[fp].b.free {
			continue
		}

		for k := 0; k < len(ss) && k < fp; k++ {
			if ss[k].b.free && ss[k].size >= ss[fp].size {
				f := ss[fp]
				ss[k].size -= f.size
				ss[fp].b = block{free: true}

				if ss[k].size > 0 {
					ss = slices.Insert(ss, k, f)
					fp++
				} else {
					ss[k] = f
				}
				break
			}
		}
	}

	s2, i := 0, 0
	for _, sq := range ss {
		if sq.b.free {
			i += sq.size
		} else {
			for sq.size > 0 {
				s2 += i * sq.b.id
				i++
				sq.size--
			}
		}
	}
	p.PartTwo(s2)
}
