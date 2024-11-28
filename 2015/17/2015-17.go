package main

import (
	"math"
	"math/bits"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year = 2015
	day  = 17
)

func main() {
	Run(year, day, "", solve)
}

func solve(p *Problem) {
	const V = 150
	var vs []int
	for p.NextLine() {
		vs = append(vs, ParseInt(p.Line()))
	}

	n := uint32(1) << len(vs)
	best := math.MaxInt
	count1, count2 := 0, 0
	for i := uint32(0); i < n; i++ {
		v, mask := 0, uint32(1)
		for j := 0; j < len(vs) && v <= V; j++ {
			if i&mask != 0 {
				v += vs[j]
			}
			mask <<= 1
		}

		if v == V {
			b := bits.OnesCount32(i)
			if b < best {
				best = b
				count2 = 0
			}

			count1++

			if b == best {
				count2++
			}
		}
	}

	p.PartOne(count1)
	p.PartTwo(count2)
}
