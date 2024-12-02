package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"slices"
)

const (
	year    = 2024
	day     = 2
	example = `

7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	s1, s2 := 0, 0
	for p.NextLine() {
		f := ParseInts(p.Line())
		if check(f) {
			s1++
			s2++
		} else {
			for s := 0; s < len(f); s++ {
				if check(slices.Delete(slices.Clone(f), s, s+1)) {
					s2++
					break
				}
			}
		}
	}

	p.PartOne(s1)
	p.PartTwo(s2)
}

func check(f []int) bool {
	inc, dec := true, true
	prev := f[0]
	for _, x := range f[1:] {
		if Abs(prev-x) > 3 {
			return false
		}
		if x >= prev {
			dec = false
		}
		if x <= prev {
			inc = false
		}
		prev = x
	}
	return inc || dec
}
