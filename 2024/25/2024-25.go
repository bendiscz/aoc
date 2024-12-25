package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2024
	day     = 25
	example = `

#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####

`
)

func main() {
	Run(year, day, example, solve)
}

type shape [5]int

func parseShape(p *Problem) shape {
	x, k, last := shape{}, 1, p.Line()
	for p.NextLine() && p.Line() != "" {
		s := p.Line()
		for i := 0; i < len(x); i++ {
			if x[i] == 0 && s[i] != last[i] {
				x[i] = k
			}
		}
		last = s
		k++
	}
	return x
}

func solve(p *Problem) {
	locks := []shape(nil)
	keys := []shape(nil)
	for p.NextLine() {
		if p.Line()[0] == '#' {
			locks = append(locks, parseShape(p))
		} else {
			keys = append(keys, parseShape(p))
		}
	}

	s1 := 0
	for _, key := range keys {
		for _, lock := range locks {
			fits := true
			for i := 0; i < len(key) && fits; i++ {
				if key[i] < lock[i] {
					fits = false
				}
			}
			if fits {
				s1++
			}
		}
	}
	p.PartOne(s1)

	p.PartTwo("Merry Christmas!")
}
