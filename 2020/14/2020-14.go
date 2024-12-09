package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"strings"
)

const (
	year    = 2020
	day     = 14
	example = `

mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	var m0, m1 uint
	var mf []uint
	mem1 := map[uint]uint{}
	mem2 := map[uint]uint{}
	for p.NextLine() {
		s := p.Line()
		if strings.HasPrefix(s, "mask") {
			mask := SplitFields(s)[2]
			m0, m1, mf = 0, 0, nil
			for i := 0; i < len(mask); i++ {
				b := uint(1) << (len(mask) - i - 1)
				switch mask[i] {
				case '1':
					m1 |= b
				case '0':
					m0 |= b
				default:
					mf = append(mf, b)
				}
			}
			continue
		}

		f := ParseUints(s)
		mem1[f[0]] = (f[1] | m1) & ^m0

		for i := 0; i < 1<<len(mf); i++ {
			addr := f[0] | m1
			for j, b := range mf {
				if i&(1<<j) != 0 {
					addr |= b
				} else {
					addr &= ^b
				}
			}
			mem2[addr] = f[1]
		}
	}

	p.PartOne(sum(mem1))
	p.PartTwo(sum(mem2))
}

func sum(mem map[uint]uint) (s uint) {
	for _, v := range mem {
		s += v
	}
	return
}
