package main

import (
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2025
	day     = 12
	example = `

0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 2 1 0 3 2

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	shapes := []int(nil)
	for p.NextLine() {
		if !strings.HasSuffix(p.Line(), ":") {
			break
		}

		area := 0
		for p.NextLine() && p.Line() != "" {
			for _, ch := range p.Line() {
				if ch == '#' {
					area++
				}
			}
		}
		shapes = append(shapes, area)
	}

	s1 := 0
	for {
		f := ParseInts(p.Line())
		area := 0
		for i, x := range f[2:] {
			area += x * shapes[i]
		}
		if area <= f[0]*f[1] {
			s1++
		}

		if !p.NextLine() {
			break
		}
	}

	p.PartOne(s1)
}
